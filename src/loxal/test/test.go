// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
    "flag" // replace by a post release.58.1 version and check whether flag.Init() exists
           //   then test whether flag.Init("name", 0) works
	"fmt"
	"net/http"
	"strings"
	"math"
)

func TestFlag(w http.ResponseWriter){
    //    var test flag.FlagSet
    var myFlag string
    var myFlag1 *string
    var myFlag2 string
    flagSetPointer := flag.NewFlagSet("google", flag.ContinueOnError)
    flagSetPointer.StringVar(&myFlag, "flag", "DEFAULT VALUE", "usage")
    myFlag1 = flagSetPointer.String("flag1", "DEFAULT VALUE", "usage")
    flagSetPointer.StringVar(&myFlag2, "flag2", "DEFAULT VALUE 2", "usage")
    args:= []string{"-flag", "value", "-flag1", "flag1 Value", "-f", "vom"}
    flagSetPointer.Usage = func() {
        fmt.Fprintln(w, "[MY USAGE]")
    }
    otherArgs := flagSetPointer.Args()
     f:=flagSetPointer.Lookup("flag")
     fmt.Fprintln(w, f.Usage)

      flagSetPointer.PrintDefaults()
      if err := flagSetPointer.Parse(args); err != nil {
                fmt.Fprintf(w, " [MY ERROR] <br/> %v", err)
    //            fmt.Fprintf(w, " error <br> %v", &myFlag2.Usage)
    //            return
        }
    fmt.Fprint(w, " BAL ")

    fmt.Fprintf(w, "flag: %q ", myFlag);
    fmt.Fprintf(w, "flag1: %q ", *myFlag1);
    fmt.Fprintf(w, "flag2: %q ", myFlag2);
    fmt.Fprintf(w, "Other: %v ", otherArgs);
}

func ParseQuery(query string) string {
    const sep = " "
    queryCmd := strings.Split(query, sep)

    var taskCmd string
    fs := flag.NewFlagSet("", flag.ContinueOnError)
    fs.StringVar(&taskCmd, "add", "my Default Task", "ADD A NEW TASK USAGE")

    if err:=fs.Parse(queryCmd[1:]); err != nil {
        panic("boom!!!")
    }

     return taskCmd
}

type Point struct { x, y float64 }
// A method on *Point
func (p *Point) Abs() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y)
}

func ReflectFunc(w http.ResponseWriter, r *http.Request) {
    // call as e.g. ?content=myContent&contentType=application/json
    content := r.FormValue("content")
    contentType := r.FormValue("contentType")

    w.Header().Set("Content-Type", contentType)
    fmt.Fprintf(w, content)
}



