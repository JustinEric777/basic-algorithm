package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 1. 二分查找
func twoPartSearch(arr []int, target int) int {
	if len(arr) < 1 {
		return -1
	}

	left, right, middle := 0, len(arr)-1, 0

	for left < right {
		middle = (left + right) / 2
		if target > arr[middle] {
			left = middle + 1
		} else if target < arr[middle] {
			right = middle - 1
		} else {
			return middle
		}
	}

	return -1
}

// 2. 两数之和（数组有序），仅有唯一答案
func twoSum(arr []int, target int) []int {
	if len(arr) < 2 {
		return []int{}
	}

	left, right := 0, len(arr)-1
	for left < right {
		tempSum := arr[left] + arr[right]
		if tempSum < target {
			left++
		} else if tempSum > target {
			right--
		} else {
			return []int{left, right}
		}
	}

	return []int{}
}

// 3. 反转数组
func reserveArr(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]

		left++
		right--
	}

	return arr
}

// 4. 回文串判断
func isPalindromeArr(arr []int) bool {
	if len(arr) < 2 {
		return false
	}

	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	fmt.Println("1. 二分查找：")
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	target := 15
	fmt.Printf("原数组为：%v，查找的目标值为：%d\n", arr, target)
	key := twoPartSearch(arr, target)
	fmt.Println("查找到的目标值key为：", key)

	line.SplitLine()

	fmt.Println("2. 两数之和：")
	arr1 := []int{1, 3, 5, 7, 9, 11, 13}
	target1 := 20
	fmt.Printf("原数组为：%v，查找的目标值和为：%d\n", arr1, target1)
	keyArr := twoSum(arr1, target1)
	fmt.Println("查找到的目标值key为：", keyArr)

	line.SplitLine()

	fmt.Println("3. 反转数组：")
	arr3 := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Println("反转前数组为：", arr3)
	res := reserveArr(arr3)
	fmt.Println("反转后的数组为：", res)

	line.SplitLine()
	fmt.Println("4. 数组的回文判断：")
	arr4 := []int{1, 3, 5, 7, 5, 3, 1}
	fmt.Println("原数组为：")
	palindromeArr := isPalindromeArr(arr4)
	fmt.Println("判责结果为：", palindromeArr)
}
