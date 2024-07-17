CMD = go
PACKAGE_NAME = url-shortener
make_arch = $(shell uname -m)

.PHONY: build
build:
	GOARCH=$(make_arch) ${CMD} build -o ${PACKAGE_NAME}
