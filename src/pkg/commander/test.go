// Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commander

import (
//	"fmt"
	"http"
	"math"
)

func test(w http.ResponseWriter) (func(int) int) {
    w.Header().Set("Content-Type", contentTypeText)
    var x int
    return func(delta int) int {
        x += delta
        return x
    }
}

type Point struct { x, y float64 }
// A method on *Point
func (p *Point) Abs() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y)
}




