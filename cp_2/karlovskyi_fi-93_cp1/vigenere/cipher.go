package vigenere

import "fmt"

type cipher struct {
	alphabet         []rune
	reversedAlphabet map[rune]int
}

func NewCipher(alphabet []rune) *cipher {
	c := cipher{
		alphabet:         alphabet,
		reversedAlphabet: map[rune]int{},
	}
	for i, r := range alphabet {
		c.reversedAlphabet[r] = i
	}
	return &c
}

func (c *cipher) Enc(text []rune, key []rune) []rune {
	newText := make([]rune, len(text))
	for i, r := range text {
		if _, ok := c.reversedAlphabet[r]; !ok {
			panic(fmt.Sprintf("%v is not allowed", r))
		}
		if _, ok := c.reversedAlphabet[key[i%len(key)]]; !ok {
			panic(fmt.Sprintf("%v is not allowed", key[i%len(key)]))
		}
		newText[i] = c.alphabet[(c.reversedAlphabet[r]+c.reversedAlphabet[key[i%len(key)]])%len(c.alphabet)]
	}
	return newText
}
