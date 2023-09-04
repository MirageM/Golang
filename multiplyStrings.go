// Leetcode Problem 43: Multiply Strings
// Time Complexity: O(n*m) where n is the length of num1 and m is the length of nums2
// Space Complexity: O(n+m)

package main

import "strconv"

func multiply(num1 string, num2 string) string {
	n1 := []byte(num1)
	n2 := []byte(nums)
	l1 := len(num1)
	l2 := len(num2)
	result := make([]int, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - l2 - 1; j >= 0; j-- {
			n1i, _ := strconv.Atoi(string(n1[i]))
			n2i, _ := strconv.Atoi(string(n2[j]))
			mul := n1i * n2i
			p1 := i + j
			p2 := p1 + 1
			sum := res[p2] + mul
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}
	i := 0
	for i < len(result) && result[i] == 0 {
		i++
	}
	str := ""
	for i < len(result) {
		str += strconv.Itoa(result[i])
		i++
	}
	if len(str) == 0 {
		return "0"
	}
	return str

}
