test:
	go test -v -cover
	go vet
	golint

