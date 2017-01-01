[![GoDoc](https://godoc.org/github.com/shuLhan/j2p?status.svg)]
(https://godoc.org/github.com/shuLhan/j2p)
[![Go Report Card](https://goreportcard.com/badge/github.com/shuLhan/j2p)]
(https://goreportcard.com/report/github.com/shuLhan/j2p)

`j2p` is a tool to help migrating from JIRA to Phabricator.

NOTE: due to limited resources on testing (time and JIRA server for testing)
this repository does not accept any reporting issue anymore on Github.


## Installation

* Get the repository
```
$ go get -u github.com/shuLhan/j2p
```
* Create `config` file (see `config.example` for an example of what you can set)

## Configuration

Configuration file is using JSON format. Below is decription for each key in
each object.

**`jira`**

This option is required to connect and query JIRA server.

* `url`: required, location of JIRA server
* `user`: required, user name
* `pass`: required, password for `user`
* `search_opt`: this option will set start and max result of query
   * `start_at`: set, any query will return search start at this index
   * `max_results`: maximum result that a query will return

**`phabricator`**

This option is required to connect and export to Phabricator server.

* `url`: required, location of Phabricator server
* `token`: required, conduit API token
* `secure`: if true, certificate of Phabricator server will be verified,
  otherwise certificate check will be skipped

**`prioritiesMap`**

This option define mapping of priorities between JIRA and Phabricator. The
key name is priority name in JIRA and their value is priority name in
Phabricator.

For example, to map `Highest` priority in JIRA to `Needs Triage` in
Phabricator, set key and value to `"Highest": 90`.

**`statutesMap`**

This option define mapping of statutes between JIRA and Phabricator. The key
name is status name in JIRA and their value is status name in Phabricator.
For example, to map `Open` status in JIRA to `open` status in Phabricator, add
`"Open": "open"`.

## Example

```
$ j2p tasks
```
  This will read the `config` file and export all projects and tasks from JIRA
  and import it to Phabricator.

```
$ j2p --projects="Project A,Project B" tasks
```
  This will only migrate tasks for project "Project A" and "Project B" only.

## Features

* Migrating tasks from JIRA to Phabricator

## Limitations

* Does not support migrating attachments
* Does not support creating sub-task

## Credits

* Thanks to Go-jira package by andygrunwald [1]

## License

Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
Use of this source code is governed by a BSD-style license that can be found
in the LICENSE file.

[1] https://github.com/andygrunwald/go-jira
