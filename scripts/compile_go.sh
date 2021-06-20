#!/usr/bin/env bash

FLAG_CGO=0

if [ -z "${BUILD_ARTIFACTS}" ]; then
    echo "Specify BUILD_ARTIFACTS environment var"
    exit 1

fi

go mod vendor

echo "Running unit tests\n"
make run-test
echo ""

for app in $(ls -d cmd/*); do
    appName=$(echo ${app} | grep -oE "[^/]+$")
    echo "Compiling app: ${appName} ..."

    inputFile="${app}/*.go"
    outputFile="${BUILD_ARTIFACTS}/${appName}"

    CGO_ENABLED=${FLAG_CGO} \
        go build \
        -ldflags "-s -w" \
        -o ${outputFile} \
         ${inputFile}

    # Compress the executable if upx is installed
    if [ -x "$(command -v upx)" ]; then
        upx --best --lzma ${BUILD_ARTIFACTS}/${appName}
    fi

    echo "Success, ${appName} compiled\n"
done
