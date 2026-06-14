package memory

import (
	"reflect"
	"testing"

	"github.com/andrlzpt/golem/internal/text"
)

func TestView(t *testing.T) {
	testcases := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "deve registrar e visualizar memória corretamente",
			input: "interação será tokenizada antes de ser lembrado",
			want:  []string{"interação", "será", "tokenizada", "antes", "de", "ser", "lembrado"},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			memory := NewStore()
			tokens := text.Tokenize(testcase.input)
			memory.Store(tokens)

			result := memory.View()

			if !reflect.DeepEqual(result, testcase.want) {
				t.Fatalf("View() return = %#v, want = %#v", result, testcase.want)
			}
		})
	}
}
