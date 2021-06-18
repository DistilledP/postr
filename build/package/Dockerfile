##########
FROM golang:latest as proto_builder

ENV PROTOC_VERSION=3.17.3
ENV PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip

RUN apt-get update && apt-get install -y unzip tree
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP} \
    && unzip -o ${PROTOC_ZIP} -d /usr/local bin/protoc \
    && unzip -o ${PROTOC_ZIP} -d /usr/local 'include/*' \
    && rm -f ${PROTOC_ZIP}

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

##########
FROM proto_builder as proto_compile

COPY . /app

WORKDIR /app

RUN /bin/sh ./scripts/compile_proto.sh

