// skipgen is a program that will generate a skiplist given a yaml file
// and optionally, board name, branch name, and environment name.

package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"flag"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// StringArray is a custom type that can be used for parsing
// yaml fields that can be either a string, or an array.
// See https://github.com/go-yaml/yaml/issues/100
type StringArray []string
// UnmarshalYAML is a custom version of UnmarshalYAML that implements
// StringArray. See https://github.com/go-yaml/yaml/issues/100
func (a *StringArray) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var multi []string
	err := unmarshal(&multi)
	if err != nil {
		var single string
		err := unmarshal(&single)
		if err != nil {
			return err
		}
		*a = []string{single}
	} else {
		*a = multi
	}
	return nil
}

// stringInSlice searches for a particular string
// in a slice of strings. If "all" is contained in
// the slice, then true is always returned.
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a || b == "all" {
			return true
		}
	}
	return false
}

// Skipfile is a map of the structure in the yaml skipfile
type Skipfile struct {
	Skiplist []struct {
		Reason string
		URL string
		Environments StringArray
		Boards StringArray
		Branches StringArray
		Tests []string
	}
}

// parseSkipfile parses a given buf and returns a Skipfile
// struct and err, if any.
func parseSkipfile(buf []byte) (Skipfile, error){
	var skips Skipfile
	err := yaml.Unmarshal(buf, &skips)
	return skips, err
}

// getSkipfileContents returns a string containing a skipfile
// given a board, environment, branch, and Skipfile struct.
func getSkipfileContents(board string, branch string, environment string, skips Skipfile) (string){
	var buf string
	buf = ""
	for _, skip := range skips.Skiplist {
		if stringInSlice(board, skip.Boards) &&
		   stringInSlice(branch, skip.Branches) &&
		   stringInSlice(environment, skip.Environments) {
			for _, test := range skip.Tests {
				buf = buf + fmt.Sprintf("%s\n", test)
			}
		}
	}
	return buf
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n    %s [--board <boardname>] [--branch <branchname>] [--environment <environmentname] <skipfile.yaml> \n", os.Args[0])
	os.Exit(1)
}

func main() {

	// This variable is overwritten during build by goreleaser
	var version = "master"

	boardPtr := flag.String("board", "all", "Board name. If not specified, skips that apply to all boards will be returned.")
	branchPtr := flag.String("branch", "all", "Branch name. If not specified, skips that apply to all branches will be returned.")
	environmentPtr := flag.String("environment", "all", "Environment name. If not specified, skips that apply to all environments will be returned.")
	versionPtr := flag.Bool("version", false, "Print skipgen version and exit.")
	flag.Parse()

	if (*versionPtr) {
		fmt.Fprintf(os.Stderr, "skipgen %s\n", version)
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "Error: skipfile not provided\n\n")
		usage()
	}
	if len(flag.Args()) > 1 {
		fmt.Fprintf(os.Stderr, "Error: unprocessed argument: '%s'\n", flag.Arg(1))
		fmt.Fprintf(os.Stderr, "Put the path to skipfile last on the command line.\n\n")
		usage()
	}

	skipfile := flag.Arg(0)

	_, err := os.Stat(skipfile)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: skipfile '%s' not found\n\n", skipfile)
		usage()
	}
	check(err)

	// Read skipfile.yaml
	buf, err := ioutil.ReadFile(skipfile)
	check(err)

	// Parse skipfile
	skips, err := parseSkipfile(buf)
	check(err)

	fmt.Printf(getSkipfileContents(*boardPtr, *branchPtr, *environmentPtr, skips))

}
