package main

import (
	"fmt"
	"unicode"
)

func caeser() {
	testStr := "abcdefghijklmnopqrstuvwxyz"
	rotate := 2
	newStr := ""

	for _, x := range testStr {
		switch {
		case unicode.IsLetter(x) && unicode.IsUpper(x):
			newStr += string(rune((int(x)+rotate-65)%26 + 65))
		case unicode.IsLetter(x) && unicode.IsLower(x):
			newStr += string(rune((int(x)+rotate-97)%26 + 97))
		default:
			newStr += string(x)
		}
	}

	fmt.Println(newStr)
}
