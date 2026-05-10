package main

func expand(s string, left int, right int) string {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left -= 1
		right += 1
	}

	return s[left+1 : right]

}

func longestPalindrome(s string) string {

	max := ""

	for i, _ := range s {
		r1 := expand(s, i, i)
		r2 := expand(s, i, i+1)
		if len(r1) > len(max) {
			max = r1
		}
		if len(r2) > len(max) {
			max = r2
		}
	}

	return max

}
