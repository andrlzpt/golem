package narrator

import (
	"testing"

	"github.com/andrlzpt/golem/internal/markov"
)

func TestSpeak(t *testing.T) {
	testcases := []struct {
		name             string
		input            string
		start            string
		maxNumberOfWords int
		want             string
	}{
		{
			name:             "deve gerar a frase correta",
			input:            "O rei foi embora!",
			start:            "o",
			maxNumberOfWords: 3,
			want:             "o rei foi",
		},
		{
			name:             "deve gerar frase vazia e cair no fallback",
			input:            "O rei foi embora!",
			start:            "",
			maxNumberOfWords: 3,
			want:             "Não sei nada sobre isso.",
		},
		{
			name:             "deve gerar frase vazia com palavra que não existe e cair no fallback",
			input:            "O rei foi embora!",
			start:            "batatinha",
			maxNumberOfWords: 3,
			want:             "Não sei nada sobre isso.",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			narrator := NewNarrator(markov.NewChain(1))
			narrator.Train(testcase.input)
			narrator.Hear(testcase.start)
			result := narrator.Speak(testcase.maxNumberOfWords)

			if result != testcase.want {
				t.Fatalf("Speak(%v, %d) result = %#v, want = %#v", testcase.start, testcase.maxNumberOfWords, result, testcase.want)
			}
		})
	}

}

func TestTellAll(t *testing.T) {
	testcases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "deve lembrar e contar memória corretamente",
			input: "lembre disso",
			want:  "lembre disso",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			narrator := NewNarrator(markov.NewChain(1))
			narrator.Hear(testcase.input)
			result := narrator.TellAll()

			if result != testcase.want {
				t.Fatalf("TellAll() result = %#v, want = %#v", result, testcase.want)
			}
		})
	}
}

func TestFromMemorySpeak(t *testing.T) {
	testcases := []struct {
		name         string
		trainingText string
		input        string
		want         string
	}{
		{
			name:         "deve escolher último item da memória como start next",
			trainingText: "o rio lembrou-se do rei ",
			input:        "Eu olhei o rio",
			want:         "rio lembrou-se do rei",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			narrator := NewNarrator(markov.NewChain(1))
			narrator.Train(testcase.trainingText)
			narrator.Hear(testcase.input)
			result := narrator.Speak(4)
			if result != testcase.want {
				t.Fatalf("FromMemorySpeak(4) result = %#v, want = %#v", result, testcase.want)
			}

		})
	}
}
