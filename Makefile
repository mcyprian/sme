SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET=$(shell echo $${PWD\#\#*/})

VERSION=1.0.0
BUILD=`git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: clean install fmt deps

all: build

build: fmt
	@go build ${LDFLAGS} -o ${TARGET}

install:
	@godep go install ${LDFLAGS}

fmt:
	@gofmt -l -w $(SRC)

deps:
	@godep save

clean:
	if [ -f ${TARGET} ] ; then rm ${TARGET} ; fi


