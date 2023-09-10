// Leetcode Problem 303: Range Sum Query - Immutable

package main

type NumArray struct {
	size        int
	prefixTable []int
}

func Constructor(nums []int) NumArray {
	numArray := new(NumArray)
	numArray.prefixTable = make([]int, len(nums))
	numArray.prefixTable[0] = nums[0]
	for k := 1; k < len(nums); k++ {
		numArray.prefixTable[k] = numArray.prefixTable[k-1] + nums[k]
	}
	return *numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixTable[right]
	} else {
		return this.prefixTable[right] - this.prefixTable[left-1]
	}

}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
