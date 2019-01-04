package aoc2018

import "strings"

func CalculateChecksum(words []string) int {
	count_2 := 0
	count_3 := 0

	for _, str := range words {
		seen := make(map[string]int)
		for _, s := range strings.Split(str, "") {
			seen[s]++
		}
		div2 := false
		div3 := false
		for _, v := range seen {
			if (v == 2) {
				div2 = true
			}
			if (v == 3) {
				div3 = true
			}
		}
		if (div2) {
			count_2++
		}
		if (div3) {
			count_3++
		}
	}
	return count_2 * count_3
}

func CalculateChecksum2(words []string) string {
	for _, str1 := range words {
		for _, str2 := range words {
			diff := 0
			diff_i := 0

			for i, c := range str1 {
				if (str2[i] != uint8(c)) {
					diff++
					diff_i = i
				}
			}

			if (diff == 1) {
				return str1[:diff_i]+str1[diff_i+1:]
			}
		}
	}
	return ""
}

