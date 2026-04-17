/**
 * Definition for a binary tree node.
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	var c int

	var dfs func(node *TreeNode, max int)

	dfs = func(node *TreeNode, max int) {
		if node == nil {
			return //return without processing
		}

		if node.Val >= max {
			max = node.Val
			c++
		}

		dfs(node.Left, max)
		dfs(node.Right, max)
	}

	dfs(root, root.Val) // put root.Val since root value is always good.

	return c
}
