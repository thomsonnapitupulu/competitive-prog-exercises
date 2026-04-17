package lc75

func increasingTriplet(nums []int) bool {
	maxV := (1 << 31) - 1
	fst, snd := maxV, maxV
	for i := 0; i < len(nums); i++ {
		if nums[i] < fst {
			fst = nums[i]
		} else if nums[i] < snd && nums[i] > fst {
			snd = nums[i]
		} else {
			if fst != maxV && snd != maxV && nums[i] > fst && nums[i] > snd {
				return true
			}
		}
	}

	return false
}
