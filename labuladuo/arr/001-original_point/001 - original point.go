package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 1. 排序数组元素去重
func noRepeatSortedArr(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	fast, slow := 0, 0
	for fast < len(arr) {
		//找到不相等的第一个数
		if arr[fast] != arr[slow] {
			slow++
			arr[slow] = arr[fast]
		}
		fast++
	}

	return arr[:slow+1]
}

// 2. 原地移除数组中特定的元素
func removeElementFromArr(arr []int, target int) []int {
	if len(arr) < 1 {
		return arr
	}

	fast, slow := 0, 0
	for fast < len(arr) {
		if arr[fast] != target {
			arr[slow] = arr[fast]
			slow++
		}
		fast++
	}

	return arr[:slow]
}

// 3. 将数组的元素0移除到最后, 且保持非零元素的顺序
func moveZeroToFinalFromArr(arr []int) []int {
	length := len(arr)
	if length < 1 {
		return arr
	}

	left, right := 0, 0
	for right < length {
		if arr[right] != 0 {
			// 相对顺序不变
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
		right++
	}

	return arr
}

func main() {
	fmt.Println("1. 排序数组去重：")
	arr := []int{1, 2, 2, 5, 5, 7}
	fmt.Println("去重前的数组为：", arr)
	sortedArrLength := noRepeatSortedArr(arr)
	fmt.Println("去重之后的数组为：", sortedArrLength)

	line.SplitLine()

	fmt.Println("2. 移除特定的元素：")
	arr1 := []int{1, 2, 2, 5, 5, 7}
	fmt.Println("移除前的数组为：", arr1)
	removedArr := removeElementFromArr(arr1, 2)
	fmt.Println("移除之后的数组为：", removedArr)

	line.SplitLine()

	fmt.Println("3. 将数组中的元素0移到末尾，且保证非0元素的顺序不变：")
	arr2 := []int{1, 3, 0, 5, 0, 7, 0, 9}
	fmt.Println("原数组为：", arr2)
	zeroToFinalFromArr := moveZeroToFinalFromArr(arr2)
	fmt.Println("移动之后的数组为:", zeroToFinalFromArr)
}
