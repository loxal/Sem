// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"fmt"
	"http"
	"io"
	"io/ioutil"
	"json"
	"os"
	"template"

	"appengine"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Hello, ...!\n")
}

func serveError(c appengine.Context, w http.ResponseWriter, err os.Error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Internal Server Error")
	c.Logf("%v", err)
}

func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Not Found")
}

func test() {
	type Test struct {
		Blub uint8
	}
	type Site struct {
		Copyright, Title, TitleDesc string
		Mail string
		Test *Test
	}
	var j string = `{
"Copyright": "Alexander Orlov. All rights reserved.", 
"Title": "Sem — Sem Entity Manager | Loxal", 
"TitleDesc": "Sem Entity Manager"
}`
	var s Site
	fmt.Println(s.Title)
	//fmt.Println(s.Copyright)
result, _ := ioutil.ReadFile("./site_properties.json")
err := json.Unmarshal([]byte(j), &s)
err2 := json.Unmarshal([]byte(result), &s)
fmt.Println(err)
fmt.Println(err2)
fmt.Println(string(result), "<<<<")
fmt.Println(s.Title)
fmt.Println(s.TitleDesc)
fmt.Println(s.Copyright)
fmt.Println(s.Mail)

}

func handleMain(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" || r.URL.Path != indexHandler {
		serve404(w)
		return
	}
//	test()
fmt.Println(w.Header(), "<<<<<")
        w.Header().Set("Content-Type", "text/html")
        if err := mainPresenter.Execute(w, nil); err != nil {
        }
}

const indexHandler = "/index"
const mainHandler = "/main"

var mainPresenter *template.Template

func init() {
mainPresenter = template.New(nil)
mainPresenter.SetDelims("{%", "%}")
if err := mainPresenter.ParseFile("main.html"); err != nil {
    panic("can't parse: " + err.String())
}
	http.HandleFunc(indexHandler, handleMain)
	http.HandleFunc(mainHandler, handleMain)
}
