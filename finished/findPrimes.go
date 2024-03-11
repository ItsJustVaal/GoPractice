package main

func findPrimes(primeNums []int, num int) []int {
	var answer []int
	for _, number := range primeNums {
		for num%number == 0 {
			num /= number
			answer = append(answer, number)
		}
	}
	if num > 1 {
		answer = append(answer, num)
	}
	return answer
}