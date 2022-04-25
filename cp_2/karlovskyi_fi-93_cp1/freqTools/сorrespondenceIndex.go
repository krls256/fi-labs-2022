package freqTools

func CorrespondenceIndex(text []rune) float64 {
	counter := countFrequencies(text)
	tmp := 0.0
	for _, value := range counter {
		tmp += float64(value) * (float64(value) - 1)
	}

	return tmp / (float64(len(text)) * (float64(len(text)) - 1))
}

func CorrespondenceIndexExpectation(text []rune) float64 {
	counter := countFrequencies(text)
	n := float64(len(text))
	res := 0.0
	for _, value := range counter {
		res += (float64(value) / n) * (float64(value) / n)
	}
	return res
}

func countFrequencies(text []rune) map[rune]int {
	counter := map[rune]int{}
	for _, r := range text {
		counter[r]++
	}
	return counter
}
