#!/bin/bash
if command -v apt >/dev/null 2>&1; then
	sudo apt install golang
elif command -v dnf >/dev/null 2>&1; then
	sudo dnf install golang 
elif command -v pacman >dev/null 2>&1; then 
	sudo pacman -Sy go 
elif command -v zypper >dev/null 2>&1; then 
	sudo zypper install go 
elif command -v apk >dev/null 2>&1; then
	sudo apk add go
else
	unable to identify distro install golang manually 
fi
git clone https://github.com/josithg/go_checksum.git
cd go_checksum
go mod tidy 
go build -o main 
./main
