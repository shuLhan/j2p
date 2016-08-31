// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package j2p_test

import (
	"fmt"
	"testing"

	"github.com/shuLhan/j2p"
)

func TestNewConfig(t *testing.T) {
	cmd := j2p.Cmd{}

	e := cmd.NewConfig()

	if e != nil {
		t.Fatal(e)
	}

	fmt.Printf("CONFIG >> %+v\n", cmd.Config)
}
