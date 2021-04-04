.PHONY: default

default: # default target
	cat Makefile

run: # run code
	go run cmd/gom/main.go $(ARGS)

test-%: # run tests
	go test internal/$*

test-main: # run tests
	go test cmd/gom/main_test.go

build: # build binary
	cd cmd/gom && go build cmd/gom/main.go

git-mirror: # mirror the repo to git hub
	git push --mirror https://github.com/jjsymes/GoMetadata.git