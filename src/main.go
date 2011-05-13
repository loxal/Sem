// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"http"
	"io"
//	"io/ioutil"
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
	Title	string
//	Body	[]byte
	Body	string
}

type Page struct {
	Title	string
//	Body	[]byte
	Body	string
}

var mainPage = template.MustParseFile("template.html", nil)
//const lenPath = len("/view/")

//func loadPage(title string) (*Page, os.Error) {
//	filename := title + ".txt"
//	body, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return nil, err
//	}
//	return &Page{Title: title, Body: body}, nil
//}

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
	if r.Method != "GET" || r.URL.Path != postHandler {
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
	////////////////
//	title := "my"
//	p, err := loadPage(title)
//	if err != nil {
//		p = &Page{Title: title}
//	}
//	t, _ := template.ParseFile("template.html", nil)
//	mainPage.Execute(w, p)

//    return &Page{Title: title, Body: body}, nil
	////////////////
//	p := &Page{Title: "ddd", Body: "oops"}
 for i := 0; i < len(gg); i++ {
//     gg[i]= &Greeting{Title: "my TITLE", Body: "my BODY"}
 }
	w.Header().Set("Content-Type", "text/html")
//	mainPage.Execute(w, p)
	if err := mainPage.Execute(w, gg); err != nil {
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
	http.Redirect(w, r, postHandler, http.StatusFound)
}

var postHandler = "/post"
var storeHandler = "/store"
var blub = "blubber"

func init() {
	http.HandleFunc("/", hello)
	http.HandleFunc(postHandler, handlePost)
	http.HandleFunc(storeHandler, handleStore)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/count", count)
}