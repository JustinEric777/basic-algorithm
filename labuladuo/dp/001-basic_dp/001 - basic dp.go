package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
	"math"
)

// 三要素：重叠子问题、最优子结构、状态转移方程

// 斐波那契数
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// 自顶向下
func fibDpTopToDown(n int) int {
	dpFib := make(map[int]int)

	var traverse func(n int) int
	traverse = func(n int) int {
		if n == 1 || n == 2 {
			return 1
		}
		if dpFib[n] != 0 {
			return dpFib[n]
		}
		dpFib[n] = traverse(n-1) + traverse(n-2)

		return dpFib[n]
	}

	return traverse(n)
}

// 自下向上
func fibDpDownToTop(n int) int {
	fibDp := make(map[int]int)
	// base case
	fibDp[0] = 0
	fibDp[1] = 1

	// 状态转义方程
	for i := 2; i < n+1; i++ {
		fibDp[i] = fibDpDownToTop(i-1) + fibDpDownToTop(i-2)
	}

	return fibDp[n]
}

// 降低空间复杂度
func fibDpMinSpace(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	var d_i int
	d_i_2 := 0
	d_i_1 := 1
	for i := 2; i < n+1; i++ {
		d_i = d_i_1 + d_i_2
		d_i_2 = d_i_1
		d_i_1 = d_i
	}

	return d_i
}

// 最少的硬币数目/零钱兑换
// K 中面值的硬币, amount 为总金额，最少几枚硬币可以凑出该金额， 不能返回 -1
// 递归
func coinChange(coins []int, amount int) int {
	// base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	res := int(^uint(0) >> 1)
	for i := 0; i < len(coins); i++ {
		sub := coinChange(coins, amount-coins[i])
		if sub == -1 {
			continue
		}
		res = int(math.Min(float64(res), float64(sub+1)))
	}

	return res
}

// dp + 备忘录递归
func coinChangeDp(coins []int, amount int) int {
	dp := make(map[int]int)
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if _, result := dp[amount]; result {
		return dp[amount]
	}

	res := int(^uint(0) >> 1)
	for i := 0; i < len(coins); i++ {
		sub := coinChangeDp(coins, amount-coins[i])
		if sub == -1 {
			continue
		}
		dp[amount] = int(math.Min(float64(res), float64(sub+1)))
	}

	if _, result := dp[amount]; result {
		return dp[amount]
	}

	return -1
}

// dp 数组迭代 - 自下向上
func coinChangeDp1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//init dp
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0

	// 子问题填充处理
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 { // 不满足跳过
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

// 最长递增子序列 - 递归
func longestIncreasingSubsequence(nums []int) int {
	res := 0
	if len(nums) == 1 {
		return 1
	}

	for i := 0; i < len(nums); i++ {
		var subProblem int
		minValue := getMin(nums[i+1:])
		if nums[i] <= minValue {
			subProblem = longestIncreasingSubsequence(nums[i+1:]) + 1
		} else {
			subProblem = longestIncreasingSubsequence(nums[i+1:])
		}

		res = int(math.Max(float64(res), float64(subProblem)))
	}

	return res
}

func getMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := int(^uint(0) >> 1)
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}

	return res
}

// 最长递增子序列 - 自下向上
func longestIncreasingSubsequenceDown(nums []int) int {
	// dp
	dp := make(map[int]int)

	//base case
	dp[0] = 0

	// 状态转移
	for i := 0; i < len(nums); i++ {
		max := getMax(nums[:i], dp, nums[i])
		dp[nums[i]] = max + 1
	}

	// res
	res := 0
	for _, v := range dp {
		res = int(math.Max(float64(res), float64(v)))
	}

	return res
}

func getMax(nums []int, dp map[int]int, i int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	for _, v := range nums {
		if v <= i {
			res = int(math.Max(float64(dp[v]), float64(res)))
		}
	}

	return res
}

type SubString struct {
	Len int
	Str string
}

// 最长回文子序列
func longestPalindromeSubsequence(str string) (int, string) {
	// dp - 表示当前结尾的回文序列长度
	dp := make([]*SubString, len(str))

	// init
	for k, _ := range str {
		s := &SubString{
			Len: 0,
			Str: "",
		}
		dp[k] = s
	}

	// 循环迭代
	for i := 1; i < len(str); i++ {
		dp[i] = dp[i-1]
		for j := 0; j < i; j++ {
			if str[j] == str[i] {
				dp[i].Len = dp[i-1].Len + 2
				dp[i].Str = fmt.Sprintf("%s%s%s", string(str[i]), dp[i-1].Str, string(str[i]))
			} else {
				continue
			}
		}
	}

	// max
	var (
		res         int
		Subsequence string
	)
	for _, val := range dp {
		if val.Len > res {
			res = val.Len
			Subsequence = val.Str
		}
	}

	return res, Subsequence
}

