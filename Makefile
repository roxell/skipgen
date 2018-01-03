FILE=skipgen

all: help

help:
	@echo "make <target>"
	@echo "targets:"
	@echo "     $(FILE)"
	@echo "     test"
	@echo "     clean"

$(FILE): test
	go build

test:
	go test -v -cover
	go vet
	golint

clean:
	rm -f ./$(FILE)
