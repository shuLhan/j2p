[![GoDoc](https://godoc.org/github.com/shuLhan/j2p?status.svg)]
(https://godoc.org/github.com/shuLhan/j2p)
[![Go Report Card](https://goreportcard.com/badge/github.com/shuLhan/j2p)]
(https://goreportcard.com/report/github.com/shuLhan/j2p)

`j2p` is a tool to help migrating from JIRA to Phabricator.

## Installation

* Get the repository
```
$ go get -u github.com/shuLhan/j2p
```
* Create `config` file (see `config.example` for an example of what you can set)

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
