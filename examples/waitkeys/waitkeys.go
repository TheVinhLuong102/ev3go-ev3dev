// Copyright ©2016 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// waitkeys demonstrates key waiting. It should be run from the command line.
// It requires ^C to terminate.
package main

import (
	"fmt"
	"log"

	"github.com/ev3go/ev3dev"
)

func main() {
	w, err := ev3dev.NewButtonWaiter()
	if err != nil {
		log.Fatalf("failed to create button waiter: %v", err)
	}

	for e := range w.Events {
		fmt.Printf("%+v\n", e)
	}
}
