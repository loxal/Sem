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
var myFlag1 *string
var myFlag2 string
//flagSetPointer.StringVar(&myFlag, "g", "value of String", "usage of string")
//var myFlag *string = flag.String("g", "value of String", "usage of string")
//flag.StringVar(&myFlag, "nameOFString", "value of String", "usage of string")
//flag.Parse()
flagSetPointer := flag.NewFlagSet("google", flag.ContinueOnError)
//var myFlag *string = flagSetPointer.String("f", "v", "u")
//var myFlag = flagSetPointer.String("f", "", "u")
flagSetPointer.StringVar(&myFlag, "flag", "v", "usage")
myFlag1 = flagSetPointer.String("flag1", "v", "u")
flagSetPointer.StringVar(&myFlag2, "mon", "DEFAULT VALUE", "u")
//args:= []string{"g", "-g", "-g=g", "-g g", "-g u", *myFlag}
args:= []string{"-flag", "value", "-flag1", "flag1 Value", "-f", "vom"}
flagSetPointer.Usage = func() {
    fmt.Fprintln(w, "[MY USAGE]")
}
//otherArgs := flagSetPointer.Args()
 f:=flagSetPointer.Lookup("flag")
 fmt.Fprintln(w, f.Usage)

  flagSetPointer.PrintDefaults()
  if err := flagSetPointer.Parse(args); err != nil {
            fmt.Fprintf(w, " [MY ERROR] <br/> %v", err)
//            fmt.Fprintf(w, " error <br> %v", &myFlag2.Usage)
//            return
    }
fmt.Fprint(w, " BAL ")

//    fmt.Fprintf(w, "Arg: %v ", flagSetPointer.NArg());
//    fmt.Fprintf(w, "Arg: %q ", flagSetPointer.Arg(0));
//    fmt.Fprintf(w, "Arg: %q ", flagSetPointer.Arg(1));
    fmt.Fprintf(w, "Arg: %q ", myFlag);
    fmt.Fprintf(w, "Arg: %q ", *myFlag1);
    fmt.Fprintf(w, "Arg: %q ", myFlag2);
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




