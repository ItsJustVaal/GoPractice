package main

import (
	"fmt"
	"unicode"
)

func phone() {
	numbers := []string{"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892"}
	newNums := []string{}

	for _, x := range numbers {
		tmpNum := ""
		for _, y := range x {
			if unicode.IsDigit(y) {
				tmpNum += string(y)
			}
		}
		newNums = append(newNums, tmpNum)
	}
	fmt.Println(newNums)
}
