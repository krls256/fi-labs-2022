package main

import (
	"bytes"
	"cp_1/pe"
	"flag"
	"fmt"
	"os"
)

var alphabet = []string{"А", "Б", "В", "Г", "Д", "Е", "Ё", "Ж", "З", "И", "Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х", "Ц", "Ч", "Ш", "Щ", "Ъ", "Ы", "Ь", "Э", "Ю", "Я"}
var alphabetWithSpace = []string{"А", "Б", "В", "Г", "Д", "Е", "Ё", "Ж", "З", "И", "Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х", "Ц", "Ч", "Ш", "Щ", "Ъ", "Ы", "Ь", "Э", "Ю", "Я", " "}
var isDetailedReview bool

func init() {
	flag.BoolVar(&isDetailedReview, "d", false, "Is final review must be detailed")
	flag.Parse()
}

func main() {
	// Читаємо текст з файлів та готуємо його для додавання у імовірносний ансамбль
	content, err := os.ReadFile("./martin-iden.txt")
	if err != nil {
		panic(err)
	}
	content = prepareContent(content)

	content2, err := os.ReadFile("./idiot.txt")
	if err != nil {
		panic(err)
	}
	content2 = prepareContent(content2)

	// Створюємо 2 ансамблі, і заповнюємо іх текстом
	eWithoutSpaces := pe.New(alphabet)
	eWithoutSpaces.AddText(string(content))
	eWithoutSpaces.AddText(string(content2))

	eWithSpaces := pe.New(alphabetWithSpace)
	eWithSpaces.AddTextWithSpaces(string(content))
	eWithSpaces.AddTextWithSpaces(string(content2))

	// отримаемо всі статастики, які сказано взяти у завдянні та форматуемо результат

	statByLettersWithoutSpaces := eWithoutSpaces.GetNGramStatWithIntersection(1)
	statByLettersWithSpaces := eWithSpaces.GetNGramStatWithIntersection(1)

	statByBigramsWithIntersAndWithoutSpaces := eWithoutSpaces.GetNGramStatWithIntersection(2)
	statByBigtamsWithIntersAndWithSpaces := eWithSpaces.GetNGramStatWithIntersection(2)

	statByBigramsWithoutIntersAndWithoutSpaces := eWithoutSpaces.GetNGramStatWithoutIntersection(2)
	statByBigtamsWithoutIntersAndWithSpaces := eWithSpaces.GetNGramStatWithoutIntersection(2)

	statByLettersWithSpaces.ForEachInOrder(writeReview(statByLettersWithSpaces, "statByLettersWithSpaces"))
	statByLettersWithoutSpaces.ForEachInOrder(writeReview(statByLettersWithoutSpaces, "statByLettersWithoutSpaces"))
	statByBigramsWithIntersAndWithoutSpaces.ForEachInOrder(writeReview(statByBigramsWithIntersAndWithoutSpaces, "statByBigramsWithIntersAndWithoutSpaces"))
	statByBigtamsWithIntersAndWithSpaces.ForEachInOrder(writeReview(statByBigtamsWithIntersAndWithSpaces, "statByBigtamsWithIntersAndWithSpaces"))
	statByBigramsWithoutIntersAndWithoutSpaces.ForEachInOrder(writeReview(statByBigramsWithoutIntersAndWithoutSpaces, "statByBigramsWithoutIntersAndWithoutSpaces"))
	statByBigtamsWithoutIntersAndWithSpaces.ForEachInOrder(writeReview(statByBigtamsWithoutIntersAndWithSpaces, "statByBigtamsWithoutIntersAndWithSpaces"))
}

func writeReview(stat *pe.NGramStat, statType string) func(string, int) {
	wasFirstLine := false
	return func(s string, i int) {
		if !wasFirstLine {
			fmt.Printf("%s has entropy: %v \n", statType, stat.Entropy())
			wasFirstLine = true
		}
		if isDetailedReview {
			fmt.Printf("%v occurs in %2.4f%% of cases\n", s, float64(i)/float64(stat.Total)*100)
		}
	}
}

func prepareContent(b []byte) []byte {
	return bytes.Trim(bytes.Replace(bytes.ToUpper(b), []byte("\n"), []byte(" "), -1), " ")
}
