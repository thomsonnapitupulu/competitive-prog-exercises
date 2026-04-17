package twopointers

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	i, j := 0, len(nums)-1
	for i < j { //while i < j
		temp := nums[i] + nums[j]
		if temp == k {
			i++
			j--
			ans++
		} else if temp > k {
			j--
		} else {
			i++
		}
	}

	return ans
}
