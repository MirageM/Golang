// Leetcode Problem 496. Next Greater Element I
// Time Complexity: O(n^2) where n is the length of nums2
// Space Complexity: O(n)

package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums2))
	for i, n := range nums2 {
		for j := i + 1; j < len(nums2); j++ {
			if n < nums2[j] {
				m[n] = nums2[j]
				break
			}
		}
		if m[n] == 0 {
			m[n] = -1
		}
	}
	ret := make([]int, len(nums1))
	for i, n := range nums1 {
		if x, ok := m[n]; ok {
			ret[i] = x
		}
	}
	return ret
}
