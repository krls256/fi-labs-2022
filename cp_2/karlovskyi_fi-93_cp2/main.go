package main

import (
	"cp_2/freqTools"
	"cp_2/vigenere"
	"errors"
	"fmt"
	"math"
	"os"
)

var alphabet = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}

var key2 = []rune("пу")
var key3 = []rune("щсу")
var key4 = []rune("иомл")
var key5 = []rune("цигкц")
var key10 = []rune("овлшытсбих")
var key20 = []rune("икуитокшщиощушитиуки")

var cipherTexts = map[int][]rune{}

func main() {
	textNotPrepared, err := os.ReadFile("./text.txt")
	if err != nil {
		panic(err)
	}
	text := []rune(string(textNotPrepared))
	c := vigenere.NewCipher(alphabet)
	cipherTexts[2] = c.Enc(text, key2)
	cipherTexts[3] = c.Enc(text, key3)
	cipherTexts[4] = c.Enc(text, key4)
	cipherTexts[5] = c.Enc(text, key5)
	cipherTexts[10] = c.Enc(text, key10)
	cipherTexts[20] = c.Enc(text, key20)

	cipherBytes, err := os.ReadFile("./cipher.txt")
	cipherText := []rune(string(cipherBytes))
	keyLen, _ := findKeyLen(cipherText, 20)
	key := findKey(cipherText, keyLen, freqTools.CountFrequencies(text))
	key[13] = c.ReversedAlphabet['и']
	key[15] = c.ReversedAlphabet['н']
	fmt.Println(string(c.Dec(cipherText, key)))
	fmt.Println(string(c.IntToRuneForKey(key)))
}

func findKey(cipherText []rune, keyLen int, defaultDistribution *freqTools.LanguageDistribution) []int {
	key := make([]int, keyLen)
	split := freqTools.SplitByKeyLen(cipherText, keyLen)
	for i := 0; i < keyLen; i++ {
		distr := freqTools.CountFrequencies(split[i])
		iterator := distr.Tree.Reverse()
		key[i] = int((iterator.Value() - defaultDistribution.Tree.Reverse().Value() + rune(len(alphabet))) % rune(len(alphabet))) /// fix
	}
	return key
}

func findKeyLen(cipherText []rune, keyLenLimit int) (int, error) {
	baseKeyLen := 2
	correspondenceIndexDelta := 1.0 / 100 // 1%
	stat := make([]float64, keyLenLimit-baseKeyLen+1)
	for i := baseKeyLen; i <= keyLenLimit; i++ {
		stat[i-baseKeyLen] = freqTools.CorrespondenceIndex(freqTools.SplitByKeyLen(cipherText, i)[0])
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
