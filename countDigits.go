// Leetcode Problem 2520: Count the Digits That Divide A Number
// Time Comlexity: O(n) where n is the number of digits in the number
// Space Complexity: O(1)
package main

func countDigits(num int) int {
	temp := num
	count := 0
	for temp > 0 {
		digit := temp % 10
		temp = temp / 10
		if num%digit == 0 {
			count++
		}
	}
	return count
}
