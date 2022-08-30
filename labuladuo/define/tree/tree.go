package tree

import (
	"fmt"
	"math"
)

type Node struct {
	Val   int   `json:"val"`
	Left  *Node `json:"left"`
	Right *Node `json:"right"`
}

func EchoTwoBranchTree(root *Node) {
	if root == nil {
		return
	}

	// 获取二叉树的宽度和深度
	height := getDepth(root, 0)

	// 数组的初始化
	arrayHeight := height*2 - 1
	arrayWidth := (2<<(height-2))*3 + 1
	resArray := make([][]string, arrayHeight)
	for i := 0; i < arrayHeight; i++ {
		lineArr := make([]string, arrayWidth)
		for j := 0; j < arrayWidth; j++ {
			lineArr[j] = " "
		}
		resArray[i] = lineArr
	}

	// 填充数组
	withArray(root, 0, arrayWidth/2, resArray, height)

	for _, line := range resArray {
		for i := 0; i < len(line); i++ {
			fmt.Print(line[i])
		}
		fmt.Println("")
	}
}

// 数组元素填充到数组
func withArray(root *Node, rowIndex, columnIndex int, res [][]string, depth int) {
	if root == nil {
		return
	}

	// 当前节点保存到数组中
	stringVal := fmt.Sprintf("%d", root.Val)
	res[rowIndex][columnIndex] = stringVal

	// 计算当前位于树的第几层
	curLevel := (rowIndex + 1) / 2
	if curLevel == depth {
		return
	}

	// 计算行元素之间的间隔
	gap := depth - curLevel - 1

	// 左子树
	if root.Left != nil {
		res[rowIndex+1][columnIndex-gap] = "/"
		withArray(root.Left, rowIndex+2, columnIndex-2*gap, res, depth)
	}

	// 右子树
	if root.Right != nil {
		res[rowIndex+1][columnIndex+gap] = "\\"
		withArray(root.Right, rowIndex+2, columnIndex+2*gap, res, depth)
	}
}

// 获取宽度
func getWidth(root *Node, maxWidth int) int {
	if root != nil {
		return 0
	}

	left := getWidth(root.Left, maxWidth)
	right := getWidth(root.Right, maxWidth)
	maxWidth = int(math.Max(float64(left+right), float64(maxWidth)))

	return int(math.Max(float64(left), float64(right))) + 1
}

// 获取深度
func getDepth(root *Node, maxDepth int) int {
	if root == nil {
		return 0
	}

	left := getDepth(root.Left, maxDepth)
	right := getDepth(root.Right, maxDepth)

	maxDepth = int(math.Max(float64(left), float64(right))) + 1

	return maxDepth
}
