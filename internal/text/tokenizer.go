package text

import (
	"strings"
	"unicode"
)

const EndToken = "<END>"

func TokenizeTrainingText(input string) []string {
	lower := strings.ToLower(input)
	clear := clean(lower, true)
	tokens := strings.Fields(clear)
	return filter(tokens, true)

}

func Tokenize(input string) []string {
	lower := strings.ToLower(input)
	clear := clean(lower, false)
	tokens := strings.Fields(clear)
	return filter(tokens, false)

}

func clean(input string, addBoundaryMarker bool) string {
	var result []rune

	for _, r := range input {
		if addBoundaryMarker && RuneIsSentenceEnding(r) {
			result = append(result, []rune(" "+EndToken+" ")...)
			continue
		}
		if unicode.IsLetter(r) || unicode.IsSpace(r) || r == '-' {
			result = append(result, r)
		}
	}

	return string(result)
}

func filter(tokens []string, keepEndToken bool) []string {
	var filtered []string
	for _, t := range tokens {
		hasLetter := strings.ContainsFunc(t, func(r rune) bool {
			return unicode.IsLetter(r)
		}) || (keepEndToken && t == EndToken)

		if !hasLetter {
			continue
		}
		filtered = append(filtered, t)
	}
	return filtered
}
