// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package j2p

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andygrunwald/go-jira"
	"github.com/shuLhan/gonduit"
)

//
// Args contain command parameter passed by user. This will define what kind
// of object that will be migrated to Phabricator.
//
type Args struct {
	Job      string
	Projects ProjectFlags
}

//
// Cmd contains user configuration, HTTP client for JIRA and Phabricator, and
// command parameters.
//
type Cmd struct {
	Config      *Config
	JiraCl      *jira.Client
	GonduitCl   *gonduit.Client
	Args        Args
	gonProjects []gonduit.Project
}

var (
	// DEBUG debug level, set using environment J2P_DEBUG
	DEBUG = 0
)

func init() {
	var e error
	DEBUG, e = strconv.Atoi(os.Getenv("J2P_DEBUG"))
	if e != nil {
		DEBUG = 0
	}
}

//
// Init will,
// (1) get JIRA and phabricator configuration,
// (2) login to JIRA,
// (3) create gonduit client.
//
func (cmd *Cmd) Init() (e error) {
	// (1)
	e = cmd.NewConfig()
	if e != nil {
		return e
	}

	fmt.Printf("CONFIG = %+v\n", cmd.Config)

	// (2)
	cmd.JiraCl, e = NewJiraClient(cmd.Config)
	if e != nil {
		return e
	}

	// (3)
	cmd.GonduitCl = gonduit.NewClient(cmd.Config.Phab.URL,
		cmd.Config.Phab.Token)

	return nil
}
