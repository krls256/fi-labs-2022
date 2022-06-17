package main

import (
	"cp_2/freqTools"
	"cp_2/vigenere"
	"fmt"
	"golang.org/x/exp/rand"
	"os"
)

var alphabet = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}

var key2 = []rune("пу")
var key3 = []rune("щсу")
var key4 = []rune("иомл")
var key5 = []rune("цигкц")
var key10 = []rune("овлшытсбих")
var key20 = []rune("икуитокшщиощушитиуки")
var keys [][]rune

func init() {
	for i := 6; i <= 30; i++ {
		key := []rune{}
		for j := 0; j < i; j++ {
			key = append(key, alphabet[rand.Intn(len(alphabet))])
		}
		keys = append(keys, key)
	}
}

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
	keyLen, _ := freqTools.FindKeyLen(cipherText, 20)

	fragmentSize := 50

	fmt.Printf("Завдання №1\n------------\n")
	fmt.Printf("Базовий текст: %v...\n", string(text[:fragmentSize]))
	fmt.Printf("Текст зашифрованый ключем довжини 2: %v...\n", string(cipherTexts[2][:fragmentSize]))
	fmt.Printf("Текст зашифрованый ключем довжини 3: %v...\n", string(cipherTexts[3][:fragmentSize]))
	fmt.Printf("Текст зашифрованый ключем довжини 4: %v...\n", string(cipherTexts[4][:fragmentSize]))
	fmt.Printf("Текст зашифрованый ключем довжини 5: %v...\n", string(cipherTexts[5][:fragmentSize]))
	fmt.Printf("\nЗавдання №2\n------------\n")
	fmt.Printf("Індекси відповідності для тексту: %1.4f\n", freqTools.CorrespondenceIndex(text))

	fmt.Printf("Індекси відповідності для ключа довжини 2: %1.4f\n", freqTools.CorrespondenceIndex(cipherTexts[2]))
	fmt.Printf("Індекси відповідності для ключа довжини 3: %1.4f\n", freqTools.CorrespondenceIndex(cipherTexts[3]))
	fmt.Printf("Індекси відповідності для ключа довжини 4: %1.4f\n", freqTools.CorrespondenceIndex(cipherTexts[4]))
	fmt.Printf("Індекси відповідності для ключа довжини 5: %1.4f\n", freqTools.CorrespondenceIndex(cipherTexts[5]))
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		fmt.Printf("Індекси відповідності для ключа довжини %v: %1.4f, ключ: %v\n", len(key), freqTools.CorrespondenceIndex(c.Enc(text, key)), string(key))
	}

	fmt.Printf("\nЗавдання №3\n------------\n")
	fmt.Printf("Довжина ключа: %v\n", keyLen)
	key := freqTools.FindKey(cipherText, keyLen, freqTools.CountFrequencies(text), alphabet)
	fmt.Printf("Базовий ключ: %v\n", string(c.IntToRuneForKey(key)))
	fmt.Printf("Результат дешифрування: %v...\n", string(c.Dec(cipherText, key)[:fragmentSize]))
	key[13] = c.ReversedAlphabet['и']
	key[15] = c.ReversedAlphabet['н']
	fmt.Printf("Результат після корегування: %v...\n", string(c.Dec(cipherText, key)[:fragmentSize]))
	key = freqTools.FindKeyByM(cipherText, keyLen, freqTools.ReverseCountFrequencies(text), c)
	fmt.Printf("Базовий ключ отриманий через M(g): %v\n", string(c.IntToRuneForKey(key)))
	fmt.Printf("Результат з ключем отриманим через M(g): %v...\n", string(c.Dec(cipherText, key)[:fragmentSize]))

	fmt.Println("-------------------")
	fmt.Printf("Повністю розшифрований текст: %v\n", string(c.Dec(cipherText, key)))
}
