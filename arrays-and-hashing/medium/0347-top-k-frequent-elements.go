// 0347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/description/

package easy

import "sort"

// 1. Sorting
// Time Complexity: O(n * log(n))
// Space Complexity: O(n)
func solution_1(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	var pairs [][2]int
	for num, count := range counter {
		pairs = append(pairs, [2]int{num, count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	var res []int
	for i := range pairs {
		res = append(res, pairs[i][0])
		if len(res) == k {
			break
		}
	}

	return res
}

// 2. Min-Heap
// Time Complexity: O(n * log(k))
// Space Complexity: O(n + k)
func solution_2(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	heap := priorityqueue.NewWith(func(a, b interface{}) int {
		countA := a.([2]int)[1]
		countB := b.([2]int)[1]
		return utils.IntComparator(countA, countB)
	})

	for num, count := range counter {
		heap.Enqueue([2]int{num, count})
		if heap.Size() > k {
			heap.Dequeue()
		}
	}

	res := make([]int, k)
	for i := k - 1; i > -1; i-- {
		num, _ := heap.Dequeue()
		res[i] = num.([2]int)[0]
	}

	return res
}

// 3. Bucket Sort
// Time Complexity: O(n)
// Space Complexity: O(n)
func solution_3(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	groups := make([][]int, len(nums)+1)
	for num, count := range counter {
		groups[count] = append(groups[count], num)
	}

	res := make([]int, 0, k)
	for i := len(nums); i > -1 && len(res) < k; i-- {
		for _, num := range groups[i] {
			res = append(res, num)
		}
	}

	return res
}
