# SPoNgEBoB
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcalini%2Fspongebob.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcalini%2Fspongebob?ref=badge_shield)

Simple project for converting normal text to sPonGeBOb tExT

![Mocking SpongeBob](spongebob.jpeg)

It uses _CUTTING EDGE™_ technology like *MARKOV CHAINS* to generate REALISTICⓇ SPonGeBoB text.️

## Getting it
```sh
go get -u github.com/calini/spongebob
```

## Using the CLI

```sh
spongebob "hello world"
> hELlO wOrlD
```

## Building it manually
```sh
go build -o ./spongebob

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


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcalini%2Fspongebob.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcalini%2Fspongebob?ref=badge_large)
