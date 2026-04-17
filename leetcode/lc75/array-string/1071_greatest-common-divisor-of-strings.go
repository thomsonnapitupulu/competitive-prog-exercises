package arraystring

func gcdOfStrings(str1 string, str2 string) string {
	if str2+str1 != str1+str2 {
		return ""
	}
	gcdLen := gcd(len(str1), len(str2))
	return str1[:gcdLen]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
