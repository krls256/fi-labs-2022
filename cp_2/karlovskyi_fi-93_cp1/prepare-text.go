package main

import (
	"io"
	"os"
	"regexp"
	"strings"
)

var reg = regexp.MustCompile(`[^а-яА-Я]`)

func main() {
	r, _ := os.ReadFile("./graf-monte-kristo.txt")
	str := string(r)
	str = strings.ToLower(str)
	str = strings.Replace(str, "ё", "е", -1)
	str = reg.ReplaceAllString(str, "")
	dst, _ := os.Create("./text.txt")
	io.Copy(dst, strings.NewReader(str))
}
