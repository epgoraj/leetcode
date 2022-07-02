/*
 Binary Tree Level Order Traversal
*/

package binarytre

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderTraversal(root *TreeNode) [][]int {
	var ans [][]int
	traversalTree(root, 0, &ans)
	return ans
}

func traversalTree(root *TreeNode, level int, ans *[][]int) {
	if root == nil {
		return
	}
	curLevel := level
	if len(*ans) <= curLevel {
		*ans = append(*ans, []int{})
	}
	(*ans)[curLevel] = append((*ans)[curLevel], root.Val)
	traversalTree(root.Left, level+1, ans)
	traversalTree(root.Right, level+1, ans)

}
