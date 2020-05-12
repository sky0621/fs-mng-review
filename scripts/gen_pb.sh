#!/usr/bin/env bash
cd ..
pwd
protoc -I schema/ schema/reviewer.proto --go_out=plugins=grpc:src/pb
