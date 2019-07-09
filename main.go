package main

import (
	"flag"
	"fmt"
	"runtime"
	"log"
	"os/exec"
)

const (
	DefaultLength int = 10
)

func main() {
	var lenPass int
	flag.IntVar(&lenPass, "l", DefaultLength, "length of password")
	containNumbers := flag.Bool("n", false, "password contains numbers (default false)")
	containSymbols := flag.Bool("s", false, "password contains symbols (default false)")
	copyPassword := flag.Bool("c", false, "copy password to clipboard (default false)")
	flag.Parse()

	if lenPass < 0 {
		lenPass = DefaultLength
	}

	password := RandStringBytes(lenPass, *containNumbers, *containSymbols)

	if *copyPassword {
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

		fmt.Println("Copied to clipboard")
	} else {
		fmt.Print(password)
	}
}
