package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	ans := []int{}

	//while queue is not empty
	for len(queue) > 0 {
		qLen := len(queue)
		//iterate for each level
		for i := 0; i < qLen; i++ {
			node := queue[0]
			queue = queue[1:] //pop the queue

			//if it is in the rightest node in the same level
			if i == qLen-1 {
				//then put into answer
				ans = append(ans, node.Val)
			}

			//fill the queue with children of current nodes
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return ans
}

// n = 1
//     i = 0
//         node = A
//         q = []
//         ans = [A]
//         q = [C]
//         q = [C, B]
// n = 2
//     i = 0
//         node = C
//         q = [B]
//         q = [B, F, E]
//     i = 1
//         node = B
//         q = [F, E]
//         ans = [A, B]
//         q = [F, E, D]
// n = 3
//     i = 0
//         node = F
//         q = [E, D]
//         q = [E, D, J, I]
//     i = 1
//         node = E
//         q = [D, J, I]
//         q = [D, J, I]
//     i = 2
//         node = D
//         q = [J, I]
//         ans = [A, B, D]
//         q = [J, I, H, G]
// n = 4
//     i = 0
//         node = J
//         q = [I, H, G]
//         q = [I, H, G, M, L]
//     i = 1
//         node = I
//         q = [H, G, M, L]
//     i = 2
//         node = H
//         q = [G, M, L]
//         q = [G, M, L, K]
//     i = 3
//         node = G
//         ans = [A, B, D, G]
//         q = [M, L, K]
// n = 3
//     i = 0
//         node = M
//         q = [L, K]
//         q = [L, K, N]
//     i = 1
//         node = L
//         q = [K, N]
//     i = 2
//         node = K
//         ans = [A, B, D, G, K]
//         q = [N]
// n = 1
//     i = 0
//         node = N
//         ans = [A, B, D, G, K, N]
