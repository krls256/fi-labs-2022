// Package pe implements probabilistic ensemble
package pe

import (
	"cp_1/slices"
)

type probabilisticEnsemble struct {
	store          map[string]int
	allowedSymbols map[string]struct{}
	analyzableText []int32
}

func New(allowedSymbols []string) *probabilisticEnsemble {
	preparedSymbols := map[string]struct{}{}
	for _, str := range allowedSymbols {
		preparedSymbols[str] = struct{}{}
	}
	return &probabilisticEnsemble{
		store:          make(map[string]int),
		allowedSymbols: preparedSymbols,
		analyzableText: make([]int32, 0),
	}
}

func (pe *probabilisticEnsemble) AddText(text string) {
	for _, letter := range text {
		if _, ok := pe.allowedSymbols[string(letter)]; ok {
			pe.analyzableText = append(pe.analyzableText, letter)
		}
	}
}

func (pe *probabilisticEnsemble) AddTextWithSpaces(text string) {
	buffer := make([]int32, 0)
	for _, letter := range text {
		if _, ok := pe.allowedSymbols[string(letter)]; ok ||
			(string(letter) == " " && len(buffer) != 0 && string(buffer[len(buffer)-1]) != " ") {
			buffer = append(buffer, letter)
		}
	}
	pe.analyzableText = append(pe.analyzableText, buffer...)
}

func (pe *probabilisticEnsemble) GetNGramStatWithIntersection(n int) *NGramStat {
	return pe.getNGramStatByStep(n, 1)
}
func (pe *probabilisticEnsemble) GetNGramStatWithoutIntersection(n int) *NGramStat {
	return pe.getNGramStatByStep(n, n)
}

func (pe *probabilisticEnsemble) getNGramStatByStep(n, step int) *NGramStat {
	stat := newStat(n)
	for _, slice := range slices.Split[int32](pe.analyzableText, ' ') {
		for i := 0; i < len(slice)-n; i += step {
			str := string(slice[i : i+n])
			stat.AddNGram(str)
		}
	}
	return stat
}
