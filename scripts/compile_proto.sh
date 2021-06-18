#!/usr/bin/env bash

SCHEMA_DIRECTORY=proto
OUTPUT_DIRECTORY=${PWD}/internal/proto

FILES=$(find ${SCHEMA_DIRECTORY} -name *.proto | xargs)

SCHEMA_FILES=""
for f in ${FILES} ; do
    f2=$(echo $f | sed 's/proto\///g')
    SCHEMA_FILES="${SCHEMA_FILES} ${f2}"
done

if [ ! -d ${OUTPUT_DIRECTORY} ]; then
    mkdir -p ${OUTPUT_DIRECTORY}
fi;

protoc --proto_path=${SCHEMA_DIRECTORY} \
       --go_out=${OUTPUT_DIRECTORY} \
       --go_opt=paths=source_relative \
       --go-grpc_out=${OUTPUT_DIRECTORY} \
       --go-grpc_opt=paths=source_relative \
       ${SCHEMA_FILES}
