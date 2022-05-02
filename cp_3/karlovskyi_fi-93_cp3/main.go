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
	fmt.Println("Після пошуку ключів, та їх фільтрації бачимо фрагменти:")
	for _, key := range keys {
		aN, err := r.Inverse(key.A)
		if err != nil {
			continue
		}
		deciphered := cipher.Dec(cipherBigram, r, cipher.Key{aN, key.B})
		monograms, err := alphabet.BigramsToMonograms(deciphered)
		handleByPanic(err)

		tmpStat.Append(monograms)
		tmpIter := tmpStat.BackIterator()
		if alphabet.CheckRealText(tmpIter, 4, 5) {
			contentToCheck, err := alphabet.MonogramsToString(monograms[:50])
			handleByPanic(err)

			fmt.Printf("Фрагмент: %v,  Ключі: (%v, %v)\n", contentToCheck, key.A, key.B)
		}
		tmpStat.Reset()
	}
	realKey := cipher.Key{72, 805}
	aN, _ := r.Inverse(realKey.A)
	realDecKey := cipher.Key{aN, realKey.B}
	fmt.Printf("Після аналізу фрагменту маємо справжній ключ: (%v, %v)\nПовний текст наведено нижче\n", realKey.A, realKey.B)
	str, _ := alphabet.BigramsToString(cipher.Dec(cipherBigram, r, realDecKey))
	fmt.Println(str)
}

func handleByPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func bruteforceKeys(cipherIter, cleanIter *stat.Iterator, r *ring.Ring, topLimit int) []cipher.Key {
	topClean := make([]int, topLimit)
	topCipher := make([]int, topLimit)
	for i := 0; i < topLimit && cipherIter.Valid() && cleanIter.Valid(); i++ {
		topCipher[i] = cipherIter.Key()
		topClean[i] = cleanIter.Key()
		cipherIter.Next()
		cleanIter.Next()
	}
	var res []cipher.Key
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
					res = append(res, cipher.Key{a, b})
				}
			}
		}
	}
	return res
}
