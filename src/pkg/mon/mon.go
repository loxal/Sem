// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mon

import (
    "flag" // replace by a post release.58.1 version and check whether flag.Init() exists
           //   then test whether flag.Init("name", 0) works
	"fmt"
	"http"
	"math"
)

func Test1(w http.ResponseWriter) (func(int) int) {
    fmt.Fprintf(w, "Hello From MON\n<br>")
    fmt.Fprintf(w, "POST-TEXTYPE-MON")
    var x int
    return func(delta int) int {
        x += delta
        return x
    }
}

func testFlag(w http.ResponseWriter){
//    var ip *int = flag.Int("flagname", 1234, "help message for flagname")
//    var test flag.FlagSet
    flag.NewFlagSet("bodommm", 0)
    var flagvar int
    flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
//    fmt.Fprintln(w, "ip has value ", *ip);
    fmt.Fprintln(w, "flagvar has valu ", flagvar);
}

type Point struct { x, y float64 }
// A method on *Point
func (p *Point) Abs() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y)
}




