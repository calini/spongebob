package main

import (
	"bufio"
	"fmt"
	"github.com/calini/spongebob/spongebob"
	"os"
	"strings"
)

func main() {
	// if string passed as param, return them
	if len(os.Args) >= 2 {
		for i := range os.Args {
			os.Args[i] = spongebob.Text(os.Args[i])
		}
		print(strings.Join(os.Args[1:], " "))
		os.Exit(0)
	}

	// else, get stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(spongebob.Text(scanner.Text()))
	}
}
