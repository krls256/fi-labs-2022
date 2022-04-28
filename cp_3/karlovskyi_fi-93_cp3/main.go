package main

import (
	"cp_3/alphabet"
	"cp_3/stat"
	"fmt"
	"os"
)

func main() {
	textB, err := os.ReadFile("./cipherText")
	handleByPanic(err)
	cipherText := string(textB)
	bi, err := alphabet.StringToBigrams(cipherText)
	handleByPanic(err)
	st := stat.Stat{}

	st.Append(bi)
	for iter := st.BackIterator(); iter.Valid(); iter.Next() {
		fmt.Printf("%v: %v\n", iter.Key(), iter.Value())
	}

}

func handleByPanic(err error) {
	if err != nil {
		panic(err)
	}
}
