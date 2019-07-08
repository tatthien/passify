package main

import (
	"flag"
	"fmt"
)

const (
	DefaultLength int = 10
)

func main() {
	var lenPass int
	flag.IntVar(&lenPass, "l", DefaultLength, "length of password")
	containNumbers := flag.Bool("n", false, "password contains numbers (default false)")
	containSymbols := flag.Bool("s", false, "password contains symbols (default false)")
	flag.Parse()

	if lenPass < 0 {
		lenPass = DefaultLength
	}

	fmt.Println(RandStringBytes(lenPass, *containNumbers, *containSymbols))
}
