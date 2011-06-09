// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package site

import (
	"http"
	"template"
)

// TODO going to be the catch-all handler?
func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	reloadMainPresenterTemplate() // auto-reload / refresh in dev mode
	if err := mainPresenter.Execute(w, nil); err != nil {
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
	http.HandleFunc(indexHandler, handleMain)
}