// 最长回文子串 - 中心向两边扩展
func longestPalindromeSubStr(str string) (int, string) {

	dp := make([]*SubString, len(str))

	// base case and init
	for k, _ := range str {
		s := &SubString{
			Len: 1,
			Str: string(str[k]),
		}
		dp[k] = s
	}

	// 循环迭代 - i代表中间
	for i := 0; i < len(str); i++ {
		left, right := i-1, i+1
		if left >= 0 && str[left] == str[i] {
			dp[i].Len = dp[i].Len + 1
			dp[i].Str = fmt.Sprintf("%s%s", string(str[left]), dp[i].Str)
			left--
		}
		if right < len(str) && str[right] == str[i] {
			dp[i].Len = dp[i].Len + 1
			dp[i].Str = fmt.Sprintf("%s%s", dp[i].Str, string(str[right]))
			right++
		}
		for left >= 0 && right < len(str) {
			if str[left] == str[right] {
				dp[i].Len = dp[i].Len + 2
				dp[i].Str = fmt.Sprintf("%s%s%s", string(str[left]), dp[i].Str, string(str[right]))
				left--
				right++
			} else {
				break
			}
		}
	}

	// 求最大值
	var (
		res    int
		subStr string
	)
	for _, val := range dp {
		if val.Len > res {
			res = val.Len
			subStr = val.Str
		}
	}

	return res, subStr
}

// 最长的回文子串 - dp
func longestPalindromeSubStr1(str string) (int, string) {
	dp := make([]*SubString, len(str))

	// base case
	for k, _ := range str {
		s := &SubString{
			Len: 0,
			Str: "",
		}
		dp[k] = s
	}

	// 循环迭代
	for i := 1; i < len(str); i++ {
		dp[i] = dp[i-1]
		if dp[i].Len == 0 {
			if i > 1 && str[i] == str[i-1] {
				dp[i].Len = dp[i].Len + 2
				dp[i].Str = fmt.Sprintf("%s%s", string(str[i]), string(str[i]))
			}
			if i > 2 && str[i] == str[i-2] {
				dp[i].Len = dp[i].Len + 3
				dp[i].Str = fmt.Sprintf("%s%s%s", string(str[i]), string(str[i-1]), string(str[i]))
			}
		} else {
			if str[i] == str[i-dp[i-1].Len-1] {
				dp[i].Len = dp[i].Len + 2
				dp[i].Str = fmt.Sprintf("%s%s%s", string(str[i]), dp[i].Str, string(str[i]))
			}
		}
	}

	// max
	var (
		res    int
		subStr string
	)
	for _, val := range dp {
		if val.Len > res {
			res = val.Len
			subStr = val.Str
		}
	}

	return res, subStr
}

// 俄罗斯套娃问题
// 给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度
// 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样
// 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）
// sort - 先对宽度 w 进行升序排序，如果遇到 w 相同的情况，则按照高度 h 降序排序；之后把所有的 h 作为一个数组，在这个数组上计算 LIS 的长度就是答案
func maxEnvelopes(nums [][]int) int {
	dp := make(map[string][][]int)

	// init
	for _, v := range nums {
		key := fmt.Sprintf("%d%s%d", v[0], "_", v[1])
		dp[key] = [][]int{{v[0], v[1]}}
	}

	for i := 0; i < len(nums); i++ {
		// get max pos
		var positions [][]int
		preMax := 0
		for j := 0; j < i; j++ {
			preKey := fmt.Sprintf("%d%s%d", nums[j][0], "_", nums[j][1])
			if preMax < len(dp[preKey]) {
				positions = dp[preKey]
			}
		}

		key := fmt.Sprintf("%d%s%d", nums[i][0], "_", nums[i][1])
		for position, _ := range positions {
			dp[key] = positions
			// judge max and min 当前元素为最大或者最小
			if (nums[i][0] > nums[position][0] && nums[i][1] > nums[position][1]) ||
				(nums[i][0] < nums[position][0] && nums[i][1] < nums[position][1]) {
				dp[key] = append(dp[key], nums[i])
			}
		}
	}

	res := 0
	for _, v := range dp {
		res = int(math.Max(float64(res), float64(len(v))))
	}

	return res
}

