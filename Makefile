export GO111MODULE=on


.PHONY: deps
deps:
	@go mod vendor

.PHONY: install-proto-gen

install-proto-gen:
	@go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	@export PATH="$PATH:$(go env GOPATH)/bin"


.PHONY: gen-go-protobuf
gen-go-protobuf:
	@protoc -I ./IDL *.proto --go_out=./goprotobuf/submit --go_out=plugins=grpc:.

echo:
	echo "${ROOT}"