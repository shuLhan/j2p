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

const cmdUsage = `
Usage: j2p [--projects="Project A,Project B, ..."] <job>

job is one of this,

	tasks
		migrate JIRA tasks from all projects to Phabricator by creating
		new project and new Maniphest task.

`

var cmd j2p.Cmd

func init() {
	flag.Var(&cmd.Args.Projects, "projects",
		"List of project to be migrated, separated by comma.")
}

//
// usage will print this command usage to standard output.
//
func usage() {
	fmt.Println(cmdUsage)
}

//
// ParseArgs will parse user arguments from command line.
//
func ParseArgs() error {
	if j2p.DEBUG >= 2 {
		fmt.Printf("[j2p] Args: %s\n", cmd.Args)
	}

	cmd.Args.Job = flag.Arg(0)

	if cmd.Args.Job == "" || cmd.Args.Job != "tasks" {
		return errors.New("Empty or invalid command arguments")
	}

	return nil
}

func main() {
	flag.Parse()

	e := ParseArgs()
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
