package main

import (
	"fmt"
	"crypto/sha256"
	"drunken-bishop-identicon/dbish"
)

func main() {
	hash := sha256.New()
	hash.Write([]byte("Hello World!"))
	fingerprint := dbish.NewFingerprint(hash.Sum(nil))
	fmt.Println(fingerprint.String())
}
