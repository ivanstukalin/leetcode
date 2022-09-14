package main

import (
	"fmt"
	"reflect"
)

func main() {
	testProductOfArrayExceptSelf()
	testSpiralMatrix()
	testAddTwoNumbersRecursive()
	oddEvenListTest()
	testValidParentheses()
	testKthSmallestElementInBST()
}

func testValidParentheses() {
	examples := []string{"()","[]","{}","[[[]]]","{[()]}","[{{]"}
	answers := []bool{true,true,true,true,true,false}
	result := -1
	for key, example := range examples {
		if isValid(example) != answers[key] {
			result = key
			break
		}
	}

	fmt.Printf("VALID PARENTHESES: ")
	if result == -1 {
		fmt.Println("All tests ok!")
	} else {
		fmt.Println("Test not okay!: ", result)
	}
}

func testKthSmallestElementInBST() {
	var examples TreeNode = TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}
	var k = map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	result := -1
	for checker, value := range k {
		res := kthSmallest(&examples, value)
		if res != checker {
			result = checker
			break
		}
	}
	fmt.Printf("KTH SMALLEST ELEMENT IN A BST: ")
	if result == -1 {
		fmt.Println("All tests ok!")
	} else {
		fmt.Println("Test not okay!: ", result)
	}
}

func testProductOfArrayExceptSelf() {
	nums := []int{1, 2, 3, 4}
	answer, result := []int{24, 12, 8, 6}, productExceptSelf(nums)
	var isResultCorrect bool
	isResultCorrect = reflect.DeepEqual(answer, result)
	fmt.Printf("PRODUCT OF ARRAY EXCEPT SELF: Result:%d, decision:%t ", result, isResultCorrect)
	println("\n")
}

func testSpiralMatrix() {
	example := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	answer, result := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, spiralOrder(example)
	var isResultCorrect bool
	isResultCorrect = reflect.DeepEqual(answer, result)
	fmt.Printf("SPIRAL MATRIX: Result:%d, decision:%t", result, isResultCorrect)
	println("\n")
}

func testAddTwoNumbersRecursive()  {
	var example1 ListNode = ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	var example2 ListNode = ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbersRecursive(&example1, &example2)
	fmt.Printf("ADD TWO NUMBERS RECURSIVE: Result:[%d,%d,%d]", result.Val, result.Next.Val, result.Next.Next.Val)
	println("\n")
}

func oddEvenListTest()  {
	var example ListNode = ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	var checker ListNode = ListNode{1, &ListNode{3, &ListNode{5, &ListNode{2, &ListNode{4, nil}}}}}
	result := oddEvenList(&example)
	var isResultCorrect bool
	fmt.Printf("ODD EVEN LIST: Result:[")
	for result != nil {
		fmt.Printf("%d", result.Val)
		isResultCorrect = result.Val == checker.Val
		if !isResultCorrect {
			result = nil
		}
		result = result.Next
		if checker.Next != nil {
			checker = *checker.Next
		}
	}
	fmt.Printf("], decision: %t", isResultCorrect)
	println("\n")
}
