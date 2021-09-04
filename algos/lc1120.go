/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package algos

import "math"

func maximumAverageSubtree(root *TreeNode) float64 {
	_, _, result := maximumAverageSubtreeHelper(root)
	return result
}

var answer int

func maximumAverageSubtreeHelper(root *TreeNode) (nodeCount int, sum int, currMaxAvg float64) {
	if root == nil {
		return 0, 0, 0
	}
	nodeCountLeft, sumLeft, currMaxAvgLeft := maximumAverageSubtreeHelper(root.Left)
	nodeCountRight, sumRight, currMaxAvgRight := maximumAverageSubtreeHelper(root.Right)
	nodeCount = 1 + nodeCountLeft + nodeCountRight
	sum = sumLeft + sumRight + root.Val
	currMaxAvg = math.Max(math.Max(currMaxAvgLeft, currMaxAvgRight), float64(sum)/float64(nodeCount))
	return
}
