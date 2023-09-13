// Leetcode Problem 389: Find the Difference
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
