package main

import (
	"os"

	"github.com/sayden/go-useful-scripts/git"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		panic("You must provide 2 arguments: user and email in this order")
	}

	thisPath := "."
	git.AddLocalUser(&args[0], &args[1], &thisPath)
}
