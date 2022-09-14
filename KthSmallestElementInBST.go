package main

// TreeNode /** Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var minVal int
	inOrderKthSmallest(root, &k, &minVal)
	return minVal
}

func inOrderKthSmallest(root *TreeNode, k *int, minVal *int)  {
	if root == nil { //check: is root leftmost
		return
	}
	inOrderKthSmallest(root.Left, k, minVal) //looking for min left value
	*k--
	if *k == 0 {
		*minVal = root.Val
		return
	}
	inOrderKthSmallest(root.Right, k, minVal) //if k > 0, we need to check right Node
}
