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
)

func serveError(w http.ResponseWriter, err os.Error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", plainTxtEnc)
	io.WriteString(w, "Internal Server Error")
	fmt.Fprintln(w, "%v", err)
}

func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", plainTxtEnc)
	io.WriteString(w, "Not Found")
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
	    panic("TESTING 'panic'")
		serve404(w)
		return
	}
    type Site struct {
		Author string "author" // this tag can be omitted and is used for demo reasons only
		Copyright, Title, TitleDesc, Mail string
		Year string "year"
	}
	content, _ := ioutil.ReadFile("pkg/site/properties.json")
	var site Site
	json.Unmarshal([]byte(content), &site)
    w.Header().Set("Content-Type", "text/html")
    reloadMainPresenterTemplate(); // auto-reload / refresh in dev mode
    if err := mainPresenter.Execute(w, &site); err != nil {
    }
}

func reloadMainPresenterTemplate() {
    mainPresenter = template.New(nil)
	mainPresenter.SetDelims("{{", "}}")
    	if err := mainPresenter.ParseFile(mainPresenterSite); err != nil {
    		panic("can't parse: " + err.String())
	}
}

const indexHandler = "/"
const plainTxtEnc = "text/plain; charset=utf-8"
const mainPresenterSite = "pkg/site/main.html"
var mainPresenter *template.Template

func init() {
    reloadMainPresenterTemplate()
	http.HandleFunc(indexHandler, handleMain)
}
