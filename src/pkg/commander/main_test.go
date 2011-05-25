// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
	"http"

	"appengine"
	"appengine/datastore"
//	"appengine/memcache"
//	"appengine/user"
)

type doubleTest struct {
	in, out int
}

var doubleTests = []doubleTest{
	doubleTest{1, 2},
	doubleTest{2, 4},
	doubleTest{-5, -10},
}

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

var cmdTests = []cmdTest {
	cmdTest{"c", "https://mail.google.com/mail/?shva=1#compose", "Compose Gmail"},
	cmdTest{"t", "http://twitter.com", "Twitter"},
	cmdTest{"sem", "https://github.com/loxal/Sem", "GitHub: Sem Project"},
	cmdTest{"verp", "https://github.com/loxal/Verp", "GitHub: Verp Project"},
	cmdTest{"lox", "https://github.com/loxal/Lox", "GitHub: Lox Project"},
}

//var cmdTests1 = []Cmd{
//	Cmd{"c", "https://mail.google.com/mail/?shva=1#compose", "Compose Gmail"},
//	Cmd{"t", "http://twitter.com", "Twitter"},
//	Cmd{"sem", "https://github.com/loxal/Sem", "GitHub: Sem Project"},
//	Cmd{"verp", "https://github.com/loxal/Verp", "GitHub: Verp Project"},
//	Cmd{"lox", "https://github.com/loxal/Lox", "GitHub: Lox Project"},
//}

func TestCmdCreation(t *testing.T) {

    cmd := &Cmd {
        Name: "blub",
        RESTcall: "blab",
        Desc: "....",
    }
//    cmd := &cmdTest {
//        name: "blub",
//        restCall: "blab",
//        desc: "....",
//    }

//    c := &appengine.NewContext(r)
//    headers := make(http.Header)
//    headers.Set("X-Appengine-Inbound-Appid", "my-app-id")
//    c := appengine.NewContext(&http.Request{Header: headers})
    c := appengine.NewContext(&http.Request{Header: make(http.Header)})
    c.Logf("%#v", c)
    c.Logf("%#v", cmd)
    datastore.Put(c, datastore.NewIncompleteKey("Cmd"), cmd)

//    for _, &ct:=range cmdTests {
////        v:=cmdCreation(ct.)
////        var r http.Request
//    r := &http.Request
//        cmdCreate(r, ct)
//    }
}

func TestDouble(t *testing.T) {
	for _, dt := range doubleTests {
		v := Double(dt.in)
		if v != dt.out {
			t.Errorf("Double(%d) = %d, want %d.", dt.in, v, dt.out)
		}
	}
}

func TestWebCmd(t *testing.T) {
//	for _, wct := range webCmdTests {
//		v := WebCmd(wct.cmd)
//		if v != wct.restCall {
//			t.Errorf("%s ==> %q != %q.", wct.cmd, v, wct.restCall)
//		}
//	}
}
