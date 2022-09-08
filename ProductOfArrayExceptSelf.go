package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Printf("Result:%d", result)
}

func productExceptSelf(nums []int) []int {
	var answer []int
	temp := 1
	for i := 0; i < len(nums); i++ {
		answer = append(answer, temp)
		temp *= nums[i]
	}
	temp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= temp
		temp *= nums[i]
	}
	return answer
}
