
GO111MODULE			?= on
export GO111MODULE
GOPROXY				?= https://proxy.golang.org
export GOPROXY

BIN_DIR 			?= $(shell pwd)/tmp/bin
GOPATH            	?= $(shell go env GOPATH)
GOBIN             	?= ${GOPATH})/bin



.PYTHON: all

all: deps install-proto-gen build

.PHONY: deps
deps: go.mod go.sum
	@go mod vendor

.PHONY: install-proto-gen

install-proto-gen:
	@echo "install proto buffer generator tool and proto buffer go by go get "
	@go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	@export PATH="$PATH:$(go env GOPATH)/bin"

.PHONY: gen-go-protobuf
gen-go-protobuf:
	@protoc -I ./apis/railway/grpc/v1beta1 api.proto --go_out=./apis/railway/grpc/v1beta1/ --go_out=plugins=grpc:.

build:
	@echo "building transportation binary in $(BIN_DIR)"
	@#go build