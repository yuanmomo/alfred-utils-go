BINARY = alfred-utils-go
GOARCH = amd64

VERSION?=?
BUILD=`date +%FT%T%z`

# Symlink into GOPATH
CURRENT_DIR=$(shell pwd)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

# Build the project
all: clean release

darwin:
	cd ${CURRENT_DIR}; \
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o build/${BINARY}-darwin-${GOARCH} . ;

clean:
	-rm -fr build
	-rm -fr release

release: darwin
	-mkdir -p release
	-cp build/${BINARY}-darwin-${GOARCH} release
	-upx --best release/${BINARY}-darwin-${GOARCH}

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make release - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make help - 查看帮助文档"

.PHONY: darwin clean release help