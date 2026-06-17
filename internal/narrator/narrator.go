package narrator

import (
	"strings"

	"github.com/andrlzpt/golem/internal/markov"
	"github.com/andrlzpt/golem/internal/memory"
	"github.com/andrlzpt/golem/internal/text"
)

type Narrator struct {
	chain *markov.Chain
	store *memory.Store
}

func NewNarrator(mc *markov.Chain) *Narrator {
	return &Narrator{
		chain: mc,
		store: memory.NewStore(),
	}
}

func (n *Narrator) Train(input string) {
	tokens := text.Tokenize(input)
	n.chain.Train(tokens)
}

func (n *Narrator) Speak(maxNumberOfWords int) string {
	memory := n.store.View()
	start := chooseStartTokens(memory, n.chain)
	tokens := markov.Generate(n.chain, start, maxNumberOfWords)
	if len(tokens) == len(start) {
		return "Não sei nada sobre isso."
	}
	return strings.Join(tokens, " ")

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

func chooseStartTokens(memory []string, chain *markov.Chain) []string {
	order := chain.Order
	if order <= 0 || len(memory) < order {
		return []string{}
	}

	for i := len(memory) - order; i >= 0; i-- {
		candidate := memory[i : i+order]
		key := strings.Join(candidate, " ")
		if chain.HasNext(key) {
			return candidate
		}
	}

	return []string{}
}
