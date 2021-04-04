# GoMetadata

## Introduction 

This is a tool to automatically search the internet for 'official' metadata from the internet for audio files with sub-standard metadata.

## How to run

* go run cmd/gom/main.go -flags /options etc
* Alternatively use Makefile

## How to test

* go test internal/metadata/metadata_test.go
* go test internal/metadata
* Alternatively use Makefile via make test target, and append name of test file e.g. make test-metadata

## How to build

* go build cmd/gom/main.go%
* Alternatively use Makefile

## Contribute

* investigate returning different types to different modules in golang
* get tags
* search for tags
* update tags

## ßNotes

* https://github.com/wtolson/go-taglib
* https://godoc.org/github.com/wtolson/go-taglib
* https://github.com/barasher/go-exiftool
* https://godoc.org/github.com/bogem/id3v2
* https://github.com/mikkyang/id3-go

## Golang notes
* Function names starting with a capital letter can be used in external modules.
* := operator is a shortcut for declaring and initializing a variable in one line