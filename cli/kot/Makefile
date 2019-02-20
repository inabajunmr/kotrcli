ARCH := 386 amd64
OS := linux darwin windows

setup:
	go get github.com/mitchellh/gox

build:
	GO111MODULE=on go build

build-all: 
	GO111MODULE=on gox -os="$(OS)" -arch="$(ARCH)" -output "./dist/{{.Dir}}_{{.OS}}_{{.Arch}}" 
