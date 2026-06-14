package memory

type Memory struct {
	tokens []string
}

func NewMemory() *Memory {
	return &Memory{
		tokens: []string{},
	}
}

func (m *Memory) Remember(t []string) {
	m.tokens = append(m.tokens, t...)
}

func (m *Memory) View() []string {
	dst := make([]string, len(m.tokens))
	copy(dst, m.tokens)
	return dst
}
