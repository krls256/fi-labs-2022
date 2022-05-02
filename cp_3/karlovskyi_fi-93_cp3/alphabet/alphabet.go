package alphabet

import (
	"cp_3/stat"
	"strings"
)

var alphabet = []string{"а", "б", "в", "г", "д", "е", "ж", "з", "и", "й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ы", "ь", "э", "ю", "я"}
var reverseAlphabet = map[string]int{}
var AlpLen = len(alphabet)
var alpBiLen = AlpLen * AlpLen

func init() {
	for i, s := range alphabet {
		reverseAlphabet[s] = i
	}
}

func MonogramsToString(text []int) (string, error) {
	str := ""
	for _, b := range text {
		if b > AlpLen {
			return "", IntIsOutOfRangeErr
		}
		str += alphabet[b]
	}
	return str, nil
}

func BigramsToString(text []int) (string, error) {
	str := strings.Builder{}

	for _, b := range text {
		if b > alpBiLen {
			return "", IntIsOutOfRangeErr
		}
		c := b % AlpLen
		a := (b - c) / AlpLen
		str.Write([]byte(alphabet[a]))
		str.Write([]byte(alphabet[c]))
	}
	return str.String(), nil
}

func StringToMonograms(str string) ([]int, error) {
	res := make([]int, 0)
	tmp := ""
	for _, r := range str {
		tmp = string(r)
		if _, ok := reverseAlphabet[tmp]; !ok {
			return nil, UnacceptableSymbolErr
		}
		res = append(res, reverseAlphabet[tmp])
	}

	return res, nil
}

func BigramsToMonograms(bitext []int) ([]int, error) {
	var mono []int

	for _, b := range bitext {
		if b > alpBiLen {
			return nil, IntIsOutOfRangeErr
		}
		c := b % AlpLen
		a := (b - c) / AlpLen
		mono = append(mono, a, c)
	}
	return mono, nil
}

func StringToBigrams(str string) ([]int, error) {
	mono, err := StringToMonograms(str)
	if err != nil {
		return nil, err
	}
	res := make([]int, (len(mono)/2)+(len(mono)%2))
	second := 0
	for i := 0; i < len(res); i++ {
		if len(mono) > i*2+1 {
			second = mono[i*2+1]
		} else {
			second = 0
		}
		res[i] = mono[i*2]*AlpLen + second
	}
	return res, nil
}

func SingleMonogramToString(monogram int) (string, error) {
	if monogram > AlpLen {
		return "", IntIsOutOfRangeErr
	}
	return alphabet[monogram], nil
}

func SingleBigramToString(bigram int) (string, error) {
	if bigram > alpBiLen {
		return "", IntIsOutOfRangeErr
	}
	c := bigram % AlpLen
	a := (bigram - c) / AlpLen
	return alphabet[a] + alphabet[c], nil
}

func CheckRealText(iter *stat.Iterator, needs, lookFor int) bool {
	mostPopular := map[string]bool{
		"o": true,
		"и": true,
		"а": true,
		"е": true,
		"н": true,
		"т": true,
		"с": true,
		"р": true,
	}
	counter := 0
	for i := 0; i < lookFor; i++ {
		str, _ := SingleMonogramToString(iter.Key())
		if mostPopular[str] == true {
			counter++
		}
		iter.Next()
	}
	return counter >= needs
}
