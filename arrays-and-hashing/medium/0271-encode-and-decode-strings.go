// 0271. Encode and Decode Strings
// https://neetcode.io/problems/string-encode-and-decode

package medium

import (
	"strconv"
	"strings"
)

// 1. Encoding and Decoding
// Time Complexity: O(n)
// Space Complexity: O(n)
func EncodeStringsSolution1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var sizes []string
	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}

	return strings.Join(sizes, ",") + "#" + strings.Join(strs, "")
}

// 1. Encoding and Decoding
// Time Complexity: O(n)
// Space Complexity: O(n)
func DecodeStringsSolution1(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	var res []string

	parts := strings.SplitN(str, "#", 2)
	sizes := strings.Split(parts[0], ",")

	l := 0
	for _, sz := range sizes {
		length, _ := strconv.Atoi(sz)
		res = append(res, parts[1][l:l+length])
		l += length
	}

	return res
}

// 2. Encoding and Decoding (Optimal)
// Time Complexity: O(n^2) because in Go += creates a new string
// Space Complexity: O(n)
func EncodeStringsSolution2(strs []string) string {
	var res string

	for _, str := range strs {
		res += strconv.Itoa(len(str)) + "#" + str
	}

	return res
}

// 2. Encoding and Decoding (Optimal)
// Time Complexity: O(n)
// Space Complexity: O(n)
func DecodeStringsSolution2(str string) []string {
	var res []string

	l := 0
	for l < len(str) {
		r := l
		for str[r] != '#' {
			r++
		}
		length, _ := strconv.Atoi(str[l:r])
		l = r + 1
		res = append(res, str[l:l+length])
		l += length
	}

	return res
}
