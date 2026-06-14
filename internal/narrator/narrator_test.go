package narrator

import (
	"testing"
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
			name:             "deve gerar frase vazia",
			input:            "O rei foi embora!",
			start:            "",
			maxNumberOfWords: 3,
			want:             "",
		},
		{
			name:             "deve gerar frase vazia com palavra que não existe",
			input:            "O rei foi embora!",
			start:            "batatinha",
			maxNumberOfWords: 3,
			want:             "",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			narrator := NewNarrator()
			narrator.Train(testcase.input)
			result := narrator.DumbSpeak(testcase.start, testcase.maxNumberOfWords)

			if result != testcase.want {
				t.Fatalf("Speak(%v, %d) result = %#v, want = %#v", testcase.start, testcase.maxNumberOfWords, result, testcase.want)
			}
		})
	}

}
