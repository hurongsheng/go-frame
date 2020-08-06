#!/bin/bash

protoc --go_out=plugins=grpc:. ./api/*.proto
cd grpc &&go build &&cd ..
cd http &&go build &&cd ..
./http/http&
./grpc/grpc :8081 1&
./grpc/grpc :8082 1&
./grpc/grpc :8083 1&