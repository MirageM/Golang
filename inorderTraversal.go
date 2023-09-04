// Leetcode Problem 94: Binary Tree Inorder Traversal
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(n) where n is the number of nodes in the tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 101)
	if root != nil {
		res = help(root, res)
	}
	return res
}
func help(root *TreeNode, res []int) []int {
	if root.Left != nil {
		res = help(root.Left, res)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = help(root.Right, res)
	}
	return res
}
