# Overview

skipgen is a program that will generate a skiplist given a yaml file and
optionally a board name, branch name, and environment name.

[![Build Status](https://travis-ci.org/Linaro/skipgen.svg?branch=master)](https://travis-ci.org/Linaro/skipgen)

## Download and Install

Download release for your OS and architecture at
https://github.com/Linaro/skipgen/releases. Extract and run the 'skipfile'
binary.

## Usage

    skipgen [--board <boardname>] [--branch <branchname>] [--environment <environmentname] [--version] <skipfile.yaml>

## Example Usage

    $ skipgen examples/skipfile.yaml
    seccomp_bpf
    $ skipgen --board x15 --environment production --branch=4.4 examples/skipfile.yaml
    test_verifier
    test_tag
    test_maps
    test_lru_map
    test_lpm_map
    test_progs
    test_align
    ...

## Skipfile Format

See [examples/skipfile.yaml](examples/skipfile.yaml).

## Building

1. Install golang. i.e. on debian-based systems, run `apt-get install golang`.
2. Set GOPATH. See https://github.com/golang/go/wiki/SettingGOPATH.
3. Install go dependencies. `go get -t ./...`
4. install golint. `go get -u github.com/golang/lint/golint`
   Don't forget to setup the path PATH="$GOPATH/bin:$PATH"
5. make skipgen
6. `./skipgen`

## Development

Print usage:
`$ go run skipgen.go`

Get default skiplist:
`$ go run skipgen.go examples/skipfile.yaml`

Get board and environment-specific skiplist:
`$ go run skipgen.go --board x15 --environment production`examples/skipfile.yaml

## Testing

skipgen includes unit tests that can be run using `go test`. The `make test`
target will also run 'go vet' and 'golint'. golint may need to be installed
(`go get -u github.com/golang/lint/golint`)
