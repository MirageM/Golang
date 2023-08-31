//Leetcode Problem 1: Two Sum
// Time Complexity: O(n^2) where n is the length of the array
// Space Complexity: O(1)
package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
