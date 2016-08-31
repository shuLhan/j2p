// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/shuLhan/j2p"
)

func usage() {
	fmt.Println(`
Usage: j2p [job]

job is one of this,

	tasks
		migrate JIRA tasks from all projects to Phabricator by creating
		new project and new Maniphest task.
`)
}

//
// parseArgs will parse user arguments from command line.
//
func parseArgs(cmd *j2p.Cmd) error {
	flag.Parse()

	cmd.Args.Job = flag.Arg(0)

	if cmd.Args.Job == "" || cmd.Args.Job != "tasks" {
		return errors.New("Empty or invalid command arguments")
	}

	return nil
}

func main() {
	var e error
	var cmd j2p.Cmd

	e = parseArgs(&cmd)
	if e != nil {
		usage()
		os.Exit(1)
	}

	e = cmd.Init()
	if e != nil {
		panic(e)
	}

	switch cmd.Args.Job {
	case "tasks":
		e = cmd.MigrateTasks()
		if e != nil {
			panic(e)
		}
	}

}
