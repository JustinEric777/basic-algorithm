package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 0-1 背包问题
// 给你一个可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。其中第 i 个物品的重量为 wt[i]，价值为 val[i]，现在让你用这个背包装物品，最多能装的价值是多少
// case:N = 3, W = 4
// wt = [2, 1, 3]
// val = [4, 2, 3]
func basicPackage(weight, num int, wt, val []int) int {
	// init and base case
	dp := make([][]int, num+1)
	for k := range dp {
		dp[k] = make([]int, weight+1)
	}

	for n := 1; n <= num; n++ {
		for w := 1; w <= weight; w++ {
			if w-wt[n-1] < 0 {
				dp[n][w] = dp[n-1][w]
			} else {
				// 前n- 1个挑最大值
				dp[n][w] = max(dp[n-1][w], dp[n-1][w-wt[n-1]]+val[n-1])
			}
		}
	}

	return dp[num][weight]
}

// 子背包问题
// 输入一个只包含正整数的非空数组 nums，请你写一个算法，判断这个数组是否可以被分割成两个子集，使得两个子集的元素和相等
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	// base case and init
	dp := make([]bool, sum+1)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := sum; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}

	return dp[sum]
}

// 完全背包问题
// 零钱兑换：给定不同面额的硬币 coins 和一个总金额 amount，写一个函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个
func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println("1. 0-1 背包问题：")
	N := 3
	W := 4
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}
	maxVal := basicPackage(W, N, wt, val)
	fmt.Println("可获取的最大价值为：", maxVal)

	line.SplitLine()

	fmt.Println("2. 子背包问题：")
	nums := []int{1, 5, 11, 5}
	res := canPartition(nums)
	fmt.Println("最终结果为：", res)

	line.SplitLine()

	fmt.Println("3. 完全背包问题：")
}
