package main

func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0

	for i < len(version1) || j < len(version2) {
		var num1, num2 int

		for i < len(version1) && version1[i] != '.' {
			num1 = num1*10 + int(version1[i]-'0')
			i++
		}

		for j < len(version2) && version2[j] != '.' {
			num2 = num2*10 + int(version2[j]-'0')
			j++
		}

		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}

		if i < len(version1) {
			i++
		}
		if j < len(version2) {
			j++
		}
	}

	return 0
}
