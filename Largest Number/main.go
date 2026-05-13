package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	arrS := make([]string,len(nums))

	for i, v := range nums {
		arrS[i] =  strconv.Itoa(v)
	}

	sort.Slice(arrS, func(i, j int) bool {
		return arrS[i]+arrS[j] > arrS[j]+arrS[i]
	})
	if arrS[0] == "0" {
		return "0"
	}
	return strings.Join(arrS, "")
}
