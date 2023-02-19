package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

var (
	searchStr = flag.String(
		"searchStr",
		"",
		`string you want to search for in a given path, eg: -searchStr=fooBar`,
	)
	path = flag.String(
		"path",
		"",
		`path where you want to find the searchStr, eg: -path=/foo/bar`,
	)
)

func parseFlags() error {
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(flag.NFlag())
	fmt.Println(flag.NArg())
	if *searchStr == "" || *path == "" {
		return errors.New("args cannot be empty")
	}
	return nil
}

func main() {
	if err := parseFlags(); err != nil {
		log.Fatalf("Error parsing command line flags: %v", err)
	}
	fmt.Println(*searchStr, *path)
}
