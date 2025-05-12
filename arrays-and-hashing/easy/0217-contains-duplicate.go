// 0217. Contains Duplicate
// https://neetcode.io/problems/duplicate-integer

package easy

import "sort"

// 1. Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func solution1(nums []int) bool {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

// 2. Sorting
// Time Complexity: O(n * log(n))
// Space Complexity: O(1) or O(n) depending on the sorting algorithm
func solution2(nums []int) bool {
	sort.Ints(nums)
	n := len(nums)

	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

// 3. Hash Set
// Time Complexity: O(n)
// Space Complexity: O(n)
func solution3(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

// 4. Hash Set Length
// Time Complexity: O(n)
// Space Complexity: O(n)
func solution4(nums []int) bool {
	seen := make(map[int]struct{})

	for _, num := range nums {
		seen[num] = struct{}{}
	}

	return len(seen) != len(nums)
}
