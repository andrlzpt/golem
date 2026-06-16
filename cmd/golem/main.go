package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/andrlzpt/golem/internal/markov"
	"github.com/andrlzpt/golem/internal/narrator"
)

var ErrReadingTrainingTextFile = errors.New("reading training text failed")

func main() {
	fmt.Println("---DER GOLEM =----------------------")

	fmt.Println("---ZARATUSTRA TRAINING TEXT --------")
	path := "corpus/zaratustra.txt"
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", ErrReadingTrainingTextFile, err)
		os.Exit(1)
	}
	zaratustra := string(data)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---UNIGRAM CHAIN: --------")
	unigramChain := markov.NewChain(1)
	unigramNarrator := narrator.NewNarrator(unigramChain)
	unigramNarrator.Train(zaratustra)
	unigramNarrator.Hear(readInput(reader))
	fmt.Printf("GOLEM SAYS: %q\n", unigramNarrator.Speak(60))

	fmt.Println("---BIGRAM CHAIN: --------")
	bigramChain := markov.NewChain(2)
	bigramNarrator := narrator.NewNarrator(bigramChain)
	bigramNarrator.Train(zaratustra)
	bigramNarrator.Hear(readInput(reader))
	fmt.Printf("GOLEM SAYS: %q\n", bigramNarrator.Speak(60))

}

func readInput(reader *bufio.Reader) string {
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	return input
}
