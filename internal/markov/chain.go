package markov

type Chain interface {
	Train(tokens []string)
	NextWordAlwaysFirstOption(word string) string
	NextWordAtRandom(word string) string
	Clear()
}
