// 0242. Valid Anagram
// https://neetcode.io/problems/is-anagram

package easy

import "sort"

// 1. Sorting
// Time Complexity: O(n * log(n) + m * log(m)) => O(n * log(n))
// Space Complexity: O(n + m) => O(n)
func ValidAnagramSolution1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRunes := []rune(s)
	sort.Slice(
		sRunes, func(i, j int) bool {
			return sRunes[i] < sRunes[j]
		},
	)

	tRunes := []rune(t)
	sort.Slice(
		tRunes, func(i, j int) bool {
			return tRunes[i] < tRunes[j]
		},
	)

	for i := range sRunes {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}

	return true
}

// 2. Hash Map
// Time Complexity: O(n + m) => O(n)
// Space Complexity: O(n)
func ValidAnagramSolution2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sCount := make(map[rune]int)
	tCount := make(map[rune]int)

	for i := range s {
		sCount[rune(s[i])]++
		tCount[rune(t[i])]++
	}

	for r := range sCount {
		if sCount[r] != tCount[r] {
			return false
		}
	}

	return true
}

// 3. Hash Table (Using Array)
// Time Complexity: O(n + 26) => O(n)
// Space Complexity: O(1) since we have at most 26 different characters
func ValidAnagramSolution3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := [26]int{}

	for i := range s {
		counter[rune(s[i])-'a']++
		counter[rune(t[i])-'a']--
	}

	for i := range counter {
		if counter[i] < 0 {
			return false
		}
	}

	return true
}
