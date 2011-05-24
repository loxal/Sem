// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"fmt"
	"http"
	"io"
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


func handleMain(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" || r.URL.Path != indexHandler {
		serve404(w)
		return
	}

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
