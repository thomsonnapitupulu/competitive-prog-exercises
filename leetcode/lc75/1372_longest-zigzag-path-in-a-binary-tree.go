package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const GOLEFT = false
const GORIGHT = true

func longestZigZag(root *TreeNode) int {
	return max(dfsCount(root.Left, GOLEFT, 0), dfsCount(root.Right, GORIGHT, 0))
}

func dfsCount(node *TreeNode, wentRight bool, counter int) int {
	if node == nil {
		return counter
	}

	if wentRight {
		return max(dfsCount(node.Left, GOLEFT, counter+1), dfsCount(node.Right, GORIGHT, 0))
	}

	return max(dfsCount(node.Left, GOLEFT, 0), dfsCount(node.Right, GORIGHT, counter+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
