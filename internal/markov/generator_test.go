package markov

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	testcases := []struct {
		name             string
		input            []string
		start            []string
		maxNumberOfWords int
		want             []string
	}{
		{
			name:             "deve gerar sequência correta",
			input:            []string{"o", "rei", "foi", "embora"},
			start:            []string{"o"},
			maxNumberOfWords: 3,
			want:             []string{"o", "rei", "foi"},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			chain := NewChain(1)
			chain.Train(testcase.input)
			result := Generate(chain, testcase.start, testcase.maxNumberOfWords)
			if !reflect.DeepEqual(result, testcase.want) {
				t.Fatalf("Generate() result = %#v, want %#v", result, testcase.want)
			}
		})
	}
}
