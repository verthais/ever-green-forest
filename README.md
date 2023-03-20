# Ever Green Forest

## Building and Running locally

```console
docker run --rm -v ${path}/ever-green-forest:/workspace golang:1.20.1 \
    /bin/bash -c "cd /workspace && go build main.go"
${path}/ever-green-forest/main
```

## Running tests

```console
docker run --rm -v ${path}/ever-green-forest:/workspace golang:1.20.1 \
    /bin/bash -c "cd /workspace && go test -v ./tests"
```

## Running the game

```console
docker run --rm dimgray/ever-green-forest
```