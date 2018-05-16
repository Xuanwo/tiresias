SHELL := /bin/bash

.PHONY: all check format　vet lint build install uninstall release clean test coverage

VERSION=$(shell cat ./constants/version.go | grep "Version\ =" | sed -e s/^.*\ //g | sed -e s/\"//g)
DIRS_TO_CHECK=$(shell ls -d */ | grep -vE "vendor|test")
PKGS_TO_CHECK=$(shell go list ./... | grep -v "/vendor/")

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check      to format, vet and lint "
	@echo "  build      to create bin directory and build tiresias"
	@echo "  install    to install tiresias to /usr/local/bin/tiresias"
	@echo "  uninstall  to uninstall tiresias"
	@echo "  release    to release tiresias"
	@echo "  clean      to clean build and test files"
	@echo "  test       to run test"
	@echo "  coverage   to test with coverage"

check: format vet lint

format:
	@echo "go fmt, skipping vendor packages"
	@for pkg in ${PKGS_TO_CHECK}; do go fmt $${pkg}; done;
	@echo "ok"

vet:
	@echo "go tool vet, skipping vendor packages"
	@go tool vet -all ${DIRS_TO_CHECK}
	@echo "ok"

lint:
	@echo "golint, skipping vendor packages"
	@lint=$$(for pkg in ${PKGS_TO_CHECK}; do golint $${pkg}; done); \
	 lint=$$(echo "$${lint}"); \
	 if [[ -n $${lint} ]]; then echo "$${lint}"; exit 1; fi
	@echo "ok"

build: check
	@echo "build tiresias"
	@mkdir -p ./bin
	@go build -o ./bin/tiresias .
	@echo "ok"

install: build
	@echo "install tiresias to GOPATH"
	@cp ./bin/tiresias ${GOPATH}/bin/tiresias
	@echo "ok"

uninstall:
	@echo "delete /usr/local/bin/tiresias"
	@rm -f /usr/local/bin/tiresias
	@echo "ok"

release:
	@echo "release tiresias"
	@mkdir -p ./release

	@echo "build for linux"
	@GOOS=linux GOARCH=amd64 go build -o ./bin/linux/tiresias_v${VERSION}_linux_amd64 .
	@tar -C ./bin/linux/ -czf ./release/tiresias_v${VERSION}_linux_amd64.tar.gz tiresias_v${VERSION}_linux_amd64

	@echo "ok"

clean:
	@rm -rf ./bin
	@rm -rf ./release
	@rm -rf ./coverage

test:
	@echo "run test"
	@go test -v ${PKGS_TO_CHECK}
	@echo "ok"

coverage:
	@echo "run test with coverage"
	@for pkg in ${PKGS_TO_CHECK}; do \
		output="coverage$${pkg#github.com/yunify/tiresias}"; \
		mkdir -p $${output}; \
		go test -v -cover -coverprofile="$${output}/profile.out" $${pkg}; \
		if [[ -e "$${output}/profile.out" ]]; then \
			go tool cover -html="$${output}/profile.out" -o "$${output}/profile.html"; \
		fi; \
	done
	@echo "ok"