// 下降路径最小和
// 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和
// 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素
// 在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）
// 具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1)
func minFallingPathSum(nums [][]int) int {
	dp := make(map[string]int)

	// init last row
	for k := 0; k < len(nums); k++ {
		key := fmt.Sprintf("%d%s%d", len(nums)-1, "_", k)
		dp[key] = nums[len(nums)-1][k]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < len(nums); j++ {
			// get min value
			minVal := int(^uint(0) >> 1)
			if j > 0 {
				nextKey1 := fmt.Sprintf("%d%s%d", i+1, "_", j-1)
				minVal = int(math.Min(float64(minVal), float64(dp[nextKey1])))

			}
			nextKey2 := fmt.Sprintf("%d%s%d", i+1, "_", j)
			minVal = int(math.Min(float64(minVal), float64(dp[nextKey2])))
			if j < len(nums)-1 {
				nextKey3 := fmt.Sprintf("%d%s%d", i+1, "_", j+1)
				minVal = int(math.Min(float64(minVal), float64(dp[nextKey3])))
			}

			// set current val
			key := fmt.Sprintf("%d%s%d", i, "_", j)
			dp[key] = nums[i][j] + minVal
		}
	}

	// get min val
	res := int(^uint(0) >> 1)
	for i := 0; i < len(nums); i++ {
		key := fmt.Sprintf("%d%s%d", 0, "_", i)
		if dp[key] < res {
			res = dp[key]
		}
	}

	return res
}

func main() {
	fmt.Println("1. 斐波那契数列：")
	n := 10
	res := fib(n)
	fmt.Println("斐波那契数列f(", n, ")")
	fmt.Println("暴力结果为：", res)

	line.SplitDottedLine()

	res1 := fibDpTopToDown(n)
	fmt.Println("斐波那契数列f(", n, ")")
	fmt.Println("dp(自顶向下)存储中间数组的结果为：", res1)

	line.SplitDottedLine()

	res2 := fibDpDownToTop(n)
	fmt.Println("斐波那契数列f(", n, ")")
	fmt.Println("dp(自下向上)存储中间数组的结果为：", res2)

	line.SplitDottedLine()

	res3 := fibDpMinSpace(n)
	fmt.Println("斐波那契数列f(", n, ")")
	fmt.Println("降低空间复杂度的结果为：", res3)

	line.SplitLine()

	fmt.Println("2. 最小的硬币数目/零钱凑数：")
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println("需要凑钱的数目和硬币集合为：", amount, ",", coins)
	coinsNum := coinChange(coins, amount)
	fmt.Println("递归解法最少需要硬币数量为:", coinsNum)

	line.SplitDottedLine()

	fmt.Println("需要凑钱的数目和硬币集合为：", amount, ",", coins)
	coinsNum1 := coinChangeDp(coins, amount)
	fmt.Println("递归解法+备忘录最少需要硬币数量为:", coinsNum1)

	line.SplitDottedLine()

	fmt.Println("需要凑钱的数目和硬币集合为：", amount, ",", coins)
	coinsNum2 := coinChangeDp1(coins, amount)
	fmt.Println("dp迭代最少需要硬币数量为:", coinsNum2)

	line.SplitLine()

	fmt.Println("3. 最长递增序列：")
	nums := []int{10, 9, 2, 5, 3, 3, 7, 101, 18}
	fmt.Println("原序列为：", nums)
	maxSubSequenceLength := longestIncreasingSubsequence(nums)
	fmt.Println("最长递增子序列长度为：", maxSubSequenceLength)

	line.SplitDottedLine()

	fmt.Println("原序列为：", nums)
	maxSubSequenceLength1 := longestIncreasingSubsequenceDown(nums)
	fmt.Println("最长递增子序列长度为(dp)：", maxSubSequenceLength1)

	line.SplitLine()

	fmt.Println("4. 最长回文子序列：")
	str := "ABCDDCEFA"
	fmt.Println("原字符串为：", str)
	length1, subsequence := longestPalindromeSubsequence(str)
	fmt.Println("最长回文队列为：", subsequence, "，长度为：", length1)

	line.SplitLine()

	fmt.Println("5. 最长的回文子串：")
	str1 := "ABCDDCEFA"
	fmt.Println("原字符串为：", str1)
	length, subStr := longestPalindromeSubStr(str1)
	fmt.Println("最长回文字符串的为 - 中心向两边扩展：", subStr, "，长度为：", length)

	line.SplitDottedLine()
	length2, subStr1 := longestPalindromeSubStr1(str1)
	fmt.Println("最长回文字符串的为 - dp：", subStr1, "，长度为：", length2)

	line.SplitLine()

	fmt.Println("6. 俄罗斯套娃：")
	envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	fmt.Println("原信封为：", envelopes)
	maxEv := maxEnvelopes(envelopes)
	fmt.Println("最大的套娃信封个数为：", maxEv)

	line.SplitLine()

	fmt.Println("7. 下降路径最小和：")
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	fmt.Println("原数组为：", matrix)
	fallingPathSum := minFallingPathSum(matrix)
	fmt.Println("下降路径的最小和为：", fallingPathSum)
}
