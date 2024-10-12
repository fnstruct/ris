// Copyright (C) 2024 fnstruct. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	glob := fmt.Sprintf("*.%s", os.Args[1])
	path := filepath.Join(os.Args[2], glob)

	files, err := filepath.Glob(path)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(files); i++ {
		filename := fmt.Sprintf("%0*d.%s", 4, i + 1, os.Args[1])
		newfile := filepath.Join(os.Args[2], filename)

		_, err := os.Stat(newfile)
		err = os.Rename(files[i], newfile)
		if err != nil {
			panic(err)
		}
	}
}
