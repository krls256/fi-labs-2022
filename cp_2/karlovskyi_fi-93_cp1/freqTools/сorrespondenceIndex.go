package freqTools

func CorrespondenceIndex(text []rune) float64 {
	counter := map[rune]int{}
	for _, r := range text {
		counter[r]++
	}
	tmp := 0.0
	for _, value := range counter {
		tmp += float64(value) * (float64(value) - 1)
	}

	return tmp / (float64(len(text)) * (float64(len(text)) - 1))
}
