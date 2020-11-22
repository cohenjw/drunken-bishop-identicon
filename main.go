package main

import (
	"crypto/sha256"
)

func main() {
	hash := sha256.New()
	hash.Write([]byte("Hello World!"))
	fingerprint = NewFingerprint(hash.Sum())
}
