package main

func twoSum(nums []int, target int) []int {
	var result []int
	temp := make(map[int]int)
	for key1, num1 := range nums {
		key2, num2 := temp[target-num1]
		if num2 {
			result = []int{key1, key2}
			break
		}
		temp[num1] = key1
	}
	return result
}
