.PHONY: build clean docker

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

# Binary output
BINARY_NAME=pod

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) main.go

docker:
	docker build -t jdnielss/pod .

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
