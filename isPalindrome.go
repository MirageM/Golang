// Leetcode Problem 9: Palindrome Number
// Time Complexity: O(n) where n is the number of digits in the number
// Space Complexity: O(1)

package main

import "strconv"

func isPalindrome(x int) bool {
	string_x := strconv.Itoa(x)
	n := len(string_x)
	for i := 0; i < n/2; i++ {
		if rune(string_x[i]) != rune(string_x[n-1-i]) {
			return false
		}
	}
	return true
}
