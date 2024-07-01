package main

import (
	"basic-algorithm/labuladuo/define/tree"
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
	"strconv"
	"strings"
)

type TreeSerialization struct {
	StrTree []string
}

// 二叉树的序列化
func (treeSerialization *TreeSerialization) Serialize(root *tree.Node) string {
	if root == nil {
		return ""
	}

	treeSerialization.StrTree = append(treeSerialization.StrTree, fmt.Sprintf("%d", root.Val))
	if root.Left == nil {
		treeSerialization.StrTree = append(treeSerialization.StrTree, "*")
	}
	if root.Right == nil {
		treeSerialization.StrTree = append(treeSerialization.StrTree, "*")
	}

	treeSerialization.Serialize(root.Left)
	treeSerialization.Serialize(root.Right)

	return strings.Join(treeSerialization.StrTree, "")
}

// 二叉树的反序列化
func (treeSerialization *TreeSerialization) Deserialize() *tree.Node {
	if len(treeSerialization.StrTree) == 0 {
		return nil
	}

	if treeSerialization.StrTree[0] == "*" {
		treeSerialization.StrTree = treeSerialization.StrTree[1:]
		return nil
	}

	val, _ := strconv.Atoi(treeSerialization.StrTree[0])
	node := &tree.Node{Left: nil, Right: nil, Val: val}
	treeSerialization.StrTree = treeSerialization.StrTree[1:]
	node.Left = treeSerialization.Deserialize()
	node.Right = treeSerialization.Deserialize()

	return node
}

var nums []string

func (treeSerialization *TreeSerialization) Deserialize2(str string) *tree.Node {
	if len(str) == 0 {
		return nil
	}

	nums = strings.Split(str, "")
	node := traverse()

	return node
}

func traverse() *tree.Node {
	if len(nums) == 0 {
		return nil
	}
	if nums[0] == "*" {
		nums = nums[1:]
		return nil
	}

	val, _ := strconv.Atoi(nums[0])
	node := &tree.Node{Left: nil, Right: nil, Val: val}
	nums = nums[1:]
	node.Left = traverse()
	node.Right = traverse()

	return node
}

// 寻找重复的子树
// dfs过程中实现序列化,转存hash map计数
var hashMap map[string]int

func findRepeatSubTree(root *tree.Node) [][]string {
	if root == nil {
		return [][]string{}
	}

	hashMap = make(map[string]int)
	traverseSubTree(root)

	var res [][]string
	// 整理hashMap
	for key, val := range hashMap {
		if val > 1 {
			subtree := strings.Replace(key, "*", "", -1)
			res = append(res, strings.Split(subtree, ""))
		}
	}

	return res
}

func traverseSubTree(root *tree.Node) {
	if root == nil {
		return
	}

	// 序列化树
	serialization := &TreeSerialization{}
	serialize := serialization.Serialize(root)
	hashMap[serialize]++

	traverseSubTree(root.Left)
	traverseSubTree(root.Right)
}

func main() {
	fmt.Println("1. 二叉树的序列化和序列化：")
	node7 := &tree.Node{Left: nil, Right: nil, Val: 1}
	node6 := &tree.Node{Left: nil, Right: nil, Val: 3}
	node5 := &tree.Node{Left: nil, Right: nil, Val: 6}
	node4 := &tree.Node{Left: nil, Right: nil, Val: 9}
	node2 := &tree.Node{Left: node7, Right: node6, Val: 2}
	node1 := &tree.Node{Left: node5, Right: node4, Val: 7}
	root := &tree.Node{Left: node2, Right: node1, Val: 4}
	fmt.Println("二叉树为：")
	tree.EchoTwoBranchTree(root)
	t := &TreeSerialization{}
	serialize := t.Serialize(root)
	fmt.Println("二叉树序列化的结果为：", serialize)
	fmt.Println("二叉树的反序列化结果为：")
	node := t.Deserialize()
	tree.EchoTwoBranchTree(node)

	line.SplitLine()

	fmt.Println("二叉树的反序列化结果为：")
	deserialize2 := t.Deserialize2(serialize)
	tree.EchoTwoBranchTree(deserialize2)

	line.SplitLine()

	node77 := &tree.Node{Left: nil, Right: nil, Val: 4}
	node66 := &tree.Node{Left: nil, Right: nil, Val: 4}
	node55 := &tree.Node{Left: nil, Right: nil, Val: 4}
	node44 := &tree.Node{Left: node77, Right: nil, Val: 2}
	node33 := &tree.Node{Left: node66, Right: nil, Val: 2}
	node22 := &tree.Node{Left: node33, Right: node55, Val: 3}
	node11 := &tree.Node{Left: node44, Right: node22, Val: 1}
	fmt.Println("2. 寻找重复的子树：")
	fmt.Println("原二叉树为：")
	tree.EchoTwoBranchTree(node11)
	subTree := findRepeatSubTree(node11)
	fmt.Println("寻找的重复子树为：", subTree)
}
