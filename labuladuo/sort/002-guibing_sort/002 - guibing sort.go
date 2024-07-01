package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

var count map[int]int

// 计算右侧小于当前元素的个数
func smallElementNums(nums []int) []int {
	if len(nums) <= 1 {
		return []int{0}
	}

	count = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		count[nums[i]] = 0
	}

	sort(nums)

	var countArr []int
	for _, v := range count {
		countArr = append(countArr, v)
	}

	return countArr
}

// 区间和的个数
// 给你一个整数数组 nums 以及两个整数 lower 和 upper 。求数组中，值位于范围 [lower, upper] （包含 lower 和 upper）之内的 区间和的个数
// 区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)
func intervalSum(nums []int, lower, upper int) int {
	return 0
}

// 翻转对
// 给定一个数组nums,如果 i < j && nums[i] > 2 * nums[j], 则（i, j）为一个翻转对，返回翻转对的数量
func flipPair(nums []int) int {
	return 0
}

// 排序数组 - 归并排序
var temp []int

func sort(nums []int) {
	temp = make([]int, len(nums))

	sortGuiBing(nums, 0, len(nums)-1)
}

func sortGuiBing(nums []int, low, high int) {
	if low == high {
		return
	}

	mid := low + (high-low)/2
	sortGuiBing(nums, low, mid)
	sortGuiBing(nums, mid+1, high)

	merge(nums, low, mid, high)
}

func merge(nums []int, low, mid, high int) {
	// cp
	for i := low; i <= high; i++ {
		temp[i] = nums[i]
	}
	left, right := low, mid+1
	for p := low; p <= high; p++ {
		if left == mid+1 {
			nums[p] = temp[right]
			right++
		} else if right == high+1 {
			nums[p] = temp[left]
			count[nums[p]] += right - mid - 1
			left++
		} else if temp[left] > temp[right] {
			nums[p] = temp[right]
			right++
		} else {
			nums[p] = temp[left]
			count[nums[p]] += right - mid - 1
			left++
		}
	}
}

func mergeSortArr(arr1, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	var res []int
	left, right := 0, 0
	for right < len(arr2) && left < len(arr1) {
		if arr1[left] > arr2[right] {
			res = append(res, arr2[right])
			right++
		} else {
			res = append(res, arr1[left])
			left++
		}
	}

	if left < len(arr1) {
		res = append(res, arr1[left:]...)
	}
	if right < len(arr2) {
		res = append(res, arr2[left:]...)
	}

	return res
}

func main() {
	fmt.Println("1. 计算右侧小于当前元素的个数：")
	nums := []int{5, 2, 6, 1}
	fmt.Println("原数组为：", nums)
	elementNums := smallElementNums(nums)
	fmt.Println("右侧小于当前元素的个数为：", elementNums)

	line.SplitLine()

	fmt.Println("2. 区间和的个数：")

	line.SplitLine()

	fmt.Println("3. 翻转树：")

	line.SplitLine()

	fmt.Println("4. 数组排序：")
	arr := []int{5, 2, 3, 1}
	fmt.Println("原数组为：", arr)
	sort(arr)
	fmt.Println("数组排序的结果为：", arr)
	arr1 := []int{5, 2, 1, 1, 0, 0}
	fmt.Println("原数组为：", arr1)
	sort(arr1)
	fmt.Println("数组排序的结果为：", arr1)

	line.SplitLine()

	fmt.Println("5. 合并两个数组：")
	arr3 := []int{1, 4, 6, 9}
	arr4 := []int{2, 5, 8}
	res5 := mergeSortArr(arr3, arr4)
	fmt.Println("需要合并的数组为：", arr3, ",", arr4)
	fmt.Println("合并之后的结果为：", res5)

}
