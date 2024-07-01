package main

import (
	"basic-algorithm/labuladuo/define/tree"
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 前序和中序构建二叉树
func buildTwoBranchTreeFromPrevSortAndMidSort(nums1, nums2 []int) *tree.Node {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	// 确认根节点的值
	nodeVal := nums1[0]
	node := &tree.Node{Left: nil, Right: nil, Val: nodeVal}

	// 确认根节点再中序遍历中的位置，确认左右子树
	position := 0
	for i := 0; i < len(nums2); i++ {
		if nodeVal == nums2[i] {
			position = i
		}
	}

	node.Left = buildTwoBranchTreeFromPrevSortAndMidSort(nums1[1:position+1], nums2[:position])
	node.Right = buildTwoBranchTreeFromPrevSortAndMidSort(nums1[position+1:], nums2[position+1:])

	return node
}

// 中序和后续构造二叉树
func buildTwoBranchTreeFromMidSortAndAfterSort(nums1, nums2 []int) *tree.Node {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	val := nums2[len(nums2)-1]
	node := &tree.Node{Left: nil, Right: nil, Val: val}
	position := 0
	for i := 0; i < len(nums1); i++ {
		if val == nums1[i] {
			position = i
		}
	}

	node.Left = buildTwoBranchTreeFromMidSortAndAfterSort(nums1[:position], nums2[:position])
	node.Right = buildTwoBranchTreeFromMidSortAndAfterSort(nums1[position+1:], nums2[position:len(nums2)-1])

	return node
}

// 构造最大二叉树
// 根最大，最大左边为左子树，右边为右子树
func buildMaxTwoBranchTree(nums []int) *tree.Node {
	if len(nums) == 0 {
		return nil
	}

	//查找根的位置
	max, position := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			position = i
		}
	}

	node := &tree.Node{Left: nil, Right: nil, Val: max}
	node.Left = buildMaxTwoBranchTree(nums[:position])
	node.Right = buildMaxTwoBranchTree(nums[position+1:])

	return node
}

// 前序和后续构建二叉树 - 多种结果返回任意一种即可
func buildTwoBranchTreeFromPrevSortAndAfterSort(nums1, nums2 []int) *tree.Node {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	// 前序数组长度为1时，直接返回该节点，避免数组越界
	if len(nums1) == 1 {
		return &tree.Node{Left: nil, Right: nil, Val: nums1[0]}
	}

	val := nums1[0]
	node := &tree.Node{Left: nil, Right: nil, Val: val}
	position := 0
	for i := 0; i < len(nums2); i++ {
		if nums1[1] == nums2[i] {
			position = i
		}
	}

	node.Left = buildTwoBranchTreeFromPrevSortAndAfterSort(nums1[1:position+2], nums2[:position+1])
	node.Right = buildTwoBranchTreeFromPrevSortAndAfterSort(nums1[position+2:], nums2[position+1:len(nums2)-1])

	return node
}

func main() {
	fmt.Println("1. 根据前序和中序构建二叉树：")
	nums1 := []int{3, 9, 20, 15, 7}
	nums2 := []int{9, 3, 15, 20, 7}
	fmt.Println("原二叉树的前序和中序遍历结果分别为：", nums1, ",", nums2)
	twoBranchTreeFromPrevSortAndMidSort := buildTwoBranchTreeFromPrevSortAndMidSort(nums1, nums2)
	fmt.Println("构建后的二叉树为：")
	tree.EchoTwoBranchTree(twoBranchTreeFromPrevSortAndMidSort)

	line.SplitLine()

	fmt.Println("2. 根据中序和后续构建二叉树：")
	nums3 := []int{1, 9, 3, 15, 20, 7}
	nums4 := []int{1, 9, 15, 7, 20, 3}
	fmt.Println("原二叉树的中序遍历和后续遍历为：", nums3, ",", nums4)
	treeFromMidSortAndAfterSort := buildTwoBranchTreeFromMidSortAndAfterSort(nums3, nums4)
	fmt.Println("构建后的二叉树为：")
	tree.EchoTwoBranchTree(treeFromMidSortAndAfterSort)

	line.SplitLine()

	fmt.Println("3. 构造最大二叉树：")
	nums := []int{3, 2, 1, 6, 0, 5}
	fmt.Println("给出的数组为：", nums)
	maxTwoBranchTree := buildMaxTwoBranchTree(nums)
	fmt.Println("构造后的最大二叉树为：")
	tree.EchoTwoBranchTree(maxTwoBranchTree)

	line.SplitLine()

	fmt.Println("4. 前序和后续构建二叉树：")
	nums5 := []int{3, 9, 20, 15, 7}
	nums6 := []int{9, 15, 7, 20, 3}
	fmt.Println("二叉树的前序遍历和后续遍历的结果为：", nums5, ",", nums6)
	treeFromPrevSortAndAfterSort := buildTwoBranchTreeFromPrevSortAndAfterSort(nums5, nums6)
	fmt.Println("构造后的二叉树为：")
	tree.EchoTwoBranchTree(treeFromPrevSortAndAfterSort)
}
