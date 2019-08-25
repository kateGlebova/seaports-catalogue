#!/bin/bash
echo "=============== Building PortDomainService service! ==============="
GO111MODULE=on go mod download
GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -o /app/repository ./domain-service/main.go

ls -l /app/repository