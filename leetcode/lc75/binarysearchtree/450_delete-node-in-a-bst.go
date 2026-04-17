package lc75

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}

		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		root.Val = findMin(root.Right)
		root.Right = deleteNode(root.Right, root.Val)
	}

	return root
}

func findMin(node *TreeNode) int {
	minVal := node.Val

	for node != nil {
		if node.Val < minVal {
			minVal = node.Val
		}
		node = node.Left
	}

	return minVal
}
