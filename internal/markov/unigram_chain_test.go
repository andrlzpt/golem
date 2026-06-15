package markov

import (
	"reflect"
	"slices"
	"testing"
)

func TestTrain(t *testing.T) {
	testcases := []struct {
		name  string
		input []string
		next  string
		want  []string
	}{
		{
			name:  "deve criar chain correta para o",
			input: []string{"o", "rei", "olhou", "o", "espelho"},
			next:  "o",
			want:  []string{"rei", "espelho"},
		},
		{
			name:  "deve criar chain correta com palavras repetidas para um mesmo next",
			input: []string{"o", "rei", "olhou", "o", "rei", "no", "espelho", "o", "outro", "rei"},
			next:  "o",
			want:  []string{"rei", "rei", "outro"},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			chain := NewUnigramChain()
			chain.Train(testcase.input)
			next := chain.Next[testcase.next]
			if !reflect.DeepEqual(next, testcase.want) {
				t.Fatalf("Train() result = %#v, want = %#v", next, testcase.want)
			}
		})
	}
}

func TestNextWord(t *testing.T) {
	testcases := []struct {
		name  string
		input []string
		next  string
		want  string
	}{
		{
			name:  "deve obter a primeira opção na chain corretamente",
			input: []string{"o", "rei", "foi"},
			next:  "o",
			want:  "rei",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			chain := NewUnigramChain()
			chain.Train(testcase.input)

			result := chain.NextWordAlwaysFirstOption(testcase.next)
			if result != testcase.want {
				t.Fatalf("NextWord(%v) result = %#v, want = %#v", testcase.next, result, testcase.want)
			}
		})
	}
}

func TestUnigramNextWordAtRandom(t *testing.T) {
	testcases := []struct {
		name  string
		input []string
		next  string
		want  []string
	}{
		{
			name:  "deve retornar uma das opcoes válidas",
			input: []string{"o", "rei", "olhou", "o", "espelho"},
			next:  "o",
			want:  []string{"rei", "espelho"},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			chain := NewUnigramChain()
			chain.Train(testcase.input)

			result := chain.NextWordAtRandom(testcase.next)

			if !slices.Contains(testcase.want, result) {
				t.Fatalf("NextWordAtRandom() result = %#v, want = %#v", result, testcase.want)
			}
		})
	}
}
