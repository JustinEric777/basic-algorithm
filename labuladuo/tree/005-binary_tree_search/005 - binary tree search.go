package main

import (
	"basic-algorithm/labuladuo/define/tree"
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 把二叉搜索树转换为累加树
// 每个节点的新值 等于 原树中大于或等于的 node.val 的值之和
func convertBinarySearchTreeToGreaterSumTree(root *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}

	traverseConvertGreaterSumTree(root)

	return root
}

// convert
var sum int

func traverseConvertGreaterSumTree(root *tree.Node) {
	if root == nil {
		return
	}

	traverseConvertGreaterSumTree(root.Right)
	sum += root.Val
	root.Val = sum
	traverseConvertGreaterSumTree(root.Left)
}

// 二叉搜索树中第k小的元素
// 二叉树存储每个子树的节点数量即可不用全遍历二叉树
var res []int

func kthSmallest(root *tree.Node, k int) int {
	if root == nil || k == 0 {
		return -1
	}

	// 中序有序
	traverseMid(root)

	return res[k-1]
}

func traverseMid(root *tree.Node) {
	if root == nil {
		return
	}

	traverseMid(root.Left)
	res = append(res, root.Val)
	traverseMid(root.Right)
}

// 删除二叉搜索树中的节点
func deleteNodeFromBST(root *tree.Node, target int) *tree.Node {
	if root == nil {
		return nil
	}

	return nil
}

// 二叉搜索树中的搜索
var result bool

func searchInBST(root *tree.Node, target int) bool {
	if root == nil {
		return false
	}

	// search
	traverseSearch(root, target)

	return result
}

func traverseSearch(root *tree.Node, target int) {
	if root == nil {
		return
	}

	if root.Val == target {
		result = true
	} else if root.Val < target {
		traverseSearch(root.Right, target)
	} else if root.Val > target {
		traverseSearch(root.Left, target)
	}
}

// 二叉搜索树中的插入操作
func insertToBST(root *tree.Node, target int) *tree.Node {

	return nil
}

// 验证二叉搜索树
func checkValidBST(root *tree.Node) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Val > root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	return checkValidBST(root.Left) && checkValidBST(root.Right)
}

// 不同的二叉搜索树

func main() {
	fmt.Println("1. 二叉搜索树转化为累加树：")
	node99 := &tree.Node{Left: nil, Right: nil, Val: 0}
	node88 := &tree.Node{Left: nil, Right: nil, Val: 3}
	node77 := &tree.Node{Left: nil, Right: nil, Val: 5}
	node66 := &tree.Node{Left: nil, Right: nil, Val: 8}
	node55 := &tree.Node{Left: nil, Right: node88, Val: 2}
	node44 := &tree.Node{Left: node99, Right: node55, Val: 1}
	node33 := &tree.Node{Left: nil, Right: node66, Val: 7}
	node22 := &tree.Node{Left: node77, Right: node33, Val: 6}
	node11 := &tree.Node{Left: node44, Right: node22, Val: 6}
	fmt.Println("原二叉树为：")
	tree.EchoTwoBranchTree(node11)
	sumTree := convertBinarySearchTreeToGreaterSumTree(node11)
	fmt.Println("转换后的二叉树为：")
	tree.EchoTwoBranchTree(sumTree)

	line.SplitLine()

	fmt.Println("2. 输出二叉搜索树中的第 k 小元素：")
	node6 := &tree.Node{Left: nil, Right: nil, Val: 1}
	node5 := &tree.Node{Left: nil, Right: nil, Val: 4}
	node4 := &tree.Node{Left: nil, Right: nil, Val: 6}
	node3 := &tree.Node{Left: node6, Right: nil, Val: 2}
	node2 := &tree.Node{Left: node3, Right: node5, Val: 3}
	node1 := &tree.Node{Left: node2, Right: node4, Val: 5}
	fmt.Println("原二叉搜索树为：")
	tree.EchoTwoBranchTree(node1)
	smallest := kthSmallest(node1, 4)
	fmt.Println("二叉搜索树中的第 k 小元素为：", smallest)

	line.SplitLine()

	fmt.Println("3. 删除二叉搜索树中的节点：")
	fmt.Println("原二叉树为：")
	tree.EchoTwoBranchTree(node1)
	delTarget := 6
	fmt.Println("需要删除的节点为：", delTarget)
	fmt.Println("删除改节点之后的二叉搜索树为：")

	line.SplitLine()

	fmt.Println("4. 二叉搜索树的查找：")
	fmt.Println("二叉搜索树为：")
	tree.EchoTwoBranchTree(node1)
	target := 7
	fmt.Println("需要查找的节点为：", target)
	searchRes := searchInBST(node1, target)
	fmt.Println("二叉搜索树的结果为：", searchRes)

	line.SplitLine()

	fmt.Println("5: 二叉搜索树中插入新的节点：")
	fmt.Println("原二叉树为：")
	tree.EchoTwoBranchTree(node1)
	inTarget := 8
	fmt.Println("二叉搜索树需要插入的节点为：", inTarget)
	fmt.Println("插入节点之后的二叉搜索树为：")

	line.SplitLine()

	fmt.Println("6. 验证二叉搜索树的合法性：")
	fmt.Println("原二叉树为：")
	tree.EchoTwoBranchTree(node1)
	validBST := checkValidBST(node1)
	fmt.Println("二叉搜索树的合法性判责结果为：", validBST)

	line.SplitLine()

	fmt.Println("7. 不同的二叉搜索树：")

}
