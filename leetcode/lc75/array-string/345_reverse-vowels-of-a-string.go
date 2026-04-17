package arraystring

import "strings"

func reverseVowels(s string) string {
	word := []rune(s)
	start := 0
	end := len(word) - 1
	vowels := "aeiouAEIOU"

	for start < end {
		for start < end && !strings.ContainsRune(vowels, word[start]) {
			start++
		}
		for start < end && !strings.ContainsRune(vowels, word[end]) {
			end--
		}

		word[start], word[end] = word[end], word[start]
		start++
		end--
	}

	return string(word)
}
