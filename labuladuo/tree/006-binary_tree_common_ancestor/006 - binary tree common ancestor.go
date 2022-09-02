package main

import (
	"basic-algorithm/labuladuo/define/tree"
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 两个指定节点的二叉树的最近公共祖先
func binaryTreeCommonAncestor(root *tree.Node, node1, node2 *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}
	if root.Val == node1.Val || root.Val == node2.Val {
		return root
	}

	left := binaryTreeCommonAncestor(root.Left, node1, node2)
	right := binaryTreeCommonAncestor(root.Right, node1, node2)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}

	return left
}

// 二叉搜索数的最近公共祖先
func BSTCommonAncestor(root *tree.Node, node1, node2 *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}

	var ancestor *tree.Node
	if root.Val == node1.Val || root.Val == node2.Val {
		return root
	}
	if node1.Val > node2.Val && root.Val < node2.Val ||
		node1.Val < node2.Val && root.Val < node1.Val {
		ancestor = BSTCommonAncestor(root.Right, node1, node2)
	}
	if node1.Val > node2.Val && root.Val > node1.Val ||
		node1.Val < node2.Val && root.Val > node2.Val {
		ancestor = BSTCommonAncestor(root.Left, node1, node2)
	}

	if ancestor != nil {
		return ancestor
	}

	return root
}

func main() {
	fmt.Println("1. 二叉树的最近公共祖先：")
	node8 := &tree.Node{Left: nil, Right: nil, Val: 6}
	node7 := &tree.Node{Left: nil, Right: nil, Val: 7}
	node6 := &tree.Node{Left: nil, Right: nil, Val: 4}
	node5 := &tree.Node{Left: nil, Right: nil, Val: 0}
	node4 := &tree.Node{Left: nil, Right: nil, Val: 8}
	node3 := &tree.Node{Left: node7, Right: node6, Val: 2}
	node2 := &tree.Node{Left: node8, Right: node3, Val: 5}
	node1 := &tree.Node{Left: node5, Right: node4, Val: 1}
	node := &tree.Node{Left: node2, Right: node1, Val: 3}
	fmt.Println("原二叉树为：")
	tree.EchoTwoBranchTree(node)
	ancestor := binaryTreeCommonAncestor(node, node2, node1)
	fmt.Println("二叉树的最近公共祖先为：", ancestor)

	line.SplitLine()

	fmt.Println("2. 二叉搜索树的最近公共祖先：")
	node88 := &tree.Node{Left: nil, Right: nil, Val: 0}
	node77 := &tree.Node{Left: nil, Right: nil, Val: 3}
	node66 := &tree.Node{Left: nil, Right: nil, Val: 5}
	node55 := &tree.Node{Left: nil, Right: nil, Val: 7}
	node44 := &tree.Node{Left: nil, Right: nil, Val: 9}
	node33 := &tree.Node{Left: node77, Right: node66, Val: 4}
	node22 := &tree.Node{Left: node88, Right: node33, Val: 2}
	node11 := &tree.Node{Left: node55, Right: node44, Val: 8}
	node0 := &tree.Node{Left: node22, Right: node11, Val: 6}
	fmt.Println("原二叉搜索树为：")
	tree.EchoTwoBranchTree(node0)
	fmt.Println("需要查找的两个节点为：", node22, node11)
	commonAncestor := BSTCommonAncestor(node0, node11, node22)
	fmt.Println("二叉搜索树的公共祖先为：", commonAncestor)

}
