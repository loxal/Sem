// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The entry package for the GAE environment
package commander

import (
	"fmt"
	"http"
	"io"
	"os"
	"strconv"
	"strings"
	"template"
	"time"
	"json"

	"flag1" // to parse r.URL.Raw

	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"appengine/user"
)

type Cmd struct {
	Name, RESTcall, Desc string
	Creator, User        string
	Created, Updated     datastore.Time
}

type Greeting struct {
	Author  string
	Content string
	Date    datastore.Time
	Title   string
	Body    string
	Pg      *page
}

type page struct {
	Title1 string
	Body1  string
}

// TODO make it a small letter pAGE
type Page1 struct {
	Title11 string
	Body11  string
}

func loadPage(title string) (*page, os.Error) {
	//	filename := title + ".txt"
	//	body, err := ioutil.ReadFile(filename)
	//	if err != nil {
	//		return nil, err
	//	}
	return &page{Title1: title, Body1: "test"}, nil
}

//func hello(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	http.Redirect(w, r, "/index", http.StatusFound)
//}

func serveError(c appengine.Context, w http.ResponseWriter, err os.Error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Internal Server Error")
	c.Logf("%v", err)
}

func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Not Found")
}

func count(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	item, err := memcache.Get(c, r.URL.Path)
	if err != nil && err != memcache.ErrCacheMiss {
		serveError(c, w, err)
		return
	}
	n := 0
	if err == nil {
		n, err = strconv.Atoi(string(item.Value))
		if err != nil {
			serveError(c, w, err)
			return
		}
	}
	n++
	item = &memcache.Item{
		Key:   r.URL.Path,
		Value: []byte(strconv.Itoa(n)),
	}
	err = memcache.Set(c, item)
	if err != nil {
		serveError(c, w, err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%q has been visited %d times", r.URL.Path, n)
}

func cmdCreation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		serve404(w)
		return
	}
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
	http.Redirect(w, r, cmdListHandler, http.StatusFound)
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
    if strings.Contains(name, "#") || strings.Contains(name, "%")  { return true }
    return
}

func cmdListing(w http.ResponseWriter, r *http.Request) (cmds []*Cmd) {
	if r.Method != "GET" {
		serve404(w)
		return
	}
	c := appengine.NewContext(r)
	if _, err := datastore.NewQuery("Cmd").GetAll(c, &cmds); err != nil {
		serveError(c, w, err)
		return
	}

	return cmds
}

func cmdListingHtml(w http.ResponseWriter, r *http.Request) {
        cmds := cmdListing(w, r)

		w.Header().Set("Content-Type", "text/html")
		if err := createCmdPresenter.Execute(w, cmds); err != nil {
			fmt.Println("%v", err)
		}
}

func cmdListingJson(w http.ResponseWriter, r *http.Request) {
        cmds := cmdListing(w, r)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		for i := range cmds {
			cmdJSONed, _ := json.Marshal(cmds[i])
			fmt.Fprintln(w, string(cmdJSONed))
		}
}

//func exec(url *http.URL) {
//
//}

// Returns the RESTful associated with a certain command
func exec(cmd string) (restCall string) {

	//	io.WriteString(os.Stdout, url.Raw + "\n")
	//	os.Stdout.WriteString(url.Raw + "\n")

	//	restCall = m[cmd]
	return
}

