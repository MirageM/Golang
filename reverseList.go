// Leetcode Problem 206: Reverse Linked List
// Time Complexity: O(n) where n is the length of the linked list
// Space Complexity: O(1)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var revHead *ListNode
	for head != nil {
		temp := head.Next
		head.Next = revHead
		revHead = head
		head = temp
	}

	return revHead
}
