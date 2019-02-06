ARCH := 386 amd64
OS := linux darwin windows

setup:
	go get github.com/mitchellh/gox

build-all: 
	cd cli/kot && GO111MODULE=on gox -os="$(OS)" -arch="$(ARCH)" -output "./dist/{{.Dir}}_{{.OS}}_{{.Arch}}"