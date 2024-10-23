#!/bin/bash

CGO_ENABLED=1 GOARCH=arm64 go build -o watchme main.go
sudo chmod +x watchme
sudo mv ./watchme /usr/local/bin