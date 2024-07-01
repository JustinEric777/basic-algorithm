package main

import (
	"basic-algorithm/labuladuo/define/queue"
	"basic-algorithm/labuladuo/define/tree"
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
	"math"
)

// 二叉树的最大深度
// 遍历二叉树
var (
	res   int
	depth int
)

func maxDepth(root *tree.Node) int {
	traverse(root)
	return res
}

func traverse(root *tree.Node) {
	if root == nil {
		return
	}
	depth++
	if root.Left == nil && root.Right == nil {
		res = int(math.Max(float64(res), float64(depth)))
	}

	traverse(root.Left)
	traverse(root.Right)
	depth--
}

// 二叉树的最大深度
// 分解问题
func maxDepth1(root *tree.Node) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth1(root.Left)
	RightMax := maxDepth1(root.Right)

	depth := int(math.Max(float64(leftMax), float64(RightMax)))

	return depth + 1
}

// 二叉树的前序遍历 - DFS
var resArr []int

func prevSort(root *tree.Node) []int {
	traversePrevOrder(root)

	return resArr
}

func traversePrevOrder(root *tree.Node) {
	if root == nil {
		return
	}
	resArr = append(resArr, root.Val)
	traversePrevOrder(root.Left)
	traversePrevOrder(root.Right)
}

// 二叉树的直径
var maxDiameter int

func maxWidth(root *tree.Node) int {
	maxDepth2(root)
	return maxDiameter
}

func maxDepth2(root *tree.Node) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth2(root.Left)
	rightMax := maxDepth2(root.Right)
	diameter := leftMax + rightMax
	maxDiameter = int(math.Max(float64(maxDiameter), float64(diameter)))

	return int(math.Max(float64(leftMax), float64(rightMax))) + 1
}

// 二叉树中的最大路径和
// 给定一个非空二叉树，返回其最大路径和
// 径被定义为一条从树中任意节点出发，达到任意节点的序列
var (
	resSum int
)

func maxRouteSum(root *tree.Node) int {
	if root == nil {
		return 0
	}
	traverseRouteSum(root)

	return resSum
}

// 根据 0 比较，避免溢出为负值
func traverseRouteSum(root *tree.Node) int {
	if root == nil {
		return 0
	}

	leftRouteSum := math.Max(float64(traverseRouteSum(root.Left)), 0)
	rightRouteSum := math.Max(float64(traverseRouteSum(root.Right)), 0)

	resSum = int(math.Max(leftRouteSum+rightRouteSum+float64(root.Val), float64(resSum)))

	return int(math.Max(leftRouteSum+float64(root.Val), rightRouteSum+float64(root.Val)))
}

// 层次遍历 - BFS
func levelTraverse(root *tree.Node) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	q := queue.NewQueue()
	q.Push(root)
	for q.Len() != 0 {
		qs := q.Len()
		for i := 0; i < qs; i++ {
			cur := q.Pop()
			curNode := cur.(*tree.Node)
			res = append(res, curNode.Val)

			if curNode.Left != nil {
				q.Push(curNode.Left)
			}
			if curNode.Right != nil {
				q.Push(curNode.Right)
			}
		}
	}

	return res
}

// 在每个树行中找到最大值 - 层次遍历直接判最大值即可
func findMaxLineValFromTree(root *tree.Node) {

}

func main() {
	root := &tree.Node{Left: &tree.Node{Left: nil, Val: 9, Right: nil}, Val: 3, Right: &tree.Node{Left: &tree.Node{Left: nil, Val: 15, Right: nil}, Val: 20, Right: &tree.Node{Left: nil, Val: 7, Right: nil}}}

	fmt.Println("1.查找二叉树的最大深度：（递归遍历）")
	fmt.Println("二叉树为：")
	tree.EchoTwoBranchTree(root)
	depth := maxDepth(root)
	fmt.Println("二叉树的最大深度为：", depth)

	fmt.Println("1.查找二叉树的最大深度：（分解遍历）")
	fmt.Println("二叉树为：")
	tree.EchoTwoBranchTree(root)
	depth1 := maxDepth1(root)
	fmt.Println("二叉树的最大深度为：", depth1)

	line.SplitLine()

	fmt.Println("2. 二叉树的前序遍历：(DFS)")
	fmt.Println("二叉树为：")
	tree.EchoTwoBranchTree(root)
	preSortedTree := prevSort(root)
	fmt.Println("前序遍历的结果为：", preSortedTree)

	line.SplitLine()

	fmt.Println("3. 二叉树的直径：")
	root1 := &tree.Node{Left: nil, Val: 9, Right: &tree.Node{Left: &tree.Node{Left: &tree.Node{Left: nil, Val: 4, Right: nil}, Val: 2, Right: &tree.Node{Left: nil, Val: 5, Right: nil}}, Val: 1, Right: &tree.Node{Left: nil, Val: 3, Right: nil}}}
	fmt.Println("二叉树为：")
	tree.EchoTwoBranchTree(root1)
	width := maxWidth(root)
	fmt.Println("二叉树的最大直径为：", width)

	line.SplitLine()

	fmt.Println("4. 二叉树的层次遍历：（BFS）")
	fmt.Println("二叉树为：")
	tree.EchoTwoBranchTree(root)
	resLevel := levelTraverse(root)
	fmt.Println("层序遍历的结果为：", resLevel)

	line.SplitLine()

	fmt.Println("5. 二叉树的最大路径和：")
	root2 := &tree.Node{Left: &tree.Node{Left: nil, Val: 6, Right: nil}, Val: 9, Right: &tree.Node{Left: &tree.Node{Left: nil, Val: -6, Right: nil}, Val: -3, Right: &tree.Node{Left: &tree.Node{Left: nil, Val: 2, Right: nil}, Val: 2, Right: nil}}}
	fmt.Println("二叉树为：")
	tree.EchoTwoBranchTree(root2)
	maxRouteSumVal := maxRouteSum(root2)
	fmt.Println("二叉树的最大路径和为：", maxRouteSumVal)
}
