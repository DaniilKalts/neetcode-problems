// 0001.Two Sum
// https://neetcode.io/problems/two-integer-sum

package easy

import (
	"sort"
)

// 1. Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func TwoSumSolution1(nums []int, target int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

// 2. Sorting
// Time Complexity: O(n * log(n))
// Space Complexity: O(1) or O(n) depending on the sorting algorithm
func TwoSumSolution2(nums []int, target int) []int {
	n := len(nums)
	items := make([][2]int, n)

	for i, num := range nums {
		items[i] = [2]int{i, num}
	}

	sort.Slice(
		items, func(i, j int) bool {
			return items[i][1] < items[j][1]
		},
	)

	l := 0
	r := n - 1
	for l < r {
		sum := items[l][1] + items[r][1]
		if sum == target {
			return []int{items[l][0], items[r][0]}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	return []int{-1, -1}
}

// 3. Hash Map
// Time Complexity: O(n)
// Space Complexity: O(n)
func TwoSumSolution3(nums []int, target int) []int {
	items := make(map[int]int)

	for i, num := range nums {
		items[num] = i
	}

	for i, num1 := range nums {
		diff := target - num1
		if j, ok := items[diff]; ok && i != j {
			return []int{i, j}
		}
	}

	return []int{-1, -1}
}

// 4. Hash Map (One Pass)
// Time Complexity: O(n)
// Space Complexity: O(n)
func TwoSumSolution4(nums []int, target int) []int {
	items := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if j, ok := items[diff]; ok && i != j {
			return []int{i, j}
		}
		items[num] = i
	}

	return []int{-1, -1}
}
