// Copyright 2010 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"solox/core"
	"flag"
	"fmt"
)

func main() {
	fmt.Printf("DOMAIN")
	core.Echo(flag.Arg(0))
	//fileStem := core.GetFileStem(flag.Arg(0))
	//fmt.Printf("%q\n", fileStem)
}
