package lc75

import "sort"

func closeStrings(word1 string, word2 string) bool {
	count := func(w string) (ch, freq [26]int) {
		for _, v := range w {
			ch[v-'a'] = 1
			freq[v-'a'] += 1
		}
		sort.Ints(freq[:])
		return ch, freq
	}

	ch1, freq1 := count(word1)
	ch2, freq2 := count(word2)

	return ch1 == ch2 && freq1 == freq2

}
