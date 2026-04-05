func maxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'i': true,
		'u': true,
		'e': true,
		'o': true,
	}

	maxCount, count := 0, 0
	l, r := 0, 0

	for r < len(s) {
		//grow window
		if _, ok := vowels[s[r]]; ok {
			count++
		}
		r++

		if r-l == k {
			maxCount = max(maxCount, count)

			//ceiling logic so that counter not exceeding the window length (k)
			if _, ok := vowels[s[l]]; ok {
				count--
			}
			l++
		}
	}

	return maxCount
}