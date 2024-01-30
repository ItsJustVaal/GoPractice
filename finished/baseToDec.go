package main

import (
	"fmt"
)

// Issue I had here was I was using the letter's ascii value
// instead of the hex value. Once changed worked fine.
func BaseToDec(numberToChage string, base int) int {
	sum := 0
	power := 1

	for i := len(numberToChage) - 1; i >= 0; i-- {
		var num int
		fmt.Sscanf(string(numberToChage[i]), "%X", &num)

		sum += power * num
		power *= base
	}
	return sum
}
