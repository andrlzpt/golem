package markov

import (
	"math/rand"
)

var _ Chain = (*UnigramChain)(nil)

type UnigramChain struct {
	Next map[string][]string
}

func NewUnigramChain() *UnigramChain {
	return &UnigramChain{
		Next: make(map[string][]string),
	}
}

func (c *UnigramChain) Train(tokens []string) {
	for i := 0; i < len(tokens)-1; i++ {
		current := tokens[i]
		next := tokens[i+1]

		c.Next[current] = append(c.Next[current], next)
	}
}

func (c *UnigramChain) NextWordAlwaysFirstOption(word string) string {
	options := c.Next[word]

	if len(options) == 0 {
		return ""
	}

	return options[0]
}

func (c *UnigramChain) NextWordAtRandom(word string) string {
	options := c.Next[word]

	if len(options) == 0 {
		return ""
	}

	return options[rand.Intn(len(options))]
}

func (c *UnigramChain) Clear() {
	c.Next = make(map[string][]string)
}
