#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <package>"
    exit 1
fi

echo "app/$1/main.go"
air --build.cmd "go build  -o tmp/$1 app/$1/main.go" --build.bin "export $(grep -v '^#' .env | xargs); ./tmp/$1"