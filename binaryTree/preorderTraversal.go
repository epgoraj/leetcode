/*
 Binary Tree Preorder Traversal
 *  Top --> Bottom
 *  Left --> Right
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var root *TreeNode
	result := preOrderTraversal(root)
	fmt.Printf("Result:%v", result)
}

func preOrderTraversal(root *TreeNode) []int {
	var ans []int
	traversalTree(root, &ans)
	return ans
}

func traversalTree(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	*ans = append(*ans, root.Val)
	traversalTree(root.Left, ans)
	traversalTree(root.Right, ans)

}
