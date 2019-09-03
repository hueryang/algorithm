package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left >= right {
		depth = left + 1
	} else {
		depth = right + 1
	}
	return depth
}