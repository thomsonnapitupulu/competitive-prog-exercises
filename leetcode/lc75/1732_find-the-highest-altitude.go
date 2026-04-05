package lc75

func largestAltitude(gain []int) int {
	maxV, tmp := 0, 0
	for i := range gain {
		tmp += gain[i]
		maxV = max(tmp, maxV)
	}

	return maxV
}
