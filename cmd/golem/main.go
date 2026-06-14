package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/andrlzpt/golem/internal/narrator"
)

func main() {
	narrator := narrator.NewNarrator()

	trainingText := "O rei olhou o espelho embaixo do rio e o espelho lembrou-se do espelho"
	fmt.Printf("Training text: %q\n", trainingText)

	narrator.Train(trainingText)
	fmt.Println("------------------------------------")
	fmt.Println("---DUMB SPEAK (ALWAYS FIRST NEXT)---")
	sentence := narrator.DumbSpeak("o", 8)
	fmt.Printf("Result: %q\n", sentence)
	fmt.Println("------------------------------------")
	fmt.Println("---RANDOM SPEAK --------------------")
	sentence = narrator.RandomSpeak("o", 8)
	fmt.Printf("Result: %q\n", sentence)
	fmt.Println("OBS: this is at random, but there is a weighted frequency selection component")
	fmt.Println("------------------------------------")
	fmt.Println("---LISTENS TO INPUT SPEAK --------------------")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	narrator.Hear(input)
	sentence = narrator.TellAll()
	fmt.Printf("Memory: %q\n", sentence)

}
