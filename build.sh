#!/bin/sh

export VERSION=$GITHUB_REF_NAME
mkdir ./build

build() {
	GOOS=$1 GOARCH=$2 go build -o ./build/qs-forward || exit 1
	chmod u+x ./build/qs-forward
	zip ./build/qs-forward_${VERSION}_$1-$2.zip ./build/qs-forward || exit 1
	tar -czf ./build/qs-forward_${VERSION}_$1-$2.tar.gz ./build/qs-forward || exit 1
	rm ./build/qs-forward
}

build linux amd64 || exit 1
build darwin amd64 || exit 1
build darwin arm64 || exit 1
build windows amd64 || exit 1
