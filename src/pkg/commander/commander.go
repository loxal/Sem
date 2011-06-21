// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// RESTful commander service backend
package commander

import (
	"fmt"
	"http"
	"io"
	"os"
	"strings"
	"time"
	"json"

	"flag"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

type Cmd struct {
	Name, RESTcall, Desc string
	Creator, User        string
	Created, Updated     datastore.Time
}

func serveError(c appengine.Context, w http.ResponseWriter, err os.Error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Internal Server Error")
}

func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Not Found")
}

func cmdCreation(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if err := r.ParseForm(); err != nil {
		serveError(c, w, err)
		return
	}

	if !cmdExists(c, r.FormValue("name")) && !cmdHasInvalidCharacters(r.FormValue("name")) {
		cmd := &Cmd{
			Name:     r.FormValue("name"),
			RESTcall: r.FormValue("restCall"),
			Desc:     r.FormValue("desc"),
			Created:  datastore.SecondsToTime(time.Seconds()),
		}
		if u := user.Current(c); u != nil {
			cmd.Creator = u.String()
		}
		if _, err := datastore.Put(c, datastore.NewIncompleteKey("Cmd"), cmd); err != nil {
			serveError(c, w, err)
			return
		}
	}
	http.Redirect(w, r, indexHandler, http.StatusFound)
}

// Constraint Check
func cmdExists(c appengine.Context, name string) (ok bool) {
	if count, err := datastore.NewQuery("Cmd").Filter("Name =", name).Count(c); err == nil && count > 0 {
		return true
	}
	return
}

// Constraint Check
func cmdHasInvalidCharacters(name string) (ok bool) {
	// command name shouldn't contain "#" because it's the HTML anchor marker and
	// might cause problems in a RESTful context (acts as a delimiter)
	if strings.Contains(name, "#") || strings.Contains(name, "%") {
		return true
	}
	return
}

func cmdListing(w http.ResponseWriter, r *http.Request) (cmds []*Cmd) {
	c := appengine.NewContext(r)
	if _, err := datastore.NewQuery("Cmd").GetAll(c, &cmds); err != nil {
		serveError(c, w, err)
		return
	}
	return cmds
}

func cmdListingJson(w http.ResponseWriter, r *http.Request) {
	cmds := cmdListing(w, r)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, `{"cmds": [`)
	lenCmds := len(cmds)
	for i := range cmds {
		cmdJSONed, _ := json.Marshal(cmds[i])
		fmt.Fprint(w, string(cmdJSONed))
		if lenCmds-1 != i {
			fmt.Fprintln(w, ",")
		}
	}
	fmt.Fprint(w, "]}")
}

// Returns the RESTful associated with a certain command
//func exec(cmd string) (restCall string) {
func exec(w http.ResponseWriter, r *http.Request) {
}

func cmd(w http.ResponseWriter, r *http.Request) {
	//	var _ = flag1.PrintDefaults // delete before submitting
	var _ = flag.PrintDefaults // delete before submitting
	c := appengine.NewContext(r)
	var cmds []*Cmd
	cmdName := r.FormValue("name")
	_, err := datastore.NewQuery("Cmd").Filter("Name =", cmdName).GetAll(c, &cmds)
	fmt.Fprintln(w, err)
	fmt.Fprintln(w, r.FormValue("cmd"))
}

func cmdDelete(cmdName string, c appengine.Context) (ok bool) {
	q := datastore.NewQuery("Cmd").Filter("Name =", cmdName).KeysOnly()
	keys, _ := q.GetAll(c, nil)
	if err := datastore.Delete(c, keys[0]); err != nil {
		return
	}
	return true
}

func cmdDeletion(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	cmdDelete(r.FormValue("name"), c)
	http.Redirect(w, r, indexHandler, http.StatusFound)
}

func cmdUpdation(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	cmd := &Cmd{
		Name:     r.FormValue("edit-name"),
		RESTcall: r.FormValue("edit-restCall"),
		Desc:     r.FormValue("edit-desc"),
		Updated:  datastore.SecondsToTime(time.Seconds()),
	}
	if u := user.Current(c); u != nil {
		cmd.User = u.String()
	}
	if ok, err := cmdUpdate(cmd, c); err != nil {
		fmt.Fprintln(w, err, ok)
	}
}

func cmdUpdate(cmd *Cmd, c appengine.Context) (ok bool, err os.Error) {
	q := datastore.NewQuery("Cmd").KeysOnly().Filter("Name =", cmd.Name)
	keys, _ := q.GetAll(c, nil)
	if _, err := datastore.Put(c, keys[0], cmd); err != nil {
		return false, err
	}
	return true, nil

	return false, os.NewError("exists")
}

const indexHandler = "/"
const cmdUpdateHandler = "/cmd/update"
const cmdDeleteHandler = "/cmd/delete"
const cmdCreateHandler = "/cmd/create"
const cmdListHandler = "/cmd/list.json"

func init() {

	http.HandleFunc(cmdDeleteHandler, cmdDeletion)
	http.HandleFunc(cmdUpdateHandler, cmdUpdation)
	http.HandleFunc(cmdCreateHandler, cmdCreation)
	http.HandleFunc(cmdListHandler, cmdListingJson)
	http.HandleFunc("/cmd", cmd)
	http.HandleFunc("/exec", exec)
}
