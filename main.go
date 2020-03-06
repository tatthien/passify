package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

const (
	DefaultLength int = 16
)

func main() {
	var lenPass int
	flag.IntVar(&lenPass, "l", DefaultLength, "length of password")
	copyPassword := flag.Bool("c", false, "copy password to clipboard (default false)")
	flag.Parse()

	if lenPass < 0 {
		lenPass = DefaultLength
	}

	password := RandStringBytes(lenPass)

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

		fmt.Println(password)
		fmt.Println("> Copied to clipboard")
	} else {
		fmt.Print(password)
	}
}
