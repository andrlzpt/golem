package narrator

import (
	"strings"

	"github.com/andrlzpt/golem/internal/markov"
	"github.com/andrlzpt/golem/internal/memory"
	"github.com/andrlzpt/golem/internal/text"
)

type Narrator struct {
	chain markov.Chain
	store *memory.Store
}

func NewNarrator() *Narrator {
	return &Narrator{
		chain: markov.NewUnigramChain(),
		store: memory.NewStore(),
	}
}

func (n *Narrator) Train(input string) {
	tokens := text.Tokenize(input)
	n.chain.Train(tokens)
}

func (n *Narrator) DumbSpeak(input string, maxNumberOfWords int) string {
	tokens := markov.GenerateNextDumb(n.chain, input, maxNumberOfWords)
	return processTokensIntoString(input, tokens)
}

func (n *Narrator) RandomSpeak(input string, maxNumberOfWords int) string {
	tokens := markov.GenerateNextAtRandom(n.chain, input, maxNumberOfWords)
	return processTokensIntoString(input, tokens)
}

func (n *Narrator) FromMemorySpeak(maxNumberOfWords int) string {
	input := n.store.Last()
	tokens := markov.GenerateNextAtRandom(n.chain, input, maxNumberOfWords)
	return processTokensIntoString(input, tokens)
}

func (n *Narrator) Hear(input string) {
	tokens := text.Tokenize(input)
	n.store.Store(tokens)
}

func (n *Narrator) TellAll() string {
	return strings.Join(n.store.View(), " ")
}

func (n *Narrator) ForgetAll() {
	n.store.Clear()
	n.chain.Clear()
}

func processTokensIntoString(input string, tokens []string) string {
	response := strings.Join(tokens, " ")
	if response == input {
		return ""
	}
	return response
}
