# Go parameters
GO111MODULE=on
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=goBash
RELEASE_NAME=goBash.zip
CONFIG_NAME=atom.json

all: build
build:
		mkdir -p bin release
		# export GOSUMDB=off
		GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)
		# $(GOBUILD) -o bin/$(BINARY_NAME)
		zip -j release/$(RELEASE_NAME) $(CONFIG_NAME) bin/$(BINARY_NAME)
clean:
		GOSUMDB=off $(GOCLEAN)
		rm -f bin/$(BINARY_NAME)
		rm -f release/$(RELEASE_NAME)
