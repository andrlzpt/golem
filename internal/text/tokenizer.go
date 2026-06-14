package text

import (
	"strings"
	"unicode"
)

func Tokenize(input string) []string {
	lower := strings.ToLower(input)
	clear := keepAllLettersAndSpaces(lower)
	tokens := strings.Fields(clear)
	return filter(tokens)

}

func keepAllLettersAndSpaces(input string) string {
	var result []rune

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsSpace(r) || r == '-' {
			result = append(result, r)
		}
	}

	return string(result)
}

func filter(tokens []string) []string {
	var filtered []string
	for _, t := range tokens {
		hasLetter := strings.ContainsFunc(t, func(r rune) bool {
			return unicode.IsLetter(r)
		})

		if !hasLetter {
			continue
		}
		filtered = append(filtered, t)
	}
	return filtered
}
