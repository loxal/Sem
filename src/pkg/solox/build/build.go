// Copyright 2010 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package build

import (
    "exec"
    "flag"
    "fmt"
    "os"
    "regexp"
)

const GolangFileExtension = ".go"
const UserHomeDir = "/Users/alex"
const Compiler = UserHomeDir + "/my/dev/env/go/bin/6g"
const Linker =  UserHomeDir + "/my/dev/env/go/bin/6l"
const BinDir = UserHomeDir + "/my/dev/prj/loxal/Solox/bin/"
const WorkDir = "./"

func compile(fileStem string) os.Error {
    cmd, err := exec.Run(Compiler, []string{Compiler, "-o" + BinDir + fileStem + ".6", flag.Arg(0)}, os.Environ(), WorkDir,
        exec.DevNull, exec.DevNull, exec.MergeWithStdout)
    if err != nil {
        return err
    }

    return cmd.Close()
}

func link(fileStem string) os.Error {
    cmd, err := exec.Run(Linker, []string{Linker, "-o" + fileStem + ".out", fileStem + ".6"}, os.Environ(), BinDir,
        exec.DevNull, exec.DevNull, exec.MergeWithStdout)
    if err != nil {
        return err
    }

    return cmd.Close()
}

func clean(fileStem string) os.Error {
	rmCmd, err := exec.LookPath("rm")
	cmd, err := exec.Run(rmCmd, []string{rmCmd, fileStem + ".6"}, os.Environ(), BinDir,
		exec.DevNull, exec.PassThrough, exec.MergeWithStdout)
	if err != nil {
		return err
	}

	return cmd.Close()
}

func GetFileStem(fileName string) string {
    var pattern *regexp.Regexp = regexp.MustCompile("^(.*)\\" + GolangFileExtension + "$")

    return pattern.FindStringSubmatch(fileName)[1]
}

func Main() {
	fmt.Printf("args: %v\n", flag.NArg())
	fileStem := GetFileStem(flag.Arg(0))
	compile(fileStem)
	link(fileStem)
	clean(fileStem)
}
