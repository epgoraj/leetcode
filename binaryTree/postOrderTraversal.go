/*
 Binary Tree Postorder Traversal
 *  Bottom --> Top
 *  Left  --> Right
*/

package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrderTraversal(root *TreeNode) []int {
	var ans []int
	traversalTree(root, &ans)
	return ans
}

func traversalTree(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	traversalTree(root.Left, ans)
	traversalTree(root.Right, ans)
	*ans = append(*ans, root.Val)
}
