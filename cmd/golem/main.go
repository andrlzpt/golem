package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/andrlzpt/golem/internal/markov"
	"github.com/andrlzpt/golem/internal/narrator"
	"github.com/andrlzpt/golem/internal/text"
)

var ErrReadingTrainingTextFile = errors.New("reading training text failed")

func main() {
	zaratustra := loadTrainingText("zaratustra")

	republica := loadTrainingText("republica")

	reader := bufio.NewReader(os.Stdin)

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

	input := readInput(reader)

	unigramNarrator.Hear(input)
	bigramNarrator.Hear(input)
	trigramNarrator.Hear(input)
	tetragramNarrator.Hear(input)

	fmt.Println("---UNIGRAM CHAIN: --------")
	fmt.Printf("GOLEM SAYS: %q\n", unigramNarrator.Speak(60))
	fmt.Println("---BIGRAM CHAIN: --------")
	fmt.Printf("GOLEM SAYS: %q\n", bigramNarrator.Speak(60))
	fmt.Println("---TRIGRAM CHAIN: --------")
	fmt.Printf("GOLEM SAYS: %q\n", trigramNarrator.Speak(60))
	fmt.Println("---TETRAGRAM CHAIN: --------")
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

func extractQuote(source string, chain *markov.Chain) {
	tokens := text.TokenizeTrainingText(source)
	sentences := splitSentences(tokens)

}

func splitSentences(tokens []string) [][]string {
	var sentences [][]string
	var current []string

	for _, token := range tokens {
		if token == text.EndToken {
			if len(current) > 0 {
				sentences = append(sentences, current)
				current = nil
			}
			continue
		}

		current = append(current, token)
	}

	if len(current) > 0 {
		sentences = append(sentences, current)
	}

	return sentences
}
