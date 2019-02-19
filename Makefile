#!/usr/bin/env nix-shell
#!nix-shell -p go ragel -i ./.make

.PHONY: dev binary dockerun image test regen run clean deps depunpin depupdate

## Variables

GENERATED_FILES := pkg/parsing/parser.go pkg/parsing/lexer.go
SOURCE_FILES := $(shell find . -name '*.go' -type f)
GO_FILES := $(SOURCE_FILES) $(GENERATED_FILES)
PROJ_DIR := $(realpath .)
GOPATH := $(shell echo "${GOPATH}" | cut -d':' -f1 )
GOYACC := $(GOPATH)/bin/goyacc
GORICH := $(GOPATH)/bin/richgo

## Phony Targets

binary: $(GOPATH)/bin/chell

dev:
	find . -type f | grep -E '(\.(go|mod|sum|rl|y)|Makefile)$$' | entr -d make

deps: go.sum $(GOYACC)
	go mod download

depunpin:
	rm -f go.sum

depupdate: depunpin deps

run: binary
	$(GOPATH)/bin/chell

clean:
	rm -f $(GENERATED_FILES)

regen: clean $(GENERATED_FILES)

image:
	docker build --force-rm --pull --tag nekroze/chell:latest .

dockerun: $(GO_FILES)
	docker run -it --rm nekroze/chell:latest

test: $(GO_FILES)
	docker build --force-rm --pull --tag nekroze/chell:tests --target test .
	docker run -it --rm nekroze/chell:tests

## Targets

$(GOYACC): go.sum
	go get golang.org/x/tools/cmd/goyacc

$(GORICH):
	env GO111MODULE=auto go get github.com/kyoh86/richgo

pkg/parsing/parser.go: pkg/parsing/parser.y $(GOYACC)
	cd $(dir $@) && $(GOYACC) -o $(notdir $@ $<)
	rm -f y.output
	goimports -w $@

pkg/parsing/lexer.go: pkg/parsing/lexer.rl
	cd $(dir $@) && ragel -Z -G2 -o $(notdir $@ $<)
	goimports -w $@

go.sum: go.mod
	go get -d -u -v ./...
	touch $@ $<

$(GOPATH)/bin/chell: deps $(GENERATED_FILES) main.go pkg/*/*.go
	go install .
