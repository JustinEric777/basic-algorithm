package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

type Difference struct {
	Diff []int `json:"nums"`
}

// 构造差数数组
func (difference *Difference) diff(nums []int) {
	difference.Diff = nums
	difference.Diff[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		difference.Diff[i] = nums[i] - nums[i-1]
	}
}

// 区间为【i， j】
func (difference *Difference) increment(i, j, val int) {
	// 仅左区间需要加，其余全抵消了，差值不变
	difference.Diff[i] += val
	// 控制右区间的后位差
	if j+1 < len(difference.Diff) {
		difference.Diff[j+1] -= val
	}
}

// 还原数组
func (difference *Difference) reduction() []int {
	nums := difference.Diff
	nums[0] = difference.Diff[0]

	for i := 1; i < len(difference.Diff); i++ {
		nums[i] = difference.Diff[i-1] + difference.Diff[i]
	}

	return nums
}

// 1. 初始状态为 n 的 0数组， 区间加法，给定多个区间进项相应的操作，输出最后的结果即可
func getMultiModifyArr(n int, updates [][]int) []int {
	nums := make([]int, n)

	difference := &Difference{Diff: nums}
	//difference.diff(nums)

	for i := 0; i < len(updates); i++ {
		difference.increment(updates[i][0], updates[i][1], updates[i][2])
	}

	res := difference.reduction()

	return res
}

// 2. 航班预定统计 - 输入航班表， 【i， j, k】 -i 到 j 上每个航班定了 k 个座位，按照航班顺序返回每个航班上预订的座位数
func flightReservation(n int, updates [][]int) []int {
	nums := make([]int, n)
	difference := &Difference{}

	difference.diff(nums)

	// 索引要减1 对应数组的 key 值
	for i := 0; i < len(updates); i++ {
		difference.increment(updates[i][0]-1, updates[i][1]-1, updates[i][2])
	}

	res := difference.reduction()

	return res
}

// 3. 拼车 - 输入乘客表， [num, start, end] 代表 num 个旅客要从 start 上车， end 下车， 输出：一次能否把所有乘客拉完， 且不能超载， 客车容量为 capacity
func carPool(n, capacity int, updates [][]int) bool {
	nums := make([]int, n)
	difference := &Difference{Diff: nums}

	difference.Diff[0] = updates[0][0]
	for i := 0; i < len(updates); i++ {
		start, end, val := updates[i][1], updates[i][2]-1, updates[i][0]
		difference.increment(start, end, val)
	}

	reduction := difference.reduction()
	for _, v := range reduction {
		if v > capacity {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("1. 数组区间加减：")
	updates := [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}
	length := 5
	fmt.Println("原数组的长度和操作三元组为：", length, updates)
	arr := getMultiModifyArr(length, updates)
	fmt.Println("数组区间加减的结果为：", arr)

	line.SplitLine()

	fmt.Println("2. 航班预定统计：")
	updates1 := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	length1 := 5
	fmt.Println("原数组的长度和操作三元组为：", length1, updates1)
	reservation := flightReservation(length1, updates1)
	fmt.Println("航班预定的统计结果为结果为：", reservation)

	line.SplitLine()

	update2 := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity := 4
	fmt.Println("车的最大容量和操作三元组为：", capacity, update2)
	res := carPool(1001, capacity, update2)
	fmt.Println("拼车的结果为：", res)
}
