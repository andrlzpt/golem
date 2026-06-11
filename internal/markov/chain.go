package markov

type Chain struct {
	Next map[string][]string
}

func NewChain() *Chain {
	return &Chain{
		Next: make(map[string][]string),
	}
}

func (c *Chain) Train(tokens []string) {
	for i := 0; i < len(tokens)-1; i++ {
		current := tokens[i]
		next := tokens[i+1]

		c.Next[current] = append(c.Next[current], next)
	}
}

func (c *Chain) NextWord(word string) string {
	options := c.Next[word]

	if len(options) == 0 {
		return ""
	}

	return options[0]
}
