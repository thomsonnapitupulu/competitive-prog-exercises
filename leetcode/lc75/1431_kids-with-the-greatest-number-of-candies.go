package main

import (
	"slices"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max_c := int(slices.Max(candies))

	result := make([]bool, len(candies))
	for i, c := range candies {
		result[i] = ((c + extraCandies) >= max_c)
	}

	return result
}
