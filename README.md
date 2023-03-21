# Ever Green Forest

Once upon a time there was a great hero, called Natelus, with some strengths and
weaknesses, as all heroes have. After battling all kinds of monsters for more
than a hundred years, Natelus now you are able to witness his heroci deeds.

You'll have an oportunity to spectate simulate a battle between Natelus and a
wild beast at command line.

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

## Developing with VS Code

Frist build docker image locally

```console
docker build -t swisscome:0.0.2 -f ${path}/ever-green-forest/environment/Dockerfile .
```

Run VS Code and `Open Folder ever-green-forest` and `Reopen In Container`.
Make sure you have `ms-vscode-remote.remote-containers` extension.