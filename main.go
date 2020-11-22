package main

import (
	"fmt"
	"crypto/sha256"
	"drunken-bishop-identicon/drunkenbishop"
)

func main() {
	hash := sha256.New()
	hash.Write([]byte("Hello World!"))
	fingerprint := drunkenbishop.NewFingerprint(hash.Sum(nil))
	fmt.Println(fingerprint.String())
}
