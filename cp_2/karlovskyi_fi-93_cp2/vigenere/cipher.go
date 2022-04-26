package vigenere

import "fmt"

type Cipher struct {
	Alphabet         []rune
	ReversedAlphabet map[rune]int
}

func NewCipher(Alphabet []rune) *Cipher {
	c := Cipher{
		Alphabet:         Alphabet,
		ReversedAlphabet: map[rune]int{},
	}
	for i, r := range Alphabet {
		c.ReversedAlphabet[r] = i
	}
	return &c
}

func (c *Cipher) Enc(text, key []rune) []rune {
	newText := make([]rune, len(text))
	for i, r := range text {
		if _, ok := c.ReversedAlphabet[r]; !ok {
			panic(fmt.Sprintf("%v is not allowed", r))
		}
		if _, ok := c.ReversedAlphabet[key[i%len(key)]]; !ok {
			panic(fmt.Sprintf("%v is not allowed", key[i%len(key)]))
		}
		newText[i] = c.Alphabet[(c.ReversedAlphabet[r]+c.ReversedAlphabet[key[i%len(key)]])%len(c.Alphabet)]
	}
	return newText
}

func (c *Cipher) Dec(text []rune, key []int) []rune {
	newText := make([]rune, len(text))
	for i, r := range text {
		newText[i] = c.Alphabet[(c.ReversedAlphabet[r]-key[i%len(key)]+len(c.Alphabet))%len(c.Alphabet)]
	}
	return newText
}

func (c *Cipher) IntToRuneForKey(key []int) []rune {
	newKey := make([]rune, len(key))
	for i, val := range key {
		newKey[i] = c.Alphabet[val]
	}
	return newKey
}
