.PHONY: build clean

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

# Binary output
BINARY_NAME=app

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) pod.go

docker:
	docker build -t jdnielss/pod .

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
