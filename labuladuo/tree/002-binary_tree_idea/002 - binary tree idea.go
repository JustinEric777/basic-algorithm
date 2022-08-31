package main

import (
	"basic-algorithm/labuladuo/define/tree"
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 反转二叉树 / 二叉树的镜像
func reserveTwoBranchTree(root *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}

	left := reserveTwoBranchTree(root.Left)
	right := reserveTwoBranchTree(root.Right)

	root.Right = left
	root.Left = right

	return root
}

// 填充每个二叉树节点的右边指针
// 完全二叉树，让当前指针 next 指向下一个右侧节点，无右侧节点则为 null
type node struct {
	val   int
	left  *node
	right *node
	next  *node
}

// 本质上一个三叉树的遍历
func addRightNodeFromTwoBranchTree(root *node) *node {
	if root == nil {
		return nil
	}

	threeTraverse(root.left, root.right)

	return root
}

func threeTraverse(node1, node2 *node) {
	if node1 == nil || node2 == nil {
		return
	}

	node1.next = node2
	threeTraverse(node1.left, node1.right)
	threeTraverse(node2.left, node2.right)

	threeTraverse(node1.right, node2.left)
}

// 将二叉树展开为链表
// 其中二叉树的 right 指针指向链表中的下一个节点，左指针始终为 null
// 展开结果与前序遍历结果相同
func convertTwoBranchTreeToLink(root *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}

	convertTwoBranchTreeToLink(root.Left)
	convertTwoBranchTreeToLink(root.Right)

	left := root.Left
	right := root.Right

	// 左子树为null
	root.Left = nil
	root.Right = left

	// 原先的右子树下移现右子树下侧
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

	return root
}

func main() {
	node7 := &tree.Node{Left: nil, Right: nil, Val: 1}
	node6 := &tree.Node{Left: nil, Right: nil, Val: 3}
	node5 := &tree.Node{Left: nil, Right: nil, Val: 6}
	node4 := &tree.Node{Left: nil, Right: nil, Val: 9}
	node2 := &tree.Node{Left: node7, Right: node6, Val: 2}
	node1 := &tree.Node{Left: node5, Right: node4, Val: 7}
	root := &tree.Node{Left: node2, Right: node1, Val: 4}
	fmt.Println("1. 二叉树的反转：(递归)")
	fmt.Println("原二叉树为：")
	tree.EchoTwoBranchTree(root)
	fmt.Println("反转后的二叉树为：")
	res := reserveTwoBranchTree(root)
	tree.EchoTwoBranchTree(res)

	line.SplitLine()

	node77 := &node{left: nil, right: nil, val: 1, next: nil}
	node66 := &node{left: nil, right: nil, val: 3, next: nil}
	node55 := &node{left: nil, right: nil, val: 6, next: nil}
	node44 := &node{left: nil, right: nil, val: 9, next: nil}
	node22 := &node{left: node77, right: node66, val: 2, next: nil}
	node11 := &node{left: node55, right: node44, val: 7, next: nil}
	root2 := &node{left: node22, right: node11, val: 4, next: nil}
	fmt.Println("2. 填充二叉树右边节点的next指针：")
	fmt.Println("原二叉树为：", root2)
	twoBranchTree := addRightNodeFromTwoBranchTree(root2)
	fmt.Println("链接右指针后的二叉树为：", twoBranchTree)

	line.SplitLine()

	node777 := &tree.Node{Left: nil, Right: nil, Val: 3}
	node666 := &tree.Node{Left: nil, Right: nil, Val: 4}
	node555 := &tree.Node{Left: nil, Right: nil, Val: 6}
	node333 := &tree.Node{Left: node777, Right: node666, Val: 2}
	node222 := &tree.Node{Left: nil, Right: node555, Val: 5}
	root3 := &tree.Node{Left: node333, Right: node222, Val: 1}
	fmt.Println("3. 二叉树展开为链表：")
	fmt.Println("原二叉树为：")
	tree.EchoTwoBranchTree(root3)
	linkTree := convertTwoBranchTreeToLink(root3)
	fmt.Println("转化为链表的结果为：")
	tree.EchoTwoBranchTree(linkTree)

}
