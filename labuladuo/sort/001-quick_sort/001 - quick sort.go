package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 数组中第k个最大元素
func maxKthElementFromArr(nums []int, k int) int {
	sort1(nums, 0, len(nums)-1, k)
	return nums[len(nums)-k]
}

func sort1(nums []int, low, high, k int) {
	if low >= high {
		return
	}

	mid := partition(nums, low, high)
	if mid == k {
		return
	}
	if mid < k {
		sort(nums, mid+1, high)
	} else {
		sort(nums, low, mid-1)
	}
}

func partition1(nums []int, low, high int) int {
	mid := low + (high-low)/2

	left, right := 0, low+1
	for left < mid && right < high+1 {
		if nums[left] > nums[mid] {
			temp := nums[mid]
			nums[mid] = nums[left]
			nums[left] = temp
			mid = left
		} else {
			left++
		}

		if nums[right] < nums[mid] {
			temp := nums[mid]
			nums[mid] = nums[right]
			nums[right] = temp
			mid = right
		} else {
			right++
		}
	}

	return mid
}

// 排序数组
func sort(nums []int, low, high int) {
	if low >= high {
		return
	}

	mid := partition(nums, low, high)
	sort(nums, low, mid-1)
	sort(nums, mid+1, high)
}

func partition(nums []int, low, high int) int {
	mid := low + (high-low)/2

	left, right := 0, low+1
	for left < mid && right < high+1 {
		if nums[left] > nums[mid] {
			temp := nums[mid]
			nums[mid] = nums[left]
			nums[left] = temp
			mid = left
		} else {
			left++
		}

		if nums[right] < nums[mid] {
			temp := nums[mid]
			nums[mid] = nums[right]
			nums[right] = temp
			mid = right
		} else {
			right++
		}
	}

	return mid
}

func main() {
	fmt.Println("1. 数组中第k个最大元素：")
	nums := []int{1, 3, 5, 7, 9, 11, 27, 17, 21}
	fmt.Println("原数组为：", nums)
	k := 3
	maxKth := maxKthElementFromArr(nums, k)
	fmt.Println("第", k, "大元素为：", maxKth)

	line.SplitLine()

	fmt.Println("2. 快速排序：")
	arr := []int{1, 5, 3, 11, 9, 7}
	fmt.Println("原数组为：", arr)
	sort(arr, 0, len(arr)-1)
	fmt.Println("排序后的数组为：", arr)

}
