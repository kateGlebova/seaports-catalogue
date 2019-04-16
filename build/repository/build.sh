#!/bin/bash
echo "=============== Building ClientAPI service! ==============="
GO111MODULE=on go mod vendor
GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -o /app/repository ./cmd/repository/main.go

ls -l /app/repository