package main

//O(n)
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
