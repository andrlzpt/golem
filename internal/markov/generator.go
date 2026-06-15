package markov

import "strings"

type next func(string) string

func GenerateNextDumb(chain Chain, word string, maxNumberOfWords int) []string {
	return generateWith(word, maxNumberOfWords, chain.NextWordAlwaysFirstOption)
}

func GenerateNextAtRandom(chain Chain, word string, maxNumberOfWords int) []string {
	return generateWith(word, maxNumberOfWords, chain.NextWordAtRandom)
}

func generateWith(input string, maxNumberOfWords int, n next) []string {
	fields := strings.Fields(input)

	if len(fields) == 1 {
		return unigramGenerate(input, maxNumberOfWords, n)
	}

	if len(fields) == 2 {
		return bigramGenerate(input, maxNumberOfWords, n)
	}

	return []string{}
}

func unigramGenerate(input string, maxNumberOfWords int, n next) []string {
	result := []string{}
	current := input
	for range maxNumberOfWords {
		if current == "" {
			return result
		}

		result = append(result, current)
		current = n(current)
	}
	return result
}

func bigramGenerate(input string, maxNumberOfWords int, n next) []string {
	parts := strings.Fields(input)
	result := []string{parts[0], parts[1]}

	first := parts[0]
	second := parts[1]

	for len(result) < maxNumberOfWords {
		key := first + " " + second
		nextWord := n(key)

		if nextWord == "" {
			return result
		}

		result = append(result, nextWord)

		first = second
		second = nextWord
	}

	return result
}
