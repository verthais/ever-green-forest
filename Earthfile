VERSION 0.6

docker:
    FROM golang:1.20.1
    COPY . /workspace
    WORKDIR /workspace
    SAVE IMAGE dev-swisscom-game:latest

# install:

test:
    FROM +docker
    RUN go test -v ./tests

check:
    FROM +test
    RUN go build -race  main.go
    RUN go build -asan  main.go

build:
    FROM +check
    RUN go build -buildvcs=true -mod=vendor main.go
    SAVE ARTIFACT main AS LOCAL main

package:
    FROM alpine:3.14
    COPY +build/main /main
    ENTRYPOINT /main
    SAVE IMAGE --push dimgray/ever-green-forest
