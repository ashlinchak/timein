#!/usr/bin/env bash

# windows
GOOS=windows GOARCH=386 go build -o tmp/windows/timein.exe
zip -j bin/windows.zip tmp/windows/timein.exe config.json data.json

# linux
GOOS=linux GOARCH=386 go build -o tmp/linux/timein
zip -j bin/linux.zip tmp/linux/timein config.json data.json

# macos
GOOS=darwin GOARCH=386 go build -o tmp/darwin/timein
zip -j bin/macos.zip tmp/darwin/timein config.json data.json

echo "Done"
