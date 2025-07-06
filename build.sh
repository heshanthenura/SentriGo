#!/bin/bash

# Exit on error
set -e

# Check if version argument is provided
if [ -z "$1" ]; then
    echo "[ERROR] No version tag provided."
    echo "Usage: ./build.sh <version-tag>"
    exit 1
fi

# Delete old build folder
rm -rf ./build/linux

# Recreate build folder
mkdir -p build/linux

# Show argument
echo "Building: sentrigo-linux-$1"

# Set environment for Linux build (use GOOS and GOARCH)
GOOS=linux GOARCH=amd64 go build -o build/linux/sentrigo-linux-$1 cmd/sentrigo/main.go
