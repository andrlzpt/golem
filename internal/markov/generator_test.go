package markov

import (
	"reflect"
	"testing"
)

func TestUnigramGenerate(t *testing.T) {
	testcases := []struct {
		name             string
		input            []string
		word             string
		maxNumberOfWords int
		want             []string
	}{
		{
			name:             "deve gerar sequência correta",
			input:            []string{"o", "rei", "foi", "embora"},
			word:             "o",
			maxNumberOfWords: 3,
			want:             []string{"o", "rei", "foi"},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			chain := NewUnigramChain()
			chain.Train(testcase.input)

			result := GenerateNextDumb(chain, testcase.word, testcase.maxNumberOfWords)

			if !reflect.DeepEqual(result, testcase.want) {
				t.Fatalf("Generate() result = %#v, want %#v", result, testcase.want)
			}
		})
	}
}

func TestBigramGenerate(t *testing.T) {
	testcases := []struct {
		name             string
		input            []string
		word             string
		maxNumberOfWords int
		want             []string
	}{
		{
			name:             "deve gerar sequência correta",
			input:            []string{"o", "homem", "é", "uma", "ponte"},
			word:             "o homem",
			maxNumberOfWords: 5,
			want:             []string{"o", "homem", "é", "uma", "ponte"},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			chain := NewBigramChain()
			chain.Train(testcase.input)

			result := GenerateNextAtRandom(chain, testcase.word, testcase.maxNumberOfWords)

			if !reflect.DeepEqual(result, testcase.want) {
				t.Fatalf("Generate() result = %#v, want %#v", result, testcase.want)
			}
		})
	}
}
