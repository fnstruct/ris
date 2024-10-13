// Copyright (C) 2024 fnstruct. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"os"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/itzg/go-flagsfiller"
)

type Options struct {
	Name     string `usage:"Specify which files to rename"`
	Index    int    `usage:"Provide custom index" default:"1"`
	Version  bool   `usage:"Print installed version"`
	Verbose  bool   `usage:"Print output when renaming"`
	Zfill    int    `usage:"Provide custom zfill" default:"4"`
}

var (
	opts Options
	Version string
)

func main() {
	flag.Usage = func() {
		msg := fmt.Sprintf("Usage: %s [OPTIONS] PATH\n", os.Args[0])
		fmt.Fprintf(os.Stderr, msg)
		flag.PrintDefaults()
	}

	err := flagsfiller.Parse(&opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if opts.Version {
		fmt.Println(os.Args[0], Version)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(0)
	}

	glob := fmt.Sprintf("*.%s", opts.Name)
	path := filepath.Join(args[0], glob)

	files, err := filepath.Glob(path)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(files); i++ {
		index := fmt.Sprintf("%0*d", opts.Zfill, i + opts.Index)
		filename := fmt.Sprintf("%s.%s", index, opts.Name)

		newfile := filepath.Join(args[0], filename)
		_, err := os.Stat(newfile)

		err = os.Rename(files[i], newfile)
		if err != nil {
			panic(err)
		}

		if opts.Verbose {
			fmt.Println(files[i], "=>", newfile)
		}
	}
}
