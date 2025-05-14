// 0238. Product of Array Except Self
// https://neetcode.io/problems/products-of-array-discluding-self

package medium

// 1. Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1) extra space and O(n) for the output array
func ProductOfArrayExceptSelfSolution1(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	for i := range n {
		product := 1
		for j := range n {
			if i == j {
				continue
			}
			product *= nums[j]
		}
		res[i] = product
	}

	return res
}

// 2. Division
// Time Complexity: O(n)
// Space Complexity: O(1) extra space and O(n) for the output array
func ProductOfArrayExceptSelfSolution2(nums []int) []int {
	zeroCount := 0
	product := 1

	for _, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			product *= num
		}
	}

	n := len(nums)
	res := make([]int, n)

	if zeroCount > 1 {
		return res
	}

	for i := range n {
		if zeroCount > 0 {
			if nums[i] == 0 {
				res[i] = product
			} else {
				res[i] = 0
			}
		} else {
			res[i] = product / nums[i]
		}
	}

	return res
}

// 3. Prefix & Sufix
// Time Complexity: O(n)
// Space Complexity: O(n)
func ProductOfArrayExceptSelfSolution3(nums []int) []int {
	n := len(nums)

	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	postfix := make([]int, n)
	postfix[n-1] = 1
	for i := n - 2; i > -1; i-- {
		postfix[i] = postfix[i+1] * nums[i+1]
	}

	res := make([]int, n)
	for i := range n {
		res[i] = prefix[i] * postfix[i]
	}

	return res
}

// 4. Prefix & Sufix (Optimal)
// Time Complexity: O(n)
// Space Complexity: O(1) extra space and O(n) for the output array
func ProductOfArrayExceptSelfSolution4(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	prefix := 1
	for i := range nums {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := n - 1; i > -1; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
