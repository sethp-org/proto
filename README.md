
### Prerequisites

- **[Go][]**, any one of the **three latest major** [releases of Go][].

For installation instructions, see Go's [Getting Started][] guide.

- **[Protocol buffer][pb] compiler**, `protoc`, [version 3][proto3].

For installation instructions, see [Protocol Buffer Compiler
Installation][pbc-install].

- **Go plugins** for the protocol compiler:

1. Install the protocol compiler plugins for Go using the following commands:

```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

2. Update your `PATH` so that the `protoc` compiler can find the plugins:

```sh
    export PATH="$PATH:$(go env GOPATH)/bin"
```

## Первый запуск в проекте

`git submodule add https://github.com/sethp-org/proto`
`git submodule update --remote`
