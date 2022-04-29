package main

import (
	"cp_3/alphabet"
	"cp_3/cipher"
	"cp_3/ring"
	"cp_3/stat"
	"fmt"
	"os"
)

func main() {

	r, err := ring.NewRing(alphabet.AlpLen * alphabet.AlpLen)
	handleByPanic(err)
	textB, err := os.ReadFile("./cipherText")
	handleByPanic(err)
	cipherText := string(textB)
	textB, err = os.ReadFile("./graf-monte-kristo.txt")
	cleanText := string(textB)
	cleanText = "стнотонаен"
	cipherBigram, err := alphabet.StringToBigrams(cipherText)
	handleByPanic(err)
	cleanBigram, err := alphabet.StringToBigrams(cleanText)
	handleByPanic(err)

	statCipher, statClean, tmpStat := stat.Stat{}, stat.Stat{}, stat.Stat{}
	statCipher.Append(cipherBigram)
	statClean.Append(cleanBigram)

	cipherIter := statCipher.BackIterator()
	cleanIter := statClean.BackIterator()

	keys := bruteforceKeys(cipherIter, cleanIter, r, 5)
	for _, v := range keys {
		a, b := v[0], v[1]
		aN, err := r.Inverse(a)
		if err != nil {
			continue
		}
		deciphered := cipher.Dec(cipherBigram, r, aN, b)
		str, err := alphabet.BigramsToString(deciphered)
		handleByPanic(err)
		monograms, err := alphabet.StringToMonograms(str)
		handleByPanic(err)
		tmpStat.Append(monograms)
		tmpIter := tmpStat.BackIterator()
		if CheckRealText(tmpIter, 2, 5) {
			fmt.Print(alphabet.MonogramsToString(monograms[:50]))
			fmt.Print(alphabet.SingleMonogramToString(tmpStat.BackIterator().Key()))
			fmt.Println(a, b)
		}
		tmpStat.Reset()
	}

}

func CheckRealText(iter *stat.Iterator, needs, lookFor int) bool {
	mostPopular := map[string]bool{
		"o": true,
		"и": true,
		"а": true,
		"е": true,
		"н": true,
		"т": true,
		"с": true,
		"р": true,
	}
	counter := 0
	for i := 0; i < lookFor; i++ {
		str, _ := alphabet.SingleMonogramToString(iter.Key())
		if mostPopular[str] == true {
			counter++
		}
		iter.Next()
	}
	return counter >= needs
}

func handleByPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func bruteforceKeys(cipherIter, cleanIter *stat.Iterator, r *ring.Ring, topLimit int) [][2]int {
	topClean := make([]int, topLimit)
	topCipher := make([]int, topLimit)
	for i := 0; i < topLimit && cipherIter.Valid() && cleanIter.Valid(); i++ {
		topCipher[i] = cipherIter.Key()
		topClean[i] = cleanIter.Key()
		cipherIter.Next()
		cleanIter.Next()
	}
	res := [][2]int{}
	// needs better cycle
	for i := 0; i < topLimit; i++ {
		for j := i + 1; j < topLimit; j++ {
			for l := 0; l < topLimit; l++ {
				for k := l + 1; k < topLimit; k++ {
					Y1 := topCipher[i]
					Y2 := topCipher[j]
					X1 := topClean[l]
					X2 := topClean[k]
					YF := (Y1 - Y2 + r.Mod()) % r.Mod()
					XF := (X1 - X2 + r.Mod()) % r.Mod()
					aH := r.SolveCongruence(XF, YF)
					if len(aH) == 0 {
						continue
					}
					a := aH[0]
					b := (Y1 - a*X1 + r.Mod()*a) % r.Mod()
					res = append(res, [2]int{a, b})
				}
			}
		}
	}
	return res
}
