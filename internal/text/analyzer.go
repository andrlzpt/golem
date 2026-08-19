package text

import (
	"slices"
	"strings"
)

type ScoredSentence struct {
	Tokens []string
	Score  int
}

type isUseful func(key string) bool

func ScoreSentences(sentences [][]string, order int, check isUseful) []ScoredSentence {
	if order <= 0 {
		return []ScoredSentence{}
	}
	var scored []ScoredSentence
	for _, sentence := range sentences {
		if len(sentence) < 6 || len(sentence) > 48 {
			continue
		}
		if !HasContentWord(sentence) {
			continue
		}
		score := 0

		for i := 0; i <= len(sentence)-order; i++ {
			key := strings.Join(sentence[i:i+order], " ")
			if check(key) {
				score += 10
			}
		}

		if score > 0 {
			scored = append(scored, ScoredSentence{Tokens: sentence, Score: score})
		}
	}
	return scored
}

func RuneIsSentenceEnding(r rune) bool {
	return slices.Contains(sentenceEndings, r)
}

var sentenceEndings = []rune{'.', '!', '?', ';', ':'}

func HasContentWord(tokens []string) bool {
	for _, token := range tokens {
		_, isStopword := portugueseStopwords[token]
		if !isStopword {
			return true
		}
	}
	return false
}

var portugueseStopwords = map[string]struct{}{
	"o":    {},
	"a":    {},
	"os":   {},
	"as":   {},
	"de":   {},
	"do":   {},
	"da":   {},
	"dos":  {},
	"das":  {},
	"que":  {},
	"e":    {},
	"se":   {},
	"em":   {},
	"no":   {},
	"na":   {},
	"nos":  {},
	"nas":  {},
	"para": {},
	"por":  {},
	"com":  {},
}
