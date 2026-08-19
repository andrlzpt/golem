package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/andrlzpt/golem/internal/markov"
	"github.com/andrlzpt/golem/internal/narrator"
	"github.com/andrlzpt/golem/internal/text"
)

var ErrReadingTrainingTextFile = errors.New("reading training text failed")

func main() {
	zaratustra := loadTrainingText("zaratustra")

	unigramChain := markov.NewChain(1)
	unigramNarrator := narrator.NewNarrator(unigramChain)
	unigramNarrator.Train(zaratustra)

	bigramChain := markov.NewChain(2)
	bigramNarrator := narrator.NewNarrator(bigramChain)
	bigramNarrator.Train(zaratustra)

	trigramChain := markov.NewChain(3)
	trigramNarrator := narrator.NewNarrator(trigramChain)
	trigramNarrator.Train(zaratustra)

	tetragramChain := markov.NewChain(4)
	tetragramNarrator := narrator.NewNarrator(tetragramChain)
	tetragramNarrator.Train(zaratustra)

	republica := loadTrainingText("republica")

	var input string

	fmt.Println("---UNIGRAM CHAIN: --------")
	input = extractQuote(republica, unigramChain)
	fmt.Printf("PLATO      SAYS: %q\n", input)
	unigramNarrator.Hear(input)
	fmt.Printf("ZARATUSTRA SAYS: %q\n", unigramNarrator.Speak(60))

	fmt.Println("---BIGRAM CHAIN: --------")
	input = extractQuote(republica, bigramChain)
	fmt.Printf("PLATO      SAYS: %q\n", input)
	bigramNarrator.Hear(input)
	fmt.Printf("ZARATUSTRA SAYS: %q\n", bigramNarrator.Speak(60))

	fmt.Println("---TRIGRAM CHAIN: --------")
	input = extractQuote(republica, trigramChain)
	fmt.Printf("PLATO      SAYS: %q\n", input)
	trigramNarrator.Hear(input)
	fmt.Printf("ZARATUSTRA SAYS: %q\n", trigramNarrator.Speak(60))

	fmt.Println("---TETRAGRAM CHAIN: --------")
	input = extractQuote(republica, tetragramChain)
	fmt.Printf("PLATO      SAYS: %q\n", input)
	tetragramNarrator.Hear(input)
	fmt.Printf("GOLEM SAYS: %q\n", tetragramNarrator.Speak(60))
}

func readInput(reader *bufio.Reader) string {
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	return input
}

func loadTrainingText(fileName string) string {
	path := "corpus/" + fileName + ".txt"
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", ErrReadingTrainingTextFile, err)
		os.Exit(1)
	}
	return string(data)
}

func extractQuote(source string, chain *markov.Chain) string {
	tokens := text.TokenizeTrainingText(source)
	sentences := text.SplitSentences(tokens)
	scored := text.ScoreSentences(sentences, chain.Order, chain.CanContinue)
	top := 40
	bestTen := scored[:top]
	return strings.Join(bestTen[rand.Intn(top)].Tokens, " ")

}
