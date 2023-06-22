package main

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
*/

//从上往下递归
//左右子树都是平衡二叉树
//左右子树的高度差不超过1

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && abs(Height(root.Right)-Height(root.Left)) <= 1

}

//计算树的高度

func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(Height(root.Left), Height(root.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

}
