// 0128. Longest Consecutive Sequence
// https://neetcode.io/problems/longest-consecutive-sequence

package medium

import "sort"

// 1. Brute Force (with optimization trick)
// Time Complexity: O(n)
// Space Complexity: O(n)
func LongestConsecutiveSequenceSolution1(nums []int) int {
	hashSet := make(map[int]struct{})
	for _, num := range nums {
		hashSet[num] = struct{}{}
	}

	longest := 0
	for _, num := range nums {
		if _, ok := hashSet[num-1]; ok {
			continue
		}
		streak, curr := 0, num
		for _, ok := hashSet[curr]; ok; _, ok = hashSet[curr] {
			streak++
			curr++
		}
		if streak > longest {
			longest = streak
		}
	}

	return longest
}

// 2. Sorting
// Time Complexity: O(n * log(n))
// Space Complexity: O(1) or O(n) depending on the sorting algorithm
func LongestConsecutiveSequenceSolution2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sort.Ints(nums)

	longest, streak := 1, 1
	for i := 1; i < n; i++ {
		switch {
		case nums[i-1] == nums[i]:
			continue
		case nums[i-1]+1 == nums[i]:
			streak++
		default:
			if streak > longest {
				longest = streak
			}
			streak = 1
		}
	}

	if streak > longest {
		longest = streak
	}

	return longest
}

// 3. Hash Set
// Time Complexity: O(n)
// Space Complexity: O(n)
func LongestConsecutiveSequenceSolution3(nums []int) int {
	hashSet := make(map[int]struct{})
	for _, num := range nums {
		hashSet[num] = struct{}{}
	}

	longest := 0
	for num := range hashSet {
		if _, exists := hashSet[num-1]; !exists {
			length := 1
			for {
				if _, found := hashSet[num+length]; found {
					length++
				} else {
					break
				}
			}
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

// 4. Hash Map
// Time Complexity: O(n)
// Space Complexity: O(n)
func LongestConsecutiveSequenceSolution4(nums []int) int {
	hashMap := make(map[int]int)
	longest := 0

	for _, num := range nums {
		if hashMap[num] == 0 {
			left := hashMap[num-1]
			right := hashMap[num+1]
			sum := left + right + 1

			hashMap[num] = sum
			hashMap[num-left] = sum
			hashMap[num+right] = sum

			if sum > longest {
				longest = sum
			}
		}
	}

	return longest
}
