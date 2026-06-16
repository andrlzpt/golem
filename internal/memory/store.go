package memory

type Store struct {
	tokens []string
}

func NewStore() *Store {
	return &Store{
		tokens: []string{},
	}
}

func (m *Store) Store(t []string) {
	m.tokens = append(m.tokens, t...)
}

func (m *Store) View() []string {
	dst := make([]string, len(m.tokens))
	copy(dst, m.tokens)
	return dst
}

func (m *Store) Clear() {
	m.tokens = nil
}
