package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, notOdd := head, head.Next
	notOddFirst := notOdd
	for notOdd != nil && notOdd.Next != nil {
		odd.Next = notOdd.Next
		odd = odd.Next
		notOdd.Next = odd.Next
		notOdd = notOdd.Next
	}
	odd.Next = notOddFirst
	return head
}
