# About
A program that utilises the Drunken Bishop algorithm to produce a unique ASCII-art indentification icon (identicon), based on a hash.

You may have seen this used by OpenSSH, providing "randomart" for their keys.

# Installation
To install this package and start generating drunken bishop ASCII identicons, install Go and run `go get`:

```console
$ go get github.com/cohenjw/drunken-bishop-identicon
```

# Usage
```go
package main

import (
	"fmt"
	"crypto/sha256"
	
	// Consider using an alias (long identifier - drunkenbishop).
	dbi "github.com/cohenjw/drunken-bishop-identicon"
)

func main() {
	// Have a []byte hash available:
	sha := sha256.New()
	sha.Write([]byte("Hello World!"))
	hash := sha.Sum(nil)

	// Pass the hash to NewFingerprint to retrieve the Fingerprint - [][]rune
	fp :=  dbi.NewFingerprint(hash)

	// Consider utilising its String() method to format it nicely in the terminal:
	fmt.Println(fp)
}
```


Output:
```
+-----------------+
|        o*.. ..  |
|        Eo=. ..  |
|       .o+.o.. ..|
|       .  . o +oo|
|        S ..o*oB+|
|         ..*++B+B|
|         .+.+o..X|
|          .. ..+o|
|               .o|
+-----------------+
```

# Future
Maybe play around with making the output not ASCII, but something a bit more interesting, perhaps utilising graphics (via https://golang.org/pkg/image/ for example).
