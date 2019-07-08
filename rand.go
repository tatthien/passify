package main

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "0123456789"
const symbolBytes = "!@#$%^&*()-_=+?[]{}"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringBytes(n int, containNumbers, containSymbols bool) string {
	characterBytes := letterBytes

	if containNumbers {
		characterBytes += numberBytes
	}
	if containSymbols {
		characterBytes += symbolBytes
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = characterBytes[rand.Intn(len(characterBytes))]
	}
	return string(b)
}
