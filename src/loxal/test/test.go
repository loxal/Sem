// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
    "flag" // replace by a post release.58.1 version and check whether flag.Init() exists
           //   then test whether flag.Init("name", 0) works
	"fmt"
	"http"
	"math"
)

func test1(w http.ResponseWriter) (func(int) int) {
    fmt.Fprintf(w, "Hello From MON\n<br>")
    fmt.Fprintf(w, "POST-TEXTYPE-MON")
    var x int
    return func(delta int) int {
        x += delta
        return x
    }
}

func TestFlag(w http.ResponseWriter){
//    var test flag.FlagSet
var myFlag string
//flagSetPointer.StringVar(&myFlag, "g", "value of String", "usage of string")
//var myFlag *string = flag.String("g", "value of String", "usage of string")
//flag.StringVar(&myFlag, "nameOFString", "value of String", "usage of string")
//flag.Parse()
flagSetPointer := flag.NewFlagSet("", flag.ContinueOnError)
//var myFlag *string = flagSetPointer.String("f", "v", "u")
//var myFlag = flagSetPointer.String("f", "", "u")
flagSetPointer.StringVar(&myFlag, "f", "v", "u")
//args:= []string{"g", "-g", "-g=g", "-g g", "-g u", *myFlag}
args:= []string{"-f", "FEST"}
//flagSetPointer.Usage = func() {}
//otherArgs := flagSetPointer.Args()


  if err := flagSetPointer.Parse(args); err != nil {
            fmt.Fprintf(w, " error <br> %v", err)
//            return
    }
fmt.Fprint(w, " test ")

//    fmt.Fprintf(w, "Arg: %v ", flagSetPointer.NArg());
//    fmt.Fprintf(w, "Arg: %q ", flagSetPointer.Arg(0));
//    fmt.Fprintf(w, "Arg: %q ", flagSetPointer.Arg(1));
    fmt.Fprintf(w, "Arg: %q ", myFlag);
//    fmt.Fprintf(w, "Arg: %q ", flagSetPointer.Arg(3));
//    fmt.Fprintf(w, "Arg: %q ", flagSetPointer.Arg(4));
//    fmt.Fprintf(w, "Arg: %q ", flagSetPointer.Arg(5));
//    fmt.Fprintf(w, "Arg: %q ", flagSetPointer.Arg(6));
//    fmt.Fprintf(w, "Other: %v ", otherArgs);
}

type Point struct { x, y float64 }
// A method on *Point
func (p *Point) Abs() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y)
}




