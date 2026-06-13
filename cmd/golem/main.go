package main

import (
	"fmt"

	"github.com/andrlzpt/golem/internal/narrator"
)

func main() {
	narrator := narrator.NewNarrator()

	trainingText := "O rei olhou o espelho embaixo do rio e o espelho lembrou-se do espelho"
	narrator.Train(trainingText)

	sentence := narrator.Speak("o", 8)

	fmt.Printf("Training text: %q\n", trainingText)
	fmt.Printf("Golem says: %q\n", sentence)
}
