// Leetcode Problem 2520: Count the Digits That Divide A Number
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