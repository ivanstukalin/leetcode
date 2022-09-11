package main

import (
	"fmt"
	"reflect"
)

func main() {
	testProductOfArrayExceptSelf()
	testSpiralMatrix()
	testAddTwoNumbersRecursive()
}

func testProductOfArrayExceptSelf() {
	nums := []int{1, 2, 3, 4}
	answer, result := []int{24, 12, 8, 6}, productExceptSelf(nums)
	var isResultCorrect bool
	isResultCorrect = reflect.DeepEqual(answer, result)
	fmt.Printf("PRODUCT OF ARRAY EXCEPT SELF: Result:%d, desicion:%t ", result, isResultCorrect)
	println("\n")
}

func testSpiralMatrix() {
	example := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	answer, result := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder(example)
	var isResultCorrect bool
	isResultCorrect = reflect.DeepEqual(answer, result)
	fmt.Printf("SPIRAL MATRIX: Result:%d, desicion:%t", result, isResultCorrect)
	println("\n")
}

func testAddTwoNumbersRecursive()  {
	var example1 *ListNode = new(ListNode)
	var example2 *ListNode = new(ListNode)
	example1.Val = 2
	example1.Next = new(ListNode)
	example1.Next.Val = 4
	example1.Next.Next = new(ListNode)
	example1.Next.Next.Val = 3
	example2.Val = 5
	example2.Next = new(ListNode)
	example2.Next.Val = 6
	example2.Next.Next = new(ListNode)
	example2.Next.Next.Val = 4
	result := addTwoNumbersRecursive(example1, example2)
	fmt.Printf("ADD TWO NUMBERS RECURSIVE: Result:[%d,%d,%d]", result.Val, result.Next.Val, result.Next.Next.Val)
	println("\n")
}
