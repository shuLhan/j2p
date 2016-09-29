// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package j2p

import (
	"encoding/json"
	"io/ioutil"
)

//
// ConfigJiraSearchOpt contains options when doing query to JIRA.
//
type ConfigJiraSearchOpt struct {
	StartAt    int `json:"start_at"`
	MaxResults int `json:"max_results"`
}

//
// ConfigJira contains information about JIRA server and API options.
//
type ConfigJira struct {
	URL       string              `json:"url"`
	User      string              `json:"user"`
	Pass      string              `json:"pass"`
	SearchOpt ConfigJiraSearchOpt `json:"search_opt"`
}

//
// ConfigPhabricator contains information of Phabricator server.
//
type ConfigPhabricator struct {
	URL    string `json:"url"`
	Token  string `json:"token"`
	Secure bool   `json:"secure"`
}

//
// Config contains user configuration, read from `config` file in current
// directory.
//
type Config struct {
	Jira          ConfigJira        `json:"jira"`
	Phab          ConfigPhabricator `json:"phabricator"`
	PrioritiesMap map[string]int    `json:"prioritiesMap"`
	StatusesMap   map[string]string `json:"statusesMap"`
}

//
// NewConfig will read `config` file and save them in Config object.
//
func (cmd *Cmd) NewConfig() (e error) {
	var bcfg []byte

	bcfg, e = ioutil.ReadFile("config")

	if e != nil {
		return e
	}

	cmd.Config = &Config{}

	e = json.Unmarshal(bcfg, &cmd.Config)

	return e
}
