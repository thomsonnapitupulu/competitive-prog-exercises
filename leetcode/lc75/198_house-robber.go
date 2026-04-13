package lc75

func rob(nums []int) int {
	memo := map[int]int{}
	var fn func(i int) int
	fn = func(i int) int {
		if v, ok := memo[i]; ok {
			return v
		}
		// base case = reach the end
		if i >= len(nums) {
			return 0
		}

		res := max(nums[i]+fn(i+2), fn(i+1))
		memo[i] = res
		return res
	}
	return fn(0)
}
