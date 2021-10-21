#!/usr/bin/env bash

# list of platforms: go tool dist list
platforms=(
    "darwin/amd64"
    "darwin/arm64"
    "linux/386"
    "linux/amd64"
    "linux/arm"
    "linux/arm64"
)

mkdir -p build

for platform in "${platforms[@]}"; do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name='passager-'$GOOS'-'$GOARCH
    echo "Building ${output_name}..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o './build/'$output_name cmd/passager/main.go
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
