#!/bin/sh

export VERSION=$GITHUB_REF_NAME
mkdir ./build

build() {
	GOOS=$1 GOARCH=$2 go build -o ./build/qs-forward || exit 1
	cd ./build
	chmod u+x ./qs-forward
	zip ./qs-forward_${VERSION}_$1-$2.zip ./qs-forward || exit 1
	tar -czf ./qs-forward_${VERSION}_$1-$2.tar.gz ./qs-forward || exit 1
	rm ./qs-forward
	cd ..
}

build linux amd64 || exit 1
build darwin amd64 || exit 1
build darwin arm64 || exit 1
build windows amd64 || exit 1
