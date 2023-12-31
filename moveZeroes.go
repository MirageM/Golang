// Leetcode Problem 283. Move Zeroes
// Time Complexity: O(n) where n is the length of nums
// Space Complexity: O(1)
package main

func moveZeroes(nums []int) {
	l := 0

	for r := range nums {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
	}
	for ; l < len(nums); l++ {
		nums[l] = 0
	}
	return
}
