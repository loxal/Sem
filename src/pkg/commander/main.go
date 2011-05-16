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

	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"appengine/user"
)

type Greeting struct {
	Author  string
	Content string
	Date    datastore.Time

	Title string
	Body  string
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
	q := datastore.NewQuery("Greeting").Order("-Date").Limit(10)
	var gg []*Greeting
	_, err := q.GetAll(c, &gg)
	if err != nil {
		serveError(c, w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := mainPage.Execute(w, gg); err != nil {
		c.Logf("%v", err)
	}

	for i := 0; i < len(gg); i++ {
        gg[i]= &Greeting{Title: "my TITLE", Body: "my BODY"}
	}
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

func cmd(w http.ResponseWriter, r *http.Request) {
	//    c := appengine.NewContext(r)
	//    c.Logf("r.URL.Path: " + r.URL.Path)
	//    c.Logf("r.FormValue(\"foo\"): " + r.FormValue("foo"))
	//    c.Logf(r.FormValue("bar"))
	//    c.Logf("r.URL.RawQuery: " + r.URL.RawQuery)
	//
	//     c.Logf("m[r.URL.RawQuery]" + m[r.URL.RawQuery])
	//    http.Redirect(w, r, m[r.URL.RawQuery], http.StatusFound)
}

// Returns the RESTful associated with a certain command
func WebCmd(cmd string) (restCall string) {
	m := map[string]string {
		"c":   "https://mail.google.com/mail/?shva=1#compose",
		"t":   "http://twitter.com",
		"sem": "https://github.com/loxal/Sem",
		"verp": "https://github.com/loxal/Verp",
		"lox": "https://github.com/loxal/Lox",
		// shortcut for adding an English Word or another unknow word to the TO_LEARN_LIST (merge with the Delingo functionality)
		// shortcut for making notes/tasks/todos
	}

	restCall = m[cmd]
	return
}

var cmdCreateHandler = "/cmdCreate"
//var postHandler = "/post"
var storeHandler = "/store"
var createCmdPresenter = template.MustParseFile("cmdCreate.html", nil)
var mainPage = template.MustParseFile("template.html", nil)

func Double(i int) int {
	return i * 2
}

func init() {
	http.HandleFunc("/", hello)
	http.HandleFunc(cmdCreateHandler, handlePost)
//	http.HandleFunc(postHandler, handlePost)
	http.HandleFunc(storeHandler, handleStore)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/count", count)
	http.HandleFunc("/cmd", cmd)
}
