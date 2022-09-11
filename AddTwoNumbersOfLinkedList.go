package main

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbersRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	var val1 int = 0
	var val2 int = 0
	var result *ListNode = new(ListNode)
	var next1 *ListNode
	var next2 *ListNode
	if l1 != nil {
		val1 = l1.Val
		next1 = l1.Next
	}
	if l2 != nil {
		val2 = l2.Val
		next2 = l2.Next
	}
	if next1 == nil && next2 == nil {
		result.Val = val1 + val2
		if result.Val > 9 {
			result.Val = result.Val - 10
			var temp *ListNode = new(ListNode)
			temp.Val = 1
			result.Next = temp
		}
		return result
	} else {
		result = addTwoNumbersRecursive(next1, next2)
		var next *ListNode = new(ListNode)
		next.Val = val1 + val2
		if next.Val > 9 {
			next.Val = next.Val - 10
			result.Val++
			if result.Val > 9 {
				result.Val = result.Val - 10
				var temp *ListNode = new(ListNode)
				temp.Val = 1
				result.Next = addTwoNumbersRecursive(result.Next, temp)
			}
		}
		next.Next = result
		result = next
	}
	return result
}
