// Leetcode Problem 389: Find the Difference
// Time Complexity: O(n) where n is the length of s
// Space Complexity: O(n) where n is the length of s

package main

func findTheDifference(s string, t string) byte {
	ms, mt := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		ms[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}
	for k, v := range mt {
		if c, ok := ms[k]; !ok || c != v {
			return k
		}
	}
	var b byte
	return b
}
