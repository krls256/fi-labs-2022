package freqTools

import (
	"github.com/igrmk/treemap/v2"
)

type LanguageDistribution struct {
	Total int
	Tree  *treemap.TreeMap[int, rune]
}

func CorrespondenceIndex(text []rune) float64 {
	counter := countFrequencies(text)
	tmp := 0.0
	for _, value := range counter {
		tmp += float64(value) * (float64(value) - 1)
	}

	return tmp / (float64(len(text)) * (float64(len(text)) - 1))
}

func CorrespondenceIndexExpectation(text []rune) float64 {
	counter := countFrequencies(text)
	n := float64(len(text))
	res := 0.0
	for _, value := range counter {
		res += (float64(value) / n) * (float64(value) / n)
	}
	return res
}

func SplitByKeyLen(text []rune, keyLen int) [][]rune {
	res := make([][]rune, keyLen)
	for i, r := range text {
		res[i%keyLen] = append(res[i%keyLen], r)
	}

	return res
}

func SplitByLen(text []rune, length int) [][]rune {
	res := make([][]rune, 0)
	for i := 0; i < len(text)/length; i++ {
		res = append(res, text[i*length:(i+1)*length])
	}

	return res
}

func countFrequencies(text []rune) map[rune]int {
	counter := map[rune]int{}
	for _, r := range text {
		counter[r]++
	}
	return counter
}

func CountFrequencies(text []rune) *LanguageDistribution {
	tr := treemap.New[int, rune]()
	for k, v := range countFrequencies(text) {
		tr.Set(v, k)
	}

	return &LanguageDistribution{
		Tree:  tr,
		Total: len(text),
	}
}
