// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()

	e := ParseArgs()
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	e = cmd.Init()

	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestProjectFlags(t *testing.T) {
	fmt.Println("Test cmd arguments:", cmd.Args)
}

func TestJiraGetProjects(t *testing.T) {
	fmt.Println("Test cmd arguments:", cmd.Args)

	_, e := cmd.JiraGetProjects()
	if e != nil {
		t.Fatal(e)
	}
}
