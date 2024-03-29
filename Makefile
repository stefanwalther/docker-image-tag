SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=

help:								## Show this help.
	@echo ''
	@echo 'Available commands:'
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ''
.PHONY: help

gen-readme:					## Generate README.md (using docker-verb)
	docker run --rm -v ${PWD}:/opt/verb stefanwalther/verb
.PHONY: gen-readme

clean:							## Clean up the release directory
	rm -rf ./dist
.PHONY: clean

# Run all the tests and opens the coverage report
cover: test
	go tool cover -html=coverage.txt
.PHONY: cover

release: clean			## Build a release
	goreleaser
	rm -rf ./dist
.PHONY: release

build:							## Build
	go build
.PHONY: build

build_release:
	./scripts/build_release.sh
.PHONY: build_release

ci: build test
.PHONY: ci

# Run all the tests
test:
	go test $(TEST_OPTIONS) -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m
.PHONY: test

.DEFAULT_GOAL := build
