package lc75

import "math"

//====== APPROACH 1: BFS + QUEUE ======
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	q := []*TreeNode{root}
	lvl, res := 1, 1

	maxSum := math.MinInt64 //large negative number, to make sure that default value is the most negative value
	// maxSum := -1 << 60 //we can also use this alternatively

	for len(q) > 0 {
		size := len(q)
		sum := 0

		for i := 0; i < size; i++ {
			//dequeue
			node := q[0]
			q = q[1:]

			sum += node.Val

			if node.Left != nil {
				//enqueue
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		if sum > maxSum {
			maxSum = sum
			res = lvl
		}
		lvl++
	}

	return res
}

//====== APPROACH 2: DFS + RECURSION ======
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sums := make(map[int]int)
	dfs(root, 1, sums)

	level := 0
	for l := range sums {
		if sums[l] > sums[level] {
			level = l
		}
	}
	return level
}

func dfs(node *TreeNode, level int, sums map[int]int) {
	if node == nil {
		return
	}

	dfs(node.Left, level+1, sums)
	sums[level] += node.Val
	dfs(node.Right, level+1, sums)
}
