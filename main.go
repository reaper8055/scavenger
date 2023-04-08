package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
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
	if *searchStr == "" && *path != "" {
		return errors.New("searchStr cannott be empty!")
	}
	if *path == "" && *searchStr != "" {
		return errors.New("path cannot be empty!")
	}
	return nil
}

func main() {
	if err := parseFlags(); err != nil {
		log.Fatalf("Error parsing command line flags: %v", err)
	}
	fmt.Println(*searchStr, *path)
	getAllFiles()
}

func getAllFiles() (string, error) {
	f, err := os.Open(*path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fInfo, err := f.Stat()
	if err != nil {
		return "", err
	}
	if fInfo.IsDir() {
		dirEntry, err := f.ReadDir(-1)
		if err != nil {
			return "", err
		}
		fmt.Println(len(dirEntry))
		return *path, nil
	}
	return *path, nil
}
