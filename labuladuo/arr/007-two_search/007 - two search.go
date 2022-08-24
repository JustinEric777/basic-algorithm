package main

import "fmt"

// 双闭区间二分查找
func midTwoSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	return -1
}

// 左边界二分查找，左闭右开
func leftTwoSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1
	}

	return left
}

func rightTwoSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if right == len(nums) || nums[left-1] != target {
		return -1
	}

	return left - 1
}

func main() {
	fmt.Println("1. 二分查找：（双闭区间）")
	nums := []int{1, 2, 2, 2, 3}
	target := 2
	fmt.Println("原数组和查找的元素分别为：", nums, ",", target)
	midSearch := midTwoSearch(nums, target)
	fmt.Println("双闭区间查找的结果为：", midSearch)
	leftSearch := leftTwoSearch(nums, target)
	fmt.Println("左闭右开区间查找的结果为：", leftSearch)
	rightSearch := rightTwoSearch(nums, target)
	fmt.Println("左开右闭区间查找的结果为：", rightSearch)
}
