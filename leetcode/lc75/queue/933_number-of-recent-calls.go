type RecentCounter struct {
	reqs []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.reqs = append(this.reqs, t)

	for this.reqs[0] < t-3000 {
		this.reqs = this.reqs[1:]
	}

	return len(this.reqs)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */