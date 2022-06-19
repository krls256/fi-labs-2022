package freqTools

import (
	"cp_2/vigenere"
	"errors"
	"fmt"
	"github.com/igrmk/treemap/v2"
	"math"
)

type LanguageDistribution struct {
	Total int
	Tree  *treemap.TreeMap[int, rune]
}

type ReversedLanguageDistribution struct {
	Total int
	Tree  *treemap.TreeMap[rune, int]
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

func ReverseCountFrequencies(text []rune) *ReversedLanguageDistribution {
	tr := treemap.New[rune, int]()
	for k, v := range countFrequencies(text) {
		tr.Set(k, v)
	}

	return &ReversedLanguageDistribution{
		Tree:  tr,
		Total: len(text),
	}
}

func FindKeyLen(cipherText []rune, keyLenLimit int) (int, error) {
	baseKeyLen := 2
	correspondenceIndexDelta := 1.0 / 100 // 1%
	stat := make([]float64, keyLenLimit-baseKeyLen+1)
	for i := baseKeyLen; i <= keyLenLimit; i++ {
		stat[i-baseKeyLen] = CorrespondenceIndex(SplitByKeyLen(cipherText, i)[0])
	}
	for k, v := range stat {
		fmt.Printf("r = %v, index = %1.2f%%\n", k+2, v*100)
	}
	keyLen := -1
	for i := 1; i < len(stat); i++ {
		if math.Abs(stat[i]-stat[i-1]) > correspondenceIndexDelta {
			keyLen = i + baseKeyLen
			if stat[i] < stat[i-1] {
				keyLen--
			}
			break
		}
	}

	if keyLen == -1 {
		return -1, errors.New("can not find key len")
	}
	return keyLen, nil
}

func FindKey(cipherText []rune, keyLen int, defaultDistribution *LanguageDistribution, alphabet []rune) []int {
	key := make([]int, keyLen)
	split := SplitByKeyLen(cipherText, keyLen)
	for i := 0; i < keyLen; i++ {
		distr := CountFrequencies(split[i])
		iterator := distr.Tree.Reverse()
		key[i] = int((iterator.Value() - defaultDistribution.Tree.Reverse().Value() + rune(len(alphabet))) % rune(len(alphabet)))
	}
	return key
}

func FindKeyByM(cipherText []rune, keyLen int, defaultDistribution *ReversedLanguageDistribution, alp *vigenere.Cipher) []int {
	key := make([]int, keyLen)
	split := SplitByKeyLen(cipherText, keyLen)
	stat := []*treemap.TreeMap[rune, float64]{}
	for i := 0; i < keyLen; i++ {
		distr := ReverseCountFrequencies(split[i])
		M := map[rune]float64{}
		for iter := distr.Tree.Iterator(); iter.Valid(); iter.Next() {
			M[iter.Key()] = m(iter.Key(), defaultDistribution, distr, alp)
		}
		var (
			maxValue float64
			maxRune  rune
		)
		statMap := treemap.New[rune, float64]()
		for r, value := range M {
			if value > maxValue {
				maxValue = value
				maxRune = r
				statMap.Set(r, value)
			}
		}
		key[i] = int(maxRune - alp.Alphabet[0])
		stat = append(stat, statMap)
	}
	//showStat(stat)
	return key
}

func showStat(stat []*treemap.TreeMap[rune, float64]) {
	for i := 0; i < len(stat); i++ {
		m := stat[i]
		fmt.Print(i+1, " ")
		for iter := m.Iterator(); iter.Valid(); iter.Next() {
			fmt.Printf("%v: %1.3f ", string(iter.Key()), iter.Value())
		}
		fmt.Print("\n")
	}

}

func m(g rune, defaultDistribution, currentDistribution *ReversedLanguageDistribution, alp *vigenere.Cipher) float64 {
	m := 0.0
	for iter := defaultDistribution.Tree.Iterator(); iter.Valid(); iter.Next() {
		p := float64(iter.Value()) / float64(defaultDistribution.Total)
		r := alp.Alphabet[(alp.ReversedAlphabet[g]+alp.ReversedAlphabet[iter.Key()])%len(alp.Alphabet)]
		N, _ := currentDistribution.Tree.Get(r)
		m += p * float64(N)
	}
	return m
}
