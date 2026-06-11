package markov

func Generate(chain *Chain, word string, maxNumberOfWords int) []string {
	result := []string{}
	current := word
	for range maxNumberOfWords {
		if current == "" {
			return result
		}

		result = append(result, current)
		current = chain.NextWord(current)
	}
	return result
}
