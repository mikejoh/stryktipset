# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=st
BINARY_UNIX=$(BINARY_NAME)_unix

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

build-linux:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

build-docker-linux:
	docker run -v ${PWD}:/st -w /st -e GOOS=linux -e GOARCH=amd64 golang:latest go get -d -v && go build -v

build-docker-darwin:
	docker run -v ${PWD}:/st -w /st -e GOOS=darwin -e GOARCH=amd64 golang:latest go get -d -v && go build -v

