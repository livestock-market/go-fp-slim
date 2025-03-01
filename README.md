# go-fp-slim

![Build](https://github.com/livestock-market/go-fp-slim/actions/workflows/build.yaml/badge.svg)
If you like functional programming style, but fp_go seems bloated to you, then this library might be for you.
This is a Slim or light weight or minimal functional programming utility for Go lang.

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
go mod tidy // for the first time only
go test ./...
```

## License
MIT License for the main library, but BSD 3-Clause License for the `XError`, since it uses [XErrors package](https://pkg.go.dev/golang.org/x/xerrors).

