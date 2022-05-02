package cipher

import "cp_3/ring"

type Key struct {
	A, B int
}

func Enc(text []int, r *ring.Ring, key Key) []int {
	cipherText := make([]int, len(text))
	for i := 0; i < len(cipherText); i++ {
		cipherText[i] = (key.A*text[i] + key.B) % r.Mod()
	}
	return cipherText
}

func Dec(cipherText []int, r *ring.Ring, kenN Key) []int {
	text := make([]int, len(cipherText))
	for i := 0; i < len(text); i++ {
		text[i] = ((kenN.A*(cipherText[i]-kenN.B))%r.Mod() + r.Mod()) % r.Mod()
	}
	return text
}
