// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package j2p

import (
	"fmt"
	"strings"
)

//
// ProjectFlags contains list of project passed through command line.
//
type ProjectFlags []string

//
// String return text representation of project flags.
//
func (prflags *ProjectFlags) String() string {
	return fmt.Sprint(*prflags)
}

//
// Set will parse the projects flags.
//
func (prflags *ProjectFlags) Set(vv string) error {
	for _, v := range strings.Split(vv, ",") {
		if v == "" {
			continue
		}
		*prflags = append(*prflags, strings.TrimSpace(v))
	}

	return nil
}
