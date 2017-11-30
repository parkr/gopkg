package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatalln("Specify exactly 1 command: 'list' or 'curr'")
	}

	command := flag.Args()[0]

	switch command {
	case "curr":
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("error retrieving cwd: %+v", err)
		}
		fmt.Println(packageNameFromSourcePath(dir))
	case "list":
		filepath.Walk(gosrc(), func(path string, info os.FileInfo, err error) error {
			if filepath.Base(path) == ".git" && info.IsDir() {
				fmt.Println(packageNameFromSourcePath(filepath.Dir(path)))
			}
			return nil
		})
	default:
		log.Fatalf("no such command: %q", command)
	}
}

func gopath() string {
	if userDefined := os.Getenv("GOPATH"); userDefined != "" {
		return userDefined
	}
	return os.Getenv("HOME") + "/go"
}

func gosrc() string {
	return gopath() + "/src"
}

func packageNameFromSourcePath(srcPath string) string {
	return strings.TrimPrefix(srcPath, gosrc()+"/")
}
