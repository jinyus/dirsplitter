package main

import (
	"fmt"
	"os"

	"github.com/i582/cfmt/cmd/cfmt"
	cli "github.com/jawher/mow.cli"
	"github.com/jinyus/confirmop"
	"github.com/jinyus/dirsplitter/commands"
)

func main() {

	app := cli.App("dirsplitter", "Split Directories")

	app.Version("v version", "v1.0.0")

	app.Command("split", "Split a directory into parts of a given size", func(cmd *cli.Cmd) {

		dir := cmd.StringArg("DIR", ".", "the directory to split")
		max := cmd.Float64Opt("m max", 5.0, "Size of each part in GB")

		cmd.Action = func() {

			checkDirAndConfirmOp(dir, cfmt.Sprintf("Split \"{{%s}}::yellow\" into parts of %f GB?", *dir, *max))

			maxSize := *max * 1024 * 1024 * 1024

			commands.SplitDir(*dir, int64(maxSize))
		}
	})

	app.Command("reverse", "Reverse a splitted directory", func(cmd *cli.Cmd) {
		dir := cmd.StringArg("DIR", ".", "the directory to reverse")

		cmd.Action = func() {

			checkDirAndConfirmOp(dir, cfmt.Sprintf("Reverse split \"{{%s}}::yellow\"?", *dir))

			commands.ReverseSplitDir(*dir)
		}
	})

	app.Run(os.Args)
}

func checkDir(dir *string) {
	info, err := os.Stat(*dir)
	if err != nil || !info.IsDir() {
		cfmt.Printf("{{Error:}}::red '%s' is not a directory\n", *dir)
		os.Exit(1)
	}
}

func checkDirAndConfirmOp(dir *string, desc string) {
	checkDir(dir)

	ans := confirmop.ConfirmOperation(
		desc,
		"",
		true,
		true,
	)

	if !ans {
		fmt.Println("Operation cancelled")
		os.Exit(0)
	}
}
