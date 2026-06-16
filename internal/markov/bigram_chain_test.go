package markov

import (
	"slices"
	"testing"
)

func TestBigramNextWordAtRandom(t *testing.T) {
	testcases := []struct {
		name  string
		input []string
		next  string
		want  []string
	}{
		{
			name:  "deve retornar uma das opcoes válidas",
			input: []string{"o", "homem", "é", "uma", "ponte"},
			next:  "o homem",
			want:  []string{"é", "uma", "ponte"},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			chain := NewBigramChain()
			chain.Train(testcase.input)

			result := chain.NextWordAtRandom(testcase.next)

			if !slices.Contains(testcase.want, result) {
				t.Fatalf("NextWordAtRandom() result = %#v, want = %#v", result, testcase.want)
			}
		})
	}
}
