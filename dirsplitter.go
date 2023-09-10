package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/jinyus/confirmop"
	"github.com/jinyus/dirsplitter/commands"
)

func main() {

	parser := argparse.NewParser("dirsplitter", "Split Directories")

	version := parser.NewCommand("version", "prints app version")

	split := parser.NewCommand("split", "Split a directory into parts of a given size")
	dirSplit := split.StringPositional(&argparse.Options{Help: "the directory to split", Default: "."})
	max := split.Float("m", "max", &argparse.Options{Help: "Size of each part in GB", Default: 5.0})

	reverse := parser.NewCommand("reverse", "Reverse a splitted directory")
	dirReverse := reverse.StringPositional(&argparse.Options{Help: "the directory to reverse", Default: "."})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if version.Happened() {
		fmt.Println("v1.0.0")
		os.Exit(0)
	}

	if split.Happened() {

		checkDir(dirSplit)

		ans := confirmop.ConfirmOperation(
			cfmt.Sprintf("Split \"{{%s}}::yellow\" into parts of %f GB?", *dirSplit, *max),
			"",
			true,
			true,
		)

		if !ans {
			os.Exit(0)
		}

		maxSize := *max * 1024 * 1024 * 1024

		commands.SplitDir(*dirSplit, int64(maxSize))

	} else if reverse.Happened() {

		checkDir(dirReverse)

		ans := confirmop.ConfirmOperation(
			cfmt.Sprintf("Reverse split \"{{%s}}::yellow\"?", *dirReverse),
			"",
			true,
			true,
		)

		if !ans {
			os.Exit(0)
		}

		commands.ReverseSplitDir(*dirReverse)
	}

	os.Exit(0)
}

func checkDir(dirSplit *string) {
	info, err := os.Stat(*dirSplit)
	if err != nil || !info.IsDir() {
		cfmt.Printf("{{Error:}}::red '%s' is not a directory\n", *dirSplit)
		os.Exit(1)
	}
}
