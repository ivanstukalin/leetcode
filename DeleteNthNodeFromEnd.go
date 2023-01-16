package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    tmp := head
    count := 1
    for tmp.Next != nil {
        count++
        tmp = tmp.Next
    }
    tmp = head
    prev := tmp
    counter := 0

    if count-n == 0 {
        return head.Next
    }

    for counter <= count-n {
        if count-n == counter {
            next := tmp.Next
            if next == nil && count == 1 {
                head = nil
                break
            }
            prev.Next = tmp.Next
            break
        }
        prev = tmp
        tmp = tmp.Next
        counter++
    }
	
    return head
}