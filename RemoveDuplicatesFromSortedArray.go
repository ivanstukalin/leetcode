package main

func removeDuplicates(nums []int) int {
	counter := 0
	for i := 1; i < len(nums); i++ {
		if  nums[counter] < nums[i] {
			counter++
			nums[counter] = nums[i]
		}
	}
	return counter+1
}