package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return sumUp(root, 0, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func sumUp(root *TreeNode, pre int, sumTarget int) int {
	if root == nil {
		return 0
	}

	sum := pre + root.Val

	c := 0

	if sum == sumTarget {
		c = 1
	}

	return c + sumUp(root.Left, sum, sumTarget) + sumUp(root.Right, sum, sumTarget)
}
