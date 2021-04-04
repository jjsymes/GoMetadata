# GoMetadata

## How to run

- go run cmd/gom/main.go -flags /options etc
- Alternatively use Makefile

## How to test

- go test internal/metadata/metadata_test.go
- go test internal/metadata
- Alternatively use Makefile via make test target, and append name of test file e.g. make test-metadata

## How to build

- go build cmd/gom/main.go%
- Alternatively use Makefile


# Golang notes
- Function names starting with a capital letter can be used in external modules.
- := operator is a shortcut for declaring and initializing a variable in one line