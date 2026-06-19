package text

import "slices"

func RuneIsSentenceEnding(r rune) bool {
	return slices.Contains(sentenceEndings, r)
}

var sentenceEndings = []rune{'.', '!', '?', ';', ':'}

func HasContentWord(tokens []string) bool {
	for _, token := range tokens {
		_, isStopword := portugueseStopwords[token]
		if !isStopword {
			return true
		}
	}
	return false
}

var portugueseStopwords = map[string]struct{}{
	"o":    {},
	"a":    {},
	"os":   {},
	"as":   {},
	"de":   {},
	"do":   {},
	"da":   {},
	"dos":  {},
	"das":  {},
	"que":  {},
	"e":    {},
	"se":   {},
	"em":   {},
	"no":   {},
	"na":   {},
	"nos":  {},
	"nas":  {},
	"para": {},
	"por":  {},
	"com":  {},
}
