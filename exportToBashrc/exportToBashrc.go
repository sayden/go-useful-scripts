package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("You must use two arguments: environment variable name " +
		"and route (both without spaces)")
	}

	s := mountString(args)
	bf := fmt.Sprintf("/%s/.bashrc", os.Getenv("HOME"))
	appendToFile(s, &bf)

	fmt.Println("Exported successfully")
	fmt.Println("Use \"source ~/.bashrc\" to add it to your current shell")
}

func mountString(args []string) *string{
	path := args[0]
	route := args[1]

	path = strings.ToUpper(path)

	s := fmt.Sprintf("export %s=%s\n", path, route)

	return &s
}

func appendToFile(s *string, fn *string) {
	f := openFile(*fn)
	defer f.Close()

	if _, err := f.WriteString(*s); err != nil {
		panic(err)
	}
}

func openFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	return f
}
