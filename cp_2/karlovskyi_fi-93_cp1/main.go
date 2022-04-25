package main

import (
	"cp_2/freqTools"
	"cp_2/vigenere"
	"fmt"
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
	fmt.Println(freqTools.CorrespondenceIndexExpectation(text))
}
