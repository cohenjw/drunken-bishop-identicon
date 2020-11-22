package drunkenbishop

import "strings"

const (
	width  = 17
	height = 9
	mappings = " .o+=*BOX@%&#/^SE"
)

// Fingerprint is a represenation of a hash value using
// the Drunkey Bishop algorithm.
type Fingerprint [height][width]rune

// NewFingerprint returns a Fingerprint of the given
// hash value.
func NewFingerprint(hash []byte) (f *Fingerprint) {
	f = &Fingerprint{}
	f.genFingerprint(hash)
	f.applyMappings()
	return
}

// genFingerprint generates the fingerprint image,
// based on the given hash.
// (mappings not yet applied - only visit counts.)
func (f *Fingerprint) genFingerprint(hash []byte) {
	// Start in the center.
	x, y := width/2, height/2

	// Probably overly complicated - but works effectively.
	for _, b := range hash {
		for i := 0; i < 4; i++ {
			pair := (b >> (2 * i)) & 0b11

			x += 2 * int(pair & 0b01) - 1
			y += 2 * int((pair & 0b10) >> 1) - 1

			x, y = clampWithinBounds(x, y)

			f[y][x]++
		}
	}

	// Start and end positions.
	f[height / 2][width / 2] = 15
	f[y][x] = 16
}

// clampWithinBounds clamps  `x` and `y` to be within
// the bounds of the indices of the image
// (defined by the `width` and `height` constants).
func clampWithinBounds(x, y int) (int, int) {
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

	return x, y
}

// applyMappings applies the `mappings` to the image
// generated by the `genFingerprint` method.
func (f *Fingerprint) applyMappings() {
	for i, row := range f {
		for j, c := range row {
			f[i][j] = rune(mappings[c])
		}
	}
}

// String returns the formatted fingerprint image,
// encased in a border.
func (f *Fingerprint) String() string {
	var sb strings.Builder

	sb.WriteString("+")
	sb.WriteString(strings.Repeat("-", width))
	sb.WriteString("+\n")

	for _, row := range f {
		sb.WriteString("|")
		for _, c := range row {
			sb.WriteRune(c)
		}
		sb.WriteString("|\n")
	}
	
	sb.WriteString("+")
	sb.WriteString(strings.Repeat("-", width))
	sb.WriteString("+\n")

	return sb.String()
}
