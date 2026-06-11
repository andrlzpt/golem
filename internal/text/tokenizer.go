package text

import (
	"strings"
	"unicode"
)

func Tokenize(input string) []string {
	lower := strings.ToLower(input)
	clear := keepAllLettersAndSpaces(lower)
	return strings.Fields(clear)
}

func keepAllLettersAndSpaces(input string) string {
	var result []rune

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			result = append(result, r)
		}
	}

	return string(result)
}
