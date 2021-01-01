package spongebob

import (
	"bytes"
	"math/rand"
	"time"
	"unicode"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Text transforms a normal string into A spOnGEbOb sTrINg
func Text(str string) string {
	var result bytes.Buffer
	upper := false

	for _, c := range str {
		if unicode.IsLetter(rune(c)) {
			// instead of flipping a coin every time,
			// we create a poor man's Markov chain
			if rand.Float32() < 0.6 {
				upper = !upper
			}

			if upper {
				c = unicode.ToUpper(rune(c))
			} else {
				c = unicode.ToLower(rune(c))
			}
		}
		result.WriteRune(c)
	}

	return result.String()
}
