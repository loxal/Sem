// Copyright 2010 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package build

import "testing"

type GetFileStemTest struct {
	in, out string
}

var GetFileStemTests = []GetFileStemTest{
	GetFileStemTest{"myFile.go", "myFile"},
	GetFileStemTest{"mySecondFILE.go", "mySecondFILE"},
}

func TestGetFileStem(t *testing.T) {
	for _, r := range GetFileStemTests {
		o := GetFileStem(r.in)
	        if r.out != o {
		        t.Errorf("%s expected %c got %q", r.in, r.out, o)
		}
	}
}

func BenchmarkGetFileStem(b *testing.B) {
	s := "myBenchFile.go"
	for i := 0; i < b.N; i++ {
		GetFileStem(s)
	}
}
