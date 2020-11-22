package main

import "strings"

const (
	width  = 17
	height = 9
)

// Fingerprint is a represenation of a hash value using
// the Drunkey Bishop algorithm. Use `NewFingerprint` to
// instantiate this struct.
type Fingerprint struct {
	hash  []byte
	image [height][width]rune
}

// NewFingerprint returns an instance of fingerprint,
// containing both the Drunken Bishop identicon, and
// the hash value used to produce
func NewFingerprint(h []byte) (f *Fingerprint) {
	f = &Fingerprint{hash: h}
	f.genFingerprint()
	f.applyMappings()
	return
}

func (f *Fingerprint) genFingerprint() {
	var image [height][width]rune

	// Start in the center.
	x, y := width/2, height/2

	// Probably overly complicated - but works.
	for _, b := range f.hash {
		for i := 0; i < 4; i++ {
			pair := (b >> (2 * i)) & 0b11

			x += 2*int(pair&0b01) - 1
			y += 2*int((pair&0b10)>>1) - 1

			x, y = clampWithinBounds(x, y)

			image[y][x]++
		}
	}

	image[4][9] = 15
	image[y][x] = 16

	f.image = image
}

func clampWithinBounds(x, y int) (int, int) {
	var x1, y1 int

	if x < 0 {
		x = 0
	} else if x > width-1 {
		x = width - 1
	}

	if y < 0 {
		y = 0
	} else if y > height-1 {
		y = height - 1
	}

	return x1, y1
}

func (f *Fingerprint) applyMappings() {
	mappings := map[uint8]rune{
		0: ' ', 1: '.', 2: 'o', 3: '+',
		4: '=', 5: '*', 6: 'B', 7: 'O',
		8: 'X', 9: '@', 10: '%', 11: '&',
		12: '#', 13: '/', 14: '^', 15: 'S',
		16: 'E',
	}

	var mappedImage [height][width]rune

	for i, row := range f.image {
		for j, c := range row {
			mappedImage[i][j] = mappings[uint8(c)]
		}
	}

	f.image = mappedImage
}

func (f *Fingerprint) String() string {
	var sb strings.Builder

	sb.WriteString("+----------------+")
	for _, row := range f.image {
		sb.WriteRune('|')
		for _, c := range row {
			sb.WriteRune(c)
		}
		sb.WriteRune('|')
	}
	sb.WriteString("+----------------+")

	return sb.String()
}
