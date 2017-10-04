package nickgen

import (
	"bytes"
	"math/rand"
)

const (
	consonants = "bcdfghjklmnprstvw"
	vocals     = "aeiou"
)

func pickRune(s string) rune {
	return rune(s[rand.Intn(len(s))])
}

// GenerateWord generates a speakable word
func GenerateWord(len int) string {
	var w bytes.Buffer
	vocalnext := rand.Intn(3) == 0
	cs := 0
	for i := 0; i < len; i++ {
		if vocalnext {
			w.WriteRune(pickRune(vocals))
		} else {
			w.WriteRune(pickRune(consonants))
		}
		if !vocalnext && cs < 2 && w.Len() > 0 {
			if rand.Intn(3) == 0 {
				cs++
			} else {
				vocalnext = true
				cs = 0
			}
		} else {
			vocalnext = !vocalnext
			if !vocalnext {
				cs = 1
			} else {
				cs = 0
			}
		}
	}
	return w.String()
}
