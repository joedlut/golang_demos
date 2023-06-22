package main

//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
/*

Definition for a binary tree node.

*/

// 递归算法，不推荐
func inorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	if root == nil {
		return nums
	}
	nums = append(nums, inorderTraversal(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, inorderTraversal(root.Right)...)
	return nums
}
