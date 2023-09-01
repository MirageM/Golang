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
