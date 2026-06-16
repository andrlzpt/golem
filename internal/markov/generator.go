package markov

import "strings"

func Generate(chain *Chain, tokens []string, maxNumberOfWords int) []string {
	if len(tokens) != chain.Order {
		return []string{}
	}

	result := make([]string, len(tokens))
	copy(result, tokens)
	window := make([]string, len(tokens))
	copy(window, tokens)

	for len(result) < maxNumberOfWords {
		key := strings.Join(window, " ")
		next := chain.NextWord(key)

		if next == "" {
			return result
		}

		result = append(result, next)
		window = append(window[1:], next)
	}

	return result
}
