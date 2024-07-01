package main

import (
	"basic-algorithm/labuladuo/define/tree"
	"fmt"
	"math"
)

// 完全二叉树的节点个数
// 完全二叉树为 叶子节点相差的层数不大于1，满二叉树为特殊的完全二叉树
func numsOfCompleteTree(root *tree.Node) int {
	if root == nil {
		return 0
	}

	var lh, rh int
	left, right := root, root
	for left != nil {
		left = left.Left
		lh++
	}
	for right != nil {
		right = right.Right
		rh++
	}

	if lh == rh {
		return int(math.Pow(float64(2), float64(rh))) - 1
	}

	return 1 + numsOfCompleteTree(root.Left) + numsOfCompleteTree(root.Right)
}

func main() {
	fmt.Println("1. 计算完全二叉树的节点个数：")
	node5 := &tree.Node{Left: nil, Right: nil, Val: 4}
	node4 := &tree.Node{Left: nil, Right: nil, Val: 5}
	node3 := &tree.Node{Left: nil, Right: nil, Val: 6}
	node2 := &tree.Node{Left: node5, Right: node4, Val: 2}
	node1 := &tree.Node{Left: node3, Right: nil, Val: 3}
	node := &tree.Node{Left: node2, Right: node1, Val: 1}
	fmt.Println("原完全二叉树为：")
	tree.EchoTwoBranchTree(node)
	completeTreeNodeNums := numsOfCompleteTree(node)
	fmt.Println("完全二叉树的节点个数为：", completeTreeNodeNums)

}
