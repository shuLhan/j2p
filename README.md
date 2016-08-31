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

* Run `j2p`
```
$ j2p tasks
```
  This will export all projects and tasks from JIRA and import it to Phabricator.

## Features

* Migrating tasks from JIRA to Phabricator

## Limitations

* Does not support migrating attachments
* Does not support creating sub-task
