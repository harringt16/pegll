// Generated by gogll. Do not edit.
//
//  Copyright 2019 Marius Ackerman
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Package slots generates a text file, grammar_slots.txt, containing the grammar slots.
package slots

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bruceiv/pegll/cfg"
	"github.com/bruceiv/pegll/gslot"
	"github.com/goccmack/goutil/ioutil"
)

func Gen(gs *gslot.GSlot) {
	if !cfg.Verbose {
		return
	}
	buf := new(bytes.Buffer)
	for _, s := range gs.Slots() {
		fmt.Fprintf(buf, "%s\n", s)
	}
	if err := ioutil.WriteFile(filepath.Join(cfg.BaseDir, "grammar_slots.txt"), buf.Bytes()); err != nil {
		fmt.Printf("Error writing grammar slots file: %s\n", err)
		os.Exit(1)
	}
}
