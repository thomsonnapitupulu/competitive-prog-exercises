package arraystring

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if len(flowerbed) < 3 {
		for _, v := range flowerbed {
			if v == 1 {
				return !(n > 0)
			}
		}
		return (n == 1)
	}
	for i := 0; i <= len(flowerbed)-2 && n > 0; i++ {
		if i == 0 || i == len(flowerbed)-2 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i+2] == 0 {
			flowerbed[i+1] = 1
			i++
			n--
		}
	}

	return !(n > 0)
}
