package markov

import (
	"math/rand"
	"slices"
	"strings"

	"github.com/andrlzpt/golem/internal/text"
)

type Chain struct {
	Order int
	Next  map[string][]string
}

func NewChain(order int) *Chain {
	return &Chain{
		Order: order,
		Next:  make(map[string][]string),
	}
}

func (c *Chain) Train(tokens []string) {
	if c.Order <= 0 {
		return
	}

	if len(tokens) <= c.Order {
		return
	}

	for i := 0; i < len(tokens)-c.Order; i++ {
		key := strings.Join(tokens[i:i+c.Order], " ")
		next := tokens[i+c.Order]
		c.Next[key] = append(c.Next[key], next)
	}
}

func (c *Chain) NextWord(key string) string {
	options := c.Next[key]

	if len(options) == 0 {
		return ""
	}

	return options[rand.Intn(len(options))]
}

func (c *Chain) CanContinue(key string) bool {
	return slices.ContainsFunc(c.Next[key], func(n string) bool {
		return n != text.EndToken
	})
}

func (c *Chain) Clear() {
	c.Next = make(map[string][]string)
}

func (c *Chain) HasNext(key string) bool {
	return len(c.Next[key]) > 0
}
