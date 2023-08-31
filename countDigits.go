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