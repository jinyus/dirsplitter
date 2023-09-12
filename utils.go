package main

import (
	"fmt"
	"os"

	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/jinyus/confirmop"
)

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
