version := $(shell egrep -o '[0-9]+\.[0-9]+\.[0-9]+' main.go)
os = $(shell go env GOOS)
arch = $(shell go env GOARCH)
out = npm-install-global

all:
	$(error pick a target)

test:
	go test -v

build:
	GOOS=darwin go build
	gzip -c $(out) > nig-$(version)-darwin-$(arch).gz
	go build
	gzip -c $(out) > nig-$(version)-$(os)-$(arch).gz
