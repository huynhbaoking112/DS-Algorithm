package main

func lengthOfLongestSubstring(s string) int {

	max := 0
	left := 0
	mapS := map[byte]int{}

	for right := 0; right < len(s); right++ {
		if v, ok := mapS[s[right]]; ok && v >= left {
			left = v + 1
		}

		mapS[s[right]] = right

		currentLen := right - left + 1
		if currentLen > max {
			max = currentLen
		}

	}
	return max

}
