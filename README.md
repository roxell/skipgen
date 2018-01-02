= Overview =

skipgen is a program that will generate a skiplist given a yaml file and
optionally a board name, branch name, and enviornment name.

== Usage ==

    $ skipgen -skipfile examples/skipfile.yaml 
    seccomp_bpf
    $ skipgen -skipfile examples/skipfile.yaml -board x15 -environment production
    test_verifier
    test_tag
    test_maps
    test_lru_map
    test_lpm_map
    test_progs
    test_align
    ...

== Skipfile Format ==

See examples/skipfile.yaml.

== Building ==

1. Install golang. i.e. on debian-based systems, run `apt-get install golang`.
2. Set GOPATH. See https://github.com/golang/go/wiki/SettingGOPATH.
3. Install go dependencies. `go get ./...`
4. go build
5. ./skipgen

== Development ==

`$ go run skipgen.go`
`$ go run skipgen.go -skipfile examples/skipfile.yaml`
`$ go run skipgen.go -skipfile examples/skipfile.yaml -board x15 -environment production`

== Testing ==

`make test`

