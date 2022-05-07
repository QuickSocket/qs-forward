#!/bin/sh

export VERSION=$GITHUB_REF_NAME
mkdir ./build

build() {
	EXTENSION=""
	if [ "$1" = "windows" ]; then
		EXTENSION=".exe"
	fi

	GOOS=$1 GOARCH=$2 go build -o ./build/qs-forward$EXTENSION || exit 1
	cd ./build

	if [ "$1" != "windows" ]; then
		chmod u+x ./qs-forward
	fi
	
	zip ./qs-forward_${VERSION}_$1-$2.zip ./qs-forward$EXTENSION || exit 1
	tar -czf ./qs-forward_${VERSION}_$1-$2.tar.gz ./qs-forward$EXTENSION || exit 1
	rm ./qs-forward$EXTENSION
	cd ..
}

build linux amd64 || exit 1
build darwin amd64 || exit 1
build darwin arm64 || exit 1
build windows amd64 || exit 1
