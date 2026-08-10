# SPoNgEBoB

[![CI](https://github.com/calini/spongebob/actions/workflows/ci.yml/badge.svg)](https://github.com/calini/spongebob/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/calini/spongebob)](https://github.com/calini/spongebob/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/calini/spongebob)](https://goreportcard.com/report/github.com/calini/spongebob)
[![License](https://img.shields.io/github/license/calini/spongebob)](LICENSE.md)

Simple project for converting normal text to sPonGeBOb tExT

<p align="center">
  <img src="spongebob.jpg" alt="Mocking SpongeBob" width="300">
</p>

It uses _CUTTING EDGE_ technology like *MARKOV CHAINS™* to generate _REALISTIC_ SPonGeBoBⓇ text.️

## Getting it

### Homebrew

```sh
brew install calini/tap/spongebob
```

### Go

```sh
go install github.com/calini/spongebob/cmd/spongebob@latest
```

## Using the CLI

```sh
spongebob "hello world"
> hELlO wOrlD
```

## Building it manually
```sh
go build -o ./spongebob ./cmd/spongebob

./spongebob "hello world"
> hELlO wOrlD
```

## Using it as a library
```go
package main

import (
	"fmt"

	"github.com/calini/spongebob"
)

func main() {
    fmt.Println(spongebob.Text("Hello world!"))
}
```