func cmd(w http.ResponseWriter, r *http.Request) {
     var _ = flag1.PrintDefaults // delete before submitting
	    c := appengine.NewContext(r)
	//    c.Logf("r.URL.Path: " + r.URL.Path)
	//    c.Logf("r.URL.RawQuery: " + r.URL.RawQuery)
	//     c.Logf("m[r.URL.RawQuery]" + m[r.URL.RawQuery])
	//	    http.Redirect(w, r, m[r.URL.RawQuery], http.StatusFound)

	//    w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//    io.WriteString(w, r.URL.Path + "\n")
	//    io.WriteString(w, r.URL.RawQuery + "\n")
	//    io.WriteString(w, r.URL.Raw + "\n")

	//	exec(r.URL)
	//	restCall := exec(r.FormValue("name"))

	var cmds []*Cmd
	cmdName := r.FormValue("name")
	_, err := datastore.NewQuery("Cmd").Filter("Name =", cmdName).GetAll(c, &cmds)
	fmt.Fprintln(w, err)
	// retrieve this from the datastore; put this as an init dataset into the datastore via *_test.go TODO
	//	m := map[string]string{
	//		"c":    "https://mail.google.com/mail/?shva=1#compose",
	//		"d":    "https://mail.google.com/tasks/canvas",
	//		"t":    "http://twitter.com",
	//		"sem":  "https://github.com/loxal/Sem",
	//		"verp": "https://github.com/loxal/Verp",
	//		"lox":  "https://github.com/loxal/Lox",
	//		// shortcut for adding an English Word or another unknow word to the TO_LEARN_LIST (merge with the Delingo functionality)
	//		// shortcut for making notes/tasks/todos
	//	}

	//    flag1.Parse1([]string{"app", "-areacode", "333", "-param1", "areacode"})
	//    flag1.Parse1([]string{"-areacode", "333", "-param1", "areacode"})
	//    var *int p
	//    flag1.IntVar(&p, "param", 3, "usage")

	//args := []string{"/tmp/dev_appserver_alex_8080_go_app_work_dir/_go_app", "-addr_http", "unix:/tmp/dev_appserver_alex_8080_socket_http", "-addr_api", "unix:/tmp/dev_appserver_alex_8080_socket_api"}
	//os.Args = args
	//	args := []string{"app", "-name", "9999999999999", "-desc", "VAL"}
	//args1 := []string{"app", "-name", "999", "-desc", "VAL"}
	//fmt.Print(args1)
	//	name := flag1.String("name", "33333", "name")
	//	var name *string
	//	flag1.StringVar(name, "name", "33333", "name")
	//	desc := flag1.String("desc", "myValue", "desc")
	//    var desc *string
	//	flag1.StringVar(desc, "desc", "myValue", "desc")
	//	flag1.Parse()
	//	var _ = fmt.Printf // delete before submitting


	//	fmt.Fprintln(w, "name ", *name)
	//	fmt.Fprintln(w, "desc", *desc)
	//  fmt.Fprintln(w, os.Args[1:]);
	//  fmt.Fprintln(w, "\n");
	//  fmt.Fprintln(w, args);

	//	fmt.Fprintln(w, &p)
	//	fmt.Fprintln(w, flag1.Args())
	//	fmt.Fprintln(w, flag1.Arg(0))
	//	fmt.Fprintln(w, flag1.Arg(1))
	//	fmt.Fprintln(w, flag1.Arg(2))
	//	fmt.Fprintln(w, os.Args)
	//	fmt.Fprintln(w, )

	//    io.WriteString(w, cmds[0].RESTcall)
	//	http.Redirect(w, r, cmds[0].RESTcall, http.StatusFound)

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
	http.Redirect(w, r, cmdListHandler, http.StatusFound)
}

func cmdUpdation(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	cmd := &Cmd{
		Name:     r.FormValue("edit-name"),
		RESTcall: r.FormValue("edit-restCall"),
		Desc:     r.FormValue("edit-desc"),
		// Creator TODO
		// Created TODO
		Updated: datastore.SecondsToTime(time.Seconds()),
		User:    user.Current(c).String(),
	}
	if ok, err := cmdUpdate(cmd, c); err != nil {
		fmt.Fprintln(w, err, ok)
	}

	http.Redirect(w, r, cmdListHandler, http.StatusFound)
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

const cmdCreateHandler = "/create"
const cmdListHandler = "/cmdList"

var createCmdPresenter *template.Template

func Double(i int) int {
	return i * 2
}

func init() {
	createCmdPresenter = template.New(nil)
	createCmdPresenter.SetDelims("{%", "%}")
	if err := createCmdPresenter.ParseFile("src/pkg/commander/main.html"); err != nil {
		panic("can't parse: " + err.String())
	}
//	http.HandleFunc("/", hello)
	http.HandleFunc("/cmdDelete", cmdDeletion)
	http.HandleFunc("/update", cmdUpdation)
//	http.HandleFunc("/hello", hello)
	http.HandleFunc(cmdCreateHandler, cmdCreation)
	http.HandleFunc(cmdListHandler, cmdListingHtml)
	http.HandleFunc(cmdListHandler + ".json", cmdListingJson)
	http.HandleFunc("/count", count)
	http.HandleFunc("/cmd", cmd)
	//		http.HandleFunc("/exec", exec)
}
