package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 根据权重随机选择/根据权重生成随机数
// 给定一个下标从0开始的权重数组w,w[i]代表第i个小标的权重
// 实现函数可以随机的从闭区间【0， w.length -i] 选出并返回一个下标，选取下标的概率为w[i]/sum[w]
// 前缀和 + 二分查找
type Solution struct {
	PreSum []int
}

var preSum []int

func NewSolution(w []int) *Solution {
	preSum = make([]int, len(w)+1)

	preSum[0] = 0
	for i := 1; i <= len(w); i++ {
		preSum[i] = preSum[i-1] + w[i-1]
	}

	return &Solution{PreSum: preSum}
}

func (solution *Solution) PickIndex() int {
	// 产生随机数
	n := len(solution.PreSum)
	//ns时间戳可以保证每次重启seed都不一样
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Int31n(int32(solution.PreSum[n-1])) + 1
	fmt.Println(randomNum)

	// 寻找最接近的下标
	key := leftTwoSearch(solution.PreSum, int(randomNum))

	// 前缀和数组和 w 有一位偏移 - 0
	return key - 1
}

// 左边界二分查找
func leftTwoSearch(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	return left
}

func main() {
	fmt.Println("1. 根据权重选择随机数：")
	w := []int{1, 3}
	fmt.Println("输入的权重为：", w)
	solution := NewSolution(w)
	index := solution.PickIndex()
	index1 := solution.PickIndex()
	index2 := solution.PickIndex()
	index3 := solution.PickIndex()
	index4 := solution.PickIndex()
	fmt.Println("返回的结果为：", index, index1, index2, index3, index4)
}
