/*
Install packages from package.json globally

"npm install -g" will install locally if we have packages.json
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type packagesFile struct {
	Dependencies map[string]string `json:"dependencies"`
}

func main() {
	showVersion := flag.Bool("version", false, "show version and exit")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [FILE] [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *showVersion {
		fmt.Println("0.1.0")
		os.Exit(0)
	}

	var inFile io.Reader

	switch flag.NArg() {
	case 0:
		inFile = os.Stdin
	case 1:
		var err error

		inFile, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprint(os.Stderr, "error: wrong number of arguments\n")
		os.Exit(1)
	}

	dec := json.NewDecoder(inFile)
	var pkg packagesFile
	if err := dec.Decode(&pkg); err != nil {
		fmt.Fprintf(os.Stderr, "error: can't parse file - %s\n", err)
		os.Exit(1)
	}

	count := 0
	for name, version := range pkg.Dependencies {
		dep := fmt.Sprintf("%s@%s", name, version)
		args := []string{"npm", "install", "-g", dep}
		fmt.Println(strings.Join(args, " "))

		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "error: can't install %s - %s\n", dep, err)
			os.Exit(1)
		}
		count++
	}

	fmt.Printf("Total of %d packages installed", count)
}
