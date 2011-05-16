// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
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

func TestDouble(t *testing.T) {
	for _, dt := range doubleTests {
		v := Double(dt.in)
		if v != dt.out {
			t.Errorf("Double(%d) = %d, want %d.", dt.in, v, dt.out)
		}
	}
}

func TestWebCmd(t *testing.T) {
	for _, wct := range webCmdTests {
		v := WebCmd(wct.cmd)
		if v != wct.restCall {
			t.Errorf("%s ==> %q != %q.", wct.cmd, v, wct.restCall)
		}
	}
}
