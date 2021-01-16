#!/usr/bin/env bash

set -e

GOOGLE_API=$1
GRPC_GATEWAY=$2
PROTO_DIR=$3

[ "_${GOOGLE_API}${GRPC_GATEWAY}${PROTO_DIR}" = "_" ] && echo "Usage: build.sh GO_GOOGLE_API_DIR GO_GRPC_GATEWAY_DIR PROTO_DIR(absolute path)" && exit 0

echo "Prep filesystem"
rm -rf ./build
mkdir build

echo "Starting build.sh ${GOOGLE_API} ${GRPC_GATEWAY} ${PROTO_DIR}"

protoc --version

echo "Domain..."

PROTOBUF="$PROTO_DIR/domain/*.proto"
protoc -I. -I$PROTO_DIR -I$GOOGLE_API --go_out=plugins=grpc:./build $PROTOBUF
protoc -I. -I$PROTO_DIR -I$GOOGLE_API --grpc-gateway_out=logtostderr=true:./build $PROTOBUF

echo "API..."
PROTOBUF="$PROTO_DIR/api/*.proto"
#protoc -I. -I$PROTO_DIR -I$GOOGLE_API --go_out=plugins=grpc:./build $PROTOBUF
#protoc -I. -I$PROTO_DIR -I$GOOGLE_API --grpc-gateway_out=logtostderr=true:./build $PROTOBUF
protoc -I. -I$PROTO_DIR -I$GOOGLE_API -I$GRPC_GATEWAY --go_out=plugins=grpc,Mdomain/domain.proto=github.com/yagacc/go-sea-port/domain/domain:./build $PROTOBUF
protoc -I. -I$PROTO_DIR -I$GOOGLE_API -I$GRPC_GATEWAY --grpc-gateway_out=logtostderr=true:./build $PROTOBUF

echo "Release"
rm -rf ./api
mv build/api .
rm -rf ./domain
mv build/domain .
rm -rf ./build

echo "Finished Build"
