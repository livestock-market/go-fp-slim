# go-fp-slim

[![Build & Publish](https://github.com/livestock-market/go-fp-slim/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/livestock-market/go-fp-slim/actions/workflows/build.yaml)

If you like functional programming style, but [fp-go](https://github.com/IBM/fp-go) seems bloated to you, then this library might be for you.
This is a slim, light weight and minimal functional programming utility for Go lang.

## Installation

```shell
go get github.com/livestock-market/go-fp-slim
```

## Usage

Checkout the [examples](./docs/examples.md) for more details.


## Features
1. Either
2. Option
3. Map function
4. Filter function

Instead of using builtin `error.Error`, this library uses a custom `XError`, with `Code` and `Message` fields.
This is useful when you want to return error codes along with the error message.

## Contributing
Feel free to open issues and PRs.

## Testing
```shell
#for the first time only
go mod tidy  

go test ./...
```

## License
MIT License for the main library, but BSD 3-Clause License for the `XError`, since it uses [XErrors package](https://pkg.go.dev/golang.org/x/xerrors).

