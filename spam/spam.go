package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := ""
	lower := 0.0
	upper := 0.0
	underscore := 0.0
	symbol := 0.0

	fmt.Scanf("%s\n", &s)
	rs := []rune(s)
	for _, r := range rs {
		if unicode.IsLower(r) {
			lower++
		} else if unicode.IsUpper(r) {
			upper++
		} else if isUnderscore(r) {
			underscore++
		} else {
			symbol++
		}
	}

	sum := lower + upper + underscore + symbol

	fmt.Println(underscore / sum)
	fmt.Println(lower / sum)
	fmt.Println(upper / sum)
	fmt.Println(symbol / sum)
}

func isUnderscore(r rune) bool {
	return r == rune('_')
}
