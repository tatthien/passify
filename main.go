package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/tatthien/passify/pkg/random"
)

const (
	// DefaultLength is the default password length
	DefaultLength int = 16
)

func main() {
	var lenPass int
	flag.IntVar(&lenPass, "l", DefaultLength, "length of password")
	flag.Parse()

	if lenPass < 0 {
		lenPass = DefaultLength
	}

	password := random.GetStringBytes(lenPass)

	arch := runtime.GOOS
	var copyCmd *exec.Cmd

	// Mac OS
	if arch == "darwin" {
		copyCmd = exec.Command("pbcopy")
	}

	// Linux
	if arch == "linux" {
		copyCmd = exec.Command("xclip")
	}

	in, err := copyCmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := in.Write([]byte(password)); err != nil {
		log.Fatal(err)
	}

	if err := copyCmd.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(">>> Copied to clipboard")
	fmt.Println(password)
}
