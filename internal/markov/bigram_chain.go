package markov

import "math/rand"

var _ Chain = (*BigramChain)(nil)

type BigramChain struct {
	Next map[string][]string
}

func NewBigramChain() *BigramChain {
	return &BigramChain{
		Next: make(map[string][]string),
	}
}

func (c *BigramChain) Train(tokens []string) {
	for i := 0; i < len(tokens)-2; i++ {
		first := tokens[i]
		second := tokens[i+1]
		next := tokens[i+2]
		key := first + " " + second
		c.Next[key] = append(c.Next[key], next)
	}
}

func (c *BigramChain) Clear() {
	c.Next = make(map[string][]string)
}

func (c *BigramChain) NextWordAtRandom(word string) string {
	options := c.Next[word]

	if len(options) == 0 {
		return ""
	}

	return options[rand.Intn(len(options))]
}

func (c *BigramChain) NextWordAlwaysFirstOption(word string) string {
	return ""
}
