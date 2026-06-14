package narrator

import (
	"strings"

	"github.com/andrlzpt/golem/internal/markov"
	"github.com/andrlzpt/golem/internal/text"
)

type Narrator struct {
	chain *markov.Chain
}

func NewNarrator() *Narrator {
	return &Narrator{
		chain: markov.NewChain(),
	}
}

func (n *Narrator) Train(input string) {
	tokens := text.Tokenize(input)
	n.chain.Train(tokens)
}

func (n *Narrator) DumbSpeak(start string, maxNumberOfWords int) string {
	tokens := markov.GenerateNextDumb(n.chain, start, maxNumberOfWords)
	response := strings.Join(tokens, " ")
	if response == start {
		return ""
	}
	return response
}

func (n *Narrator) RandomSpeak(start string, maxNumberOfWords int) string {
	tokens := markov.GenerateNextAtRandom(n.chain, start, maxNumberOfWords)
	response := strings.Join(tokens, " ")
	if response == start {
		return ""
	}
	return response
}
