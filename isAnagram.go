// Leetcode Problem 242: Valid Anagram
// Time Complexity: O(n) where n is the length of the string
// Space Complexity: O(1)
package main

func isAnagram(s string, t string) bool {
	chars := make(map[rune]int)

	for _, v := range s {
		chars[v]++
	}

	for _, v := range t {
		chars[v]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}
