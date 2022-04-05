package main

import (
	"bytes"
	"cp_1/pe"
	"fmt"
	"os"
)

var alphabet = []string{"А", "Б", "В", "Г", "Д", "Е", "Ё", "Ж", "З", "И", "Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х", "Ц", "Ч", "Ш", "Щ", "Ъ", "Ы", "Ь", "Э", "Ю", "Я"}

func main() {
	content, err := os.ReadFile("./martin-iden.txt")
	if err != nil {
		panic(err)
	}
	content = bytes.Trim(bytes.Replace(bytes.ToUpper(content), []byte("\n"), []byte(" "), -1), " ")

	content2, err := os.ReadFile("./idiot.txt")
	if err != nil {
		panic(err)
	}
	content2 = bytes.Trim(bytes.Replace(bytes.ToUpper(content2), []byte("\n"), []byte(" "), -1), " ")
	eWithoutSpaces := pe.New(alphabet)
	eWithoutSpaces.AddText(string(content))
	eWithoutSpaces.AddText(string(content2))

	eWithSpaces := pe.New(alphabet)
	eWithSpaces.AddTextWithSpaces(string(content))
	eWithSpaces.AddTextWithSpaces(string(content2))

	stat1 := eWithoutSpaces.GetNGramStatWithIntersection(2)
	stat2 := eWithSpaces.GetNGramStatWithIntersection(2)
	fmt.Println(stat1.Entropy())
	fmt.Println(stat2.Entropy())
}
