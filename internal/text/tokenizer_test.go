package text

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	testcases := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "deve tokenizar corretamente",
			input: "Pois é, meu rei, markov chain em GO!!!123##$",
			want:  []string{"pois", "é", "meu", "rei", "markov", "chain", "em", "go"},
		},
		{
			name:  "deve filtrar dashes corretamente",
			input: " -, ---, --------, +, +++, $$$, ##, lembrou-se, --deve passar, --a,",
			want:  []string{"lembrou-se", "--deve", "passar", "--a"},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			result := Tokenize(testcase.input)
			if !reflect.DeepEqual(result, testcase.want) {
				t.Fatalf("Tokenize() result = %#v, want: = %#v", result, testcase.want)
			}
		})
	}
}
