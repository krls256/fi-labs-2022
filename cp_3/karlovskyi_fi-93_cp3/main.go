package main

import (
	"cp_3/alphabet"
	"fmt"
	"os"
)

func main() {
	textB, _ := os.ReadFile("./cipherText")
	text := string(textB)

	fmt.Println(text)
	bi, _ := alphabet.StringToBigrams(text)
	fmt.Println(alphabet.BigramsToString(bi))
}
