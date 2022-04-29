package cipher

import "cp_3/ring"

func Enc(text []int, r *ring.Ring, a, b int) []int {
	cipherText := make([]int, len(text))
	for i := 0; i < len(cipherText); i++ {
		cipherText[i] = (a*text[i] + b) % r.Mod()
	}
	return cipherText
}

func Dec(cipherText []int, r *ring.Ring, aN, b int) []int {
	text := make([]int, len(cipherText))
	for i := 0; i < len(text); i++ {
		text[i] = ((aN*(cipherText[i]-b))%r.Mod() + r.Mod()) % r.Mod()
	}
	return text
}
