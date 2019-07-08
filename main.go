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
	flag.IntVar(&lenPass, "l", DefaultLength, "Length of password")
	flag.Parse()

	if lenPass < 0 {
		lenPass = DefaultLength
	}

	fmt.Println(RandStringBytes(lenPass))
}
