package algos

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//543. Diameter of Binary Trees
func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := helper(root)
	return diameter

}

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftHeight, diamLeft := helper(root.Left)
	rightHeight, diamRight := helper(root.Right)
	return 1 + max(leftHeight, rightHeight), max(max(diamLeft, diamRight), leftHeight+rightHeight)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
