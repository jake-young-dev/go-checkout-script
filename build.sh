#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o ./builds/linux-amd64-build
GOOS=windows GOARCH=amd64 go build -o ./builds/windows-amd64-build