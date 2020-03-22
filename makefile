NAME := a-thing-tells-global-ip-on-telegram
VERSION := v1.0.0
REVISION := $(shell git rev-parse --short HEAD)

SRCS:=$(shell find . -type f -name '*.go')
LDFLAGS:= -ldflags="-s -w -X \"main.version=$(VERSION)\" -X \"main.revision=$(REVISION)\" -extldflags \"-static\""

bin/$(NAME): $(SRCS)
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

.PHONY: clean
clean:
	rm -rf bin/*