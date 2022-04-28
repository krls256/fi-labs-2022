package main

import (
	"cp_3/alphabet"
	"cp_3/ring"
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
	r, err := ring.NewRing(alphabet.AlpLen)
	handleByPanic(err)
	fmt.Println(r.SolveCongruence(4, 46))
}

func handleByPanic(err error) {
	if err != nil {
		panic(err)
	}
}
