package main

func ReverseInteger(x int) int {
	palendrome := 0
	for x/10 > 0 || -x/10 > 0 {
		temp := x % 10
		palendrome = palendrome*10 + temp
		x = x / 10
	}

	palendrome = palendrome*10 + x
	if palendrome > 2147483647 || palendrome < -2147483648 {
		return 0
	}
	return palendrome
}
