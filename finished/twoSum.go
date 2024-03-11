package main

func twoSum(nums []int, toFind int) (int, int) {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if _, ok := numsMap[num]; ok {
			return numsMap[num], i
		}
		numsMap[toFind-num] = i
	}
	return -1, -1
}