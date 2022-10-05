package main

func plusOne(digits []int) []int {
	remainder := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if i < len(digits)-1 && remainder == 0 {
			break
		}
		current := digits[i] + remainder
		if current > 9 {
			remainder = 1
			digits[i] = 0
		} else {
			remainder = 0
			digits[i] = current
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
