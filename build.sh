#!/bin/bash
echo "Start build macOS 32 bit"
GOOS=darwin GOARCH=386 go build -o build/macos/32/main
echo "Completed"
echo "Start build macOS 64 bit"
GOOS=darwin GOARCH=amd64 go build -o build/macos/64/main
echo  "Completed"
echo "Start build linux 32 bit"
GOOS=linux GOARCH=386 go build -o build/linux/32/main
echo "Completed"
echo "Start build linux 64 bit"
GOOS=linux GOARCH=amd64 go build -o build/linux/64/main
echo "Completed"
echo "Start build windows 32 bit"
GOOS=windows GOARCH=386 go build -o build/windows/32/main.exe
echo "Completed"
echo "Start build windows 64 bit"
GOOS=windows GOARCH=amd64 go build -o build/windows/64/main.exe
echo "Completed"
