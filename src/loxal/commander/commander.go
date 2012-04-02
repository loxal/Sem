// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// RESTful commander service backend
package commander

import (
	"fmt"
	"net/http"
	"io"
	"flag"
	"strings"
	"time"
	"encoding/json"
	"loxal/test"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

type Cmd struct {
	Name, Call, Desc    string
	Creator, User       string
	Created, Updated    time.Time
}

func serveError(c appengine.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", contentTypeText)
	io.WriteString(w, "Internal Server Error")
	c.Errorf("%s", err)
}

func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", contentTypeText)
	io.WriteString(w, "Not Found")
}

func getUser(c appengine.Context) string {
    u := user.Current(c)
    if u == nil {
        return ""
    }

    return u.Email
}

func createCmd(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	if !cmdExists(c, r.FormValue("name")) && !cmdHasInvalidCharacters(r.FormValue("name")) {
	    currentUser := getUser(c);
		cmd := &Cmd{
			Name:     r.FormValue("name"),
			Call:     r.FormValue("call"),
			Desc:     r.FormValue("desc"),
			Creator:  currentUser,
			User:     currentUser,
			Created:  time.Now(),
			Updated:  time.Now(),
		}

		if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Cmd", nil), cmd); err != nil {
			serveError(c, w, err)
			return
		}

		addCacheItem(r, cmd)
		return
	}

    w.WriteHeader(http.StatusBadRequest)
}

// Constraint Check
func cmdExists(c appengine.Context, name string) (ok bool) {
    count, _ := datastore.NewQuery("Cmd").Filter("Name =", strings.ToLower(name)).Filter("Creator =", getUser(c)).Count(c)

	if count > 0 {
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

func listCmds(r *http.Request) (cmds []*Cmd) {
	c := appengine.NewContext(r)
	datastore.NewQuery("Cmd").Filter("Creator =", getUser(c)).GetAll(c, &cmds)

	return cmds
}

func listCmdsAsJSON(w http.ResponseWriter, r *http.Request) {
	cmds := listCmds(r)

	w.Header().Set("Content-Type", contentTypeJSON)
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
//func exec(cmd string) (call string) {
func getCmd(r *http.Request) (call, query string) {
    const sep = "+"
    rawQuery := strings.Split(r.URL.RawQuery, sep)
    getCacheItem(r, rawQuery[0])

    cmds := listCmds(r)

    for i := range cmds {
        if cmds[i].Name == rawQuery[0] {
            call = cmds[0].Call
            query = strings.Join(rawQuery[1:], sep)
            return call, query
        }
    }

    const defaultRestCall = "http://www.google.com/search?q="
    call = defaultRestCall
    query = strings.Join(rawQuery[:], sep)

    return call, query
}

func exec(w http.ResponseWriter, r *http.Request) {
//test.TestFlag(w)
fmt.Fprintf(w, test.ParseQuery("task --add='myNewTask bond'"))

    call, query := getCmd(r)
    http.Redirect(w, r, call + query, http.StatusFound)
}

func cmd(w http.ResponseWriter, r *http.Request) {
	var _ = flag.PrintDefaults // delete before submitting
	c := appengine.NewContext(r)
	var cmds []*Cmd
	cmdName := r.FormValue("name")
	_, err := datastore.NewQuery("Cmd").Filter("Name =", cmdName).GetAll(c, &cmds)
	fmt.Fprintln(w, err)
	fmt.Fprintln(w, r.FormValue("cmd"))
}

func cmdDelete(cmdName string, c appengine.Context) (ok bool) {
	q := datastore.NewQuery("Cmd").Filter("Name =", cmdName).Filter("Creator =", getUser(c)).KeysOnly()
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
		Call: r.FormValue("edit-call"),
		Desc:     r.FormValue("edit-desc"),
		User:  getUser(c),
		Updated:  time.Now(),
	}

	if ok, err := cmdUpdate(cmd, c); err != nil {
		fmt.Fprintln(w, err, ok)
	}
}

func cmdUpdate(cmd *Cmd, c appengine.Context) (ok bool, err error) {
	q := datastore.NewQuery("Cmd").Filter("Name =", cmd.Name).Filter("Creator =", getUser(c)).KeysOnly()
	keys, _ := q.GetAll(c, nil)

	var cmdInDS Cmd
	datastore.Get(c, keys[0], &cmdInDS)
    cmd.Creator = cmdInDS.Creator
    cmd.Created = cmdInDS.Created

	if _, err := datastore.Put(c, keys[0], cmd); err != nil {
		return false, err
	}
	return true, nil
}

func authentication(r *http.Request) (usr, url string, isAdmin bool) {
    c := appengine.NewContext(r)
    u := user.Current(c)

    if u == nil {
        url, _ = user.LoginURL(c, indexHandler)
    } else {
        usr = u.Email
        isAdmin = user.IsAdmin(c)
        url, _ = user.LogoutURL(c, indexHandler)
    }

    return usr, url, isAdmin
}

func authenticate(w http.ResponseWriter, r *http.Request) {
    user, url, isAdmin := authentication(r)

    w.Header().Set("Content-Type", contentTypeJSON)
    fmt.Fprintf(w, `{"user": "%s", "isAdmin": "%t", "url": "%s"}`, user, isAdmin, url)
}

const indexHandler = "/"
const cmdUpdateHandler = "/cmd/update"
const cmdDeleteHandler = "/cmd/delete"
const cmdCreateHandler = "/cmd/create"
const cmdListHandler = "/cmd/list.json"
const reflectHandler = "/test/reflect"

const contentTypeJSON = "application/json; charset=utf-8"
const contentTypeText = "text/plain; charset=utf-8"

func init() {
	http.HandleFunc(reflectHandler, test.ReflectFunc)
	http.HandleFunc("/cmd/auth.json", authenticate)
	http.HandleFunc(cmdDeleteHandler, cmdDeletion)
	http.HandleFunc(cmdUpdateHandler, cmdUpdation)
	http.HandleFunc(cmdCreateHandler, createCmd)
	http.HandleFunc(cmdListHandler, listCmdsAsJSON)
	http.HandleFunc("/cmd", cmd)
	http.HandleFunc("/cmd/exec", exec)
}
