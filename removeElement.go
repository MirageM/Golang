// Leetcode Problem 27: Remove Element
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1)
package main

func removeElement(nums []int, val int) int {
	index := 0
	for _, num := range nums {
		if num != val {
			nums[index] = num
			index += 1
		}
	}
	return index
}
