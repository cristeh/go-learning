/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//297. Serialize and Deserialize Binary Tree
package algos

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	sb strings.Builder
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			this.sb.WriteString("nil ")
			return
		}
		this.sb.WriteString(fmt.Sprintf("%d ", node.Val))
		helper(node.Left)
		helper(node.Right)
	}
	helper(root)
	fmt.Println(this.sb.String())
	return this.sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	serializedNodes := strings.Split(data, " ")
	var helper func() *TreeNode
	helper = func() *TreeNode {
		if len(serializedNodes) == 0 {
			return nil
		}
		if serializedNodes[0] == "nil" {
			serializedNodes = serializedNodes[1:]
			return nil
		}
		val, _ := strconv.Atoi(serializedNodes[0])
		node := &TreeNode{
			Val: val,
		}
		serializedNodes = serializedNodes[1:]
		node.Left = helper()
		node.Right = helper()
		return node
	}
	root := helper()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

//1
