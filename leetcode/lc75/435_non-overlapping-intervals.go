
func eraseOverlapIntervals(intervals [][]int) int {
	//the idea is, we should sort them by interval end ascendingly
	//and then iterate from the left, make sure the end of interval i is lower than start of interval i+1

	//sort the interval end ascendingly
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	lastEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < lastEnd {
			count++
		} else {
			//lastEnd only slides when it is not overlapping
			//if it is overlapping, count++ and no slide bcs we assume we just removed that interval
			lastEnd = intervals[i][1]
		}
	}

	return count
}
