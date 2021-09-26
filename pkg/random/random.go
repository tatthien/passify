package random

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "0123456789"
const symbolBytes = "!@#$%&"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetStringBytes generates random string from character bytes
func GetStringBytes(strLen int) string {
	characterBytes := letterBytes + numberBytes + symbolBytes

	b := make([]byte, strLen)

	for i := range b {
		b[i] = characterBytes[rand.Intn(len(characterBytes))]
	}

	return string(b)
}
