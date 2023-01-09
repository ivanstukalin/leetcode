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
	testSetMatrixZeroes()
	testGroupAnagrams()
	testRemoveDuplicates()
	testPlusOne()
	testTwoSum()
	testReverseInteger()
}

func testReverseInteger() {
	examples := []int{123, 456, 789, 123456, 1534236469}
	answers := []int{321, 654, 987, 654321, 0}
	errors := 0
	for key, value := range examples {
		result := ReverseInteger(value)
		if result != answers[key] {
			fmt.Println("Test failed! Input: ", value, " Expected: ", answers[key], " Result: ", result)
			errors++
		}
	}
	if errors > 0 {
		fmt.Println("Tests failed!")
	} else {
		fmt.Println("Tests passed!")
	}
}

func testTwoSum() {
	exampleNums := []int{2, 7, 11, 15}
	exampleTarget := 9
	test := twoSum(exampleNums, exampleTarget)
	fmt.Printf("Test two sum example %d: result: %d \n", exampleNums, test)
}

func testRemoveDuplicates() {
	example := []int{1, 2, 2, 3, 4}
	answer := 4
	test := removeDuplicates(example)
	fmt.Println("Test remove duplicates:", test == answer)
}

func testPlusOne() {
	example := []int{1, 2, 3, 3}
	answer := []int{1, 2, 3, 4}
	test := plusOne(example)
	fmt.Println("Test plus one:", test[len(test)-1] == answer[len(answer)-1])
}

func testValidParentheses() {
	examples := []string{"()", "[]", "{}", "[[[]]]", "{[()]}", "[{{]"}
	answers := []bool{true, true, true, true, true, false}
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
	var examples = TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}
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

func testGroupAnagrams() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	checker := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	output := groupAnagrams(input)
	fmt.Println("GROUP ANAGRAMS:", input, output, "correct result:", checker)
}

func testSetMatrixZeroes() {
	example := [][]int{{1, 1, 1}, {0, 1, 1}, {1, 1, 1}}
	result := setZeroes(example)
	fmt.Println("SET MATRIX ZEROES:", result)
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

func testAddTwoNumbersRecursive() {
	var example1 = ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	var example2 = ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbersRecursive(&example1, &example2)
	fmt.Printf("ADD TWO NUMBERS RECURSIVE: Result:[%d,%d,%d]", result.Val, result.Next.Val, result.Next.Next.Val)
	println("\n")
}

func oddEvenListTest() {
	var example = ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	var checker = ListNode{1, &ListNode{3, &ListNode{5, &ListNode{2, &ListNode{4, nil}}}}}
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
