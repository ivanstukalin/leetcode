package main

import (
	"fmt"
	"reflect"
)

func main() {
	testProductOfArrayExceptSelf()
	TestSpiralMatrix()
}

func testProductOfArrayExceptSelf() {
	nums := []int{1, 2, 3, 4}
	answer, result := []int{24, 12, 8, 6}, productExceptSelf(nums)
	var isResultCorrect bool
	isResultCorrect = reflect.DeepEqual(answer, result)
	fmt.Printf("PRODUCT OF ARRAY EXCEPT SELF: Result:%d, desicion:%t ", result, isResultCorrect)
	println("\n")
}

func TestSpiralMatrix() {
	example := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	answer, result := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder(example)
	var isResultCorrect bool
	isResultCorrect = reflect.DeepEqual(answer, result)
	fmt.Printf("SPIRAL MATRIX: Result:%d, desicion:%t", result, isResultCorrect)
	println("\n")
}
