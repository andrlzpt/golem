package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/andrlzpt/golem/internal/narrator"
)

var ErrReadingTrainingTextFile = errors.New("reading training text failed")

func main() {
	narrator := narrator.NewNarrator()
	fmt.Println("------------------------------------")
	fmt.Println("---------TERMINAL GOLEM -----------")

	reader := bufio.NewReader(os.Stdin)

	// trainingText := "O rei olhou o espelho embaixo do rio e o espelho lembrou-se do espelho"
	// fmt.Printf("Training text: %q\n", trainingText)

	// narrator.Train(trainingText)
	// fmt.Println("------------------------------------")
	// fmt.Println("---DUMB SPEAK (ALWAYS FIRST NEXT)---")
	// sentence := narrator.DumbSpeak("o", 8)
	// fmt.Printf("Result: %q\n", sentence)
	// fmt.Println("------------------------------------")
	// fmt.Println("---RANDOM SPEAK --------------------")
	// sentence = narrator.RandomSpeak("o", 8)
	// fmt.Printf("Result: %q\n", sentence)
	// fmt.Println("OBS: this is at random, but there is a weighted frequency selection component")
	// fmt.Println("------------------------------------")
	// fmt.Println("---LISTENS TO INPUT SPEAK --------------------")
	// fmt.Print("Enter text: ")
	// input, _ := reader.ReadString('\n')
	// narrator.Hear(input)
	// sentence = narrator.TellAll()
	// fmt.Printf("Memory: %q\n", sentence)
	// fmt.Println("---FROM MEMORY SPEAK --------------------")
	// sentence = narrator.FromMemorySpeak(4)
	// fmt.Printf("Result: %q\n", sentence)
	// fmt.Println("------------------------------------")
	fmt.Println("---ZARATUSTRA TRAINING TEXT --------")
	narrator.ForgetAll()
	path := "corpus/zaratustra.txt"
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", ErrReadingTrainingTextFile, err)
		os.Exit(1)
	}
	trainingText := string(data)
	narrator.Train(trainingText)
	fmt.Println("read the Zaratustra Book. Ready to say something")
	fmt.Println("---LISTENS TO INPUT SPEAK --------------------")
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	narrator.Hear(input)
	fmt.Println("---FROM MEMORY SPEAK --------------------")
	sentence := narrator.FromMemorySpeak(38)
	fmt.Printf("Result: %q\n", sentence)
}
