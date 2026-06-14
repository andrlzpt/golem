package markov

type next func(string) string

func GenerateNextDumb(chain *Chain, word string, maxNumberOfWords int) []string {
	return generateWith(word, maxNumberOfWords, chain.NextWordAlwaysFirstOption)
}

func GenerateNextAtRandom(chain *Chain, word string, maxNumberOfWords int) []string {
	return generateWith(word, maxNumberOfWords, chain.NextWordAtRandom)
}

func generateWith(word string, maxNumberOfWords int, n next) []string {
	result := []string{}
	current := word
	for range maxNumberOfWords {
		if current == "" {
			return result
		}

		result = append(result, current)
		current = n(current)
	}
	return result
}
