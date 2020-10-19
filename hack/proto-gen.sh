#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail
ROOT=$(dirname "${BASH_SOURCE[0]}")/..

function protoc::generate_proto() {
  kube::golang::setup_env
  go install k8s.io/kubernetes/vendor/k8s.io/code-generator/cmd/go-to-protobuf/protoc-gen-gogo

  kube::protoc::check_protoc

  local package=${1}
  protoc::protoc "${package}"
  protoc::format "${package}"
}

function kube::protoc::protoc() {
  local package=${1}

  PATH="${gogopath}:${PATH}" protoc \
    --proto_path="${package}" \
    --proto_path="${KUBE_ROOT}/vendor" \
    --gogo_out=plugins=grpc:"${package}" "${package}/api.proto"
}

function kube::protoc::format() {
  local package=${1}

  # Update boilerplate for the generated file.
  cat hack/boilerplate/boilerplate.generatego.txt "${package}/api.pb.go" >tmpfile && mv tmpfile "${package}/api.pb.go"

  # Run gofmt to clean up the generated code.
  kube::golang::verify_go_version
  gofmt -l -s -w "${package}/api.pb.go"
}
