/*
 Binary Tree Inorder Traversal
 *  Left-->root-->Right
 */


 type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
 }

 package main

 func inOrderTraversal(root *TreeNode) []int {
	 var ans []int
	 traversalTree(root, &ans)
	 return ans
 }

 func traversalTree(root *TreeNode, ans *[]int) {
	 if root == nil {
		 return
	 }

	 traversalTree(root.Left, ans)
	 *ans = append(*ans, root.Val)
	 traversalTree(root.Right, ans)
 }