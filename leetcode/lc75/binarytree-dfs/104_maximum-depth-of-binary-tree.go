package lc75

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	farRight := maxDepth(root.Right)
	farLeft := maxDepth(root.Left)

	return 1 + max(farRight, farLeft)
}
