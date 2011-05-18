// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The entry package for the GAE environment
package main

import (
	"fmt"
	"http"
	"io"
	"os"
	"strconv"
	"template"
	"time"
	"json"
	//		"flag" // to parse r.URL.Raw
	//    "./flag_osArgs-less"
	//    "../flag1/flag1"

	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"appengine/user"
)

type Cmd struct {
	Name, RESTcall, Desc string
	User                 string
	Created              datastore.Time
}

type Greeting struct {
	Author  string
	Content string
	Date    datastore.Time
	Title   string
	Body    string
	//	Pg  *page
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

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Hello, ...!\n")
}

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

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" || r.URL.Path != cmdCreateHandler {
		serve404(w)
		return
	}
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Greeting")
	//	q := datastore.NewQuery("Greeting").Order("-Date").Limit(10)
	var gg []*Greeting
	_, err := q.GetAll(c, &gg)
	if err != nil {
		serveError(c, w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	//	if err := mainPage.Execute(w, gg); err != nil {
	//		c.Logf("%v", err)
	//	}

	//	for i := 0; i < len(gg); i++ {
	//        gg[i]= &Greeting{Title: "my TITLE", Body: "my BODY", Pg: &page{Title1: "fest1111", Body1: "test"}}
	//        gg[i]= &Greeting{Title: "my TITLE", Body: "my BODY"}
	//	}

	//    pg1 := &Page1{Title11: "my1111", Body11: "yours"}
	if err := createCmdPresenter.Execute(w, gg); err != nil {
		c.Logf("%v", err)
	}

}


func handleStore(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		serve404(w)
		return
	}
	c := appengine.NewContext(r)
	if err := r.ParseForm(); err != nil {
		serveError(c, w, err)
		return
	}
	g := &Greeting{
		Content: r.FormValue("content"),
		Date:    datastore.SecondsToTime(time.Seconds()),
	}
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	if _, err := datastore.Put(c, datastore.NewIncompleteKey("Greeting"), g); err != nil {
		serveError(c, w, err)
		return
	}
	http.Redirect(w, r, cmdCreateHandler, http.StatusFound)
}


// Returns the RESTful associated with a certain command
func WebCmd(cmd string) (restCall string) {
	m := map[string]string{
		"c":    "https://mail.google.com/mail/?shva=1#compose",
		"d":    "https://mail.google.com/tasks/canvas",
		"t":    "http://twitter.com",
		"sem":  "https://github.com/loxal/Sem",
		"verp": "https://github.com/loxal/Verp",
		"lox":  "https://github.com/loxal/Lox",
		// shortcut for adding an English Word or another unknow word to the TO_LEARN_LIST (merge with the Delingo functionality)
		// shortcut for making notes/tasks/todos
	}

	restCall = m[cmd]
	return
}

func cmdCreation(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	cmd := &Cmd{
		Name:     r.FormValue("name"),
		RESTcall: r.FormValue("restCall"),
		Desc:     r.FormValue("desc"),
		User:     user.Current(c).String(),
		Created:  datastore.SecondsToTime(time.Seconds()),
	}
	datastore.Put(c, datastore.NewIncompleteKey("Cmd"), cmd)
}

func cmdListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var cmds []*Cmd
	q := datastore.NewQuery("Cmd")

	if keys, err := q.GetAll(appengine.NewContext(r), &cmds); err == nil {
		for i := range keys {
			cmdJSONed, _ := json.Marshal(cmds[i])
			fmt.Fprintln(w, i, string(cmdJSONed))
		}
	}

	fmt.Fprintln(w, os.Args, ",,,,,,,,,,,,")

	//	fmt.Fprintln(w, flag.Args(), ",,,,,,,,,,,,")

	//	var keys []*datastore.Key
	//    q1 :=q.KeysOnly()
	//    count,e := q1.Filter("Name=", "my2").Count(c)


}

func exec(url *http.URL) {
	//	io.WriteString(os.Stdout, url.Raw + "\n")
	//	os.Stdout.WriteString(url.Raw + "\n")

	//	    http.Redirect(w, r, "http://sol.loxal.net", http.StatusFound)
}

func cmd(w http.ResponseWriter, r *http.Request) {
	//    c := appengine.NewContext(r)
	//    c.Logf("r.URL.Path: " + r.URL.Path)
	//    c.Logf("r.URL.RawQuery: " + r.URL.RawQuery)
	//     c.Logf("m[r.URL.RawQuery]" + m[r.URL.RawQuery])
	//	    http.Redirect(w, r, m[r.URL.RawQuery], http.StatusFound)

	//    w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//    io.WriteString(w, r.URL.Path + "\n")
	//    io.WriteString(w, r.URL.RawQuery + "\n")
	//    io.WriteString(w, r.URL.Raw + "\n")

	exec(r.URL)
}

func cmdDelete(cmdName string, c appengine.Context) (deleted bool) {
	q := datastore.NewQuery("Cmd").Filter("Name =", cmdName).KeysOnly()
	keys, _ := q.GetAll(c, nil)
    if err := datastore.Delete(c, keys[0]); err != nil {
        return
    }
    return true
}

func cmdDeletion(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	fmt.Println(cmdDelete(r.FormValue("name"), c))
}

func cmdUpdate(current *datastore.Key, new Cmd)

var cmdCreateHandler = "/cmdCreate"
var postHandler = "/post"
var storeHandler = "/store"
var createCmdPresenter = template.MustParseFile("cmdCreate.html", nil)
var mainPage = template.MustParseFile("template.html", nil)

func Double(i int) int {
	return i * 2
}

func init() {
	//flag.args = os.Args
	fmt.Println(os.Args, ",,,,,,,,,,,,<<OS<")
	//fmt.Println(flag.Args(), ",,,,,,,,,,,,<<<")
	http.HandleFunc("/", cmd)
	http.HandleFunc("/cmdDelete", cmdDeletion)
	http.HandleFunc(cmdCreateHandler, handlePost)
	http.HandleFunc(postHandler, handlePost)
	http.HandleFunc(storeHandler, handleStore)
	http.HandleFunc("/hello", hello)
	http.HandleFunc(cmdCreateHandler, cmdCreation)
	http.HandleFunc("/cmdList", cmdListing)
	http.HandleFunc("/count", count)
	http.HandleFunc("/cmd.*", cmd)
	//		http.HandleFunc("/exec", exec)
}
