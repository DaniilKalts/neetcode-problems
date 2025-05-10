// 0049. Group Anagrams
// https://leetcode.com/problems/group-anagrams/description/

package easy

import "sort"

// 1. Sorting
// Time Complexity: O(m * n * log(n))
// Space Complexity: O(m * n)
func solution_1(strs []string) [][]string {
	hashMap := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		hashMap[sortedStr] = append(hashMap[sortedStr], str)
	}

	var res [][]string
	for _, group := range hashMap {
		res = append(res, group)
	}

	return res
}

func sortString(str string) string {
	chars := []rune(str)

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}

// 2. Hash Table
// Time Complexity: O(n * m)
// Space Complexity: O(n * m)
func solution_2(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)
	for _, str := range strs {
		var counter [26]int
		for _, char := range str {
			counter[char-'a']++
		}
		hashMap[counter] = append(hashMap[counter], str)
	}

	var res [][]string
	for _, group := range hashMap {
		res = append(res, group)
	}

	return res
}
