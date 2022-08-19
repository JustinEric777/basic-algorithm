package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 1. 一维数组中的前缀和, 求数组【i， j】索引的总和，数组不可变
// 实现numsArr 类
type numsArr struct {
}

var preSum []int

// 暂记前一步计算的结果，直接空间换取时间
func NewNumsArr(arr []int) *numsArr {
	preSum = append(arr, 0)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + arr[i-1]
	}

	return &numsArr{}
}

func (nums *numsArr) sumRange(left, right int) int {
	return preSum[right+1] - preSum[left]
}

// 2. 二维矩阵中的前缀和（计算矩阵范围内元素的总和） - 矩阵不可变 -- 空间换时间 -- 根据原点拼凑
func sumRegion() int {

	return 0
}

func main() {
	fmt.Println("1. 一维数组的前缀和：")
	arr := []int{1, 3, 5, 7, 9, 11}
	fmt.Println("原数组为：", arr)
	newNumsArr := NewNumsArr(arr)
	res1 := newNumsArr.sumRange(0, 3)
	res2 := newNumsArr.sumRange(1, 3)
	res3 := newNumsArr.sumRange(2, 5)
	fmt.Println("计算的结果为：", append([]int{}, res1, res2, res3))

	line.SplitLine()

	fmt.Println("2. 二维数组的前缀和：")

}
