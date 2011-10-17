// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commander

import (
	"testing"
	"http"

	"appengine"
	"appengine/datastore"
)

type webCmdTest struct {
	cmd, restCall string
}

var webCmdTests = []webCmdTest{
	webCmdTest{"c", "https://mail.google.com/mail/?shva=1#compose"},
	webCmdTest{"t", "http://twitter.com"},
	webCmdTest{"sem", "https://github.com/loxal/Sem"},
	webCmdTest{"verp", "https://github.com/loxal/Verp"},
	webCmdTest{"lox", "https://github.com/loxal/Lox"},
}

type cmdTest struct {
	name, restCall, desc string
}

var cmdTests = []cmdTest{
	cmdTest{"c", "https://mail.google.com/mail/?shva=1#compose", "Compose Gmail"},
	cmdTest{"t", "http://twitter.com", "Twitter"},
	cmdTest{"sem", "https://github.com/loxal/Sem", "GitHub: Sem Project"},
	cmdTest{"verp", "https://github.com/loxal/Verp", "GitHub: Verp Project"},
	cmdTest{"lox", "https://github.com/loxal/Lox", "GitHub: Lox Project"},
}

func TestCmdCreation(t *testing.T) {

	cmd := &Cmd{
		Name:     "blub",
		RESTcall: "blab",
		Desc:     "....",
	}

	c := appengine.NewContext(&http.Request{Header: make(http.Header)})
	c.Logf("%#v", c)
	c.Logf("%#v", cmd)
	datastore.Put(c, datastore.NewIncompleteKey(c, "Cmd", nil), cmd)
}
