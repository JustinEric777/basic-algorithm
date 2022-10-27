package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
	"math"
)

// 编辑距离
// 给你两个单词word1 和word2， 请返回将word1转换成word2 所使用的最少操作数
// 你可以对一个单词进行如下三种操作：插入、删除、替换
func minDistance(word1 string, word2 string) int {
	i := len(word1) - 1
	j := len(word2) - 1

	return traverse(word1, word2, i, j)
}

func traverse(s1, s2 string, i, j int) int {
	// base case - 走完了
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}

	// judge skip
	if s1[i] == s2[j] {
		return traverse(s1, s2, i-1, j-1)
	}

	// operate - delete/insert/update
	return min(
		traverse(s1, s2, i-1, j)+1,
		traverse(s1, s2, i, j-1)+1,
		traverse(s1, s2, i-1, j-1)+1)
}

func min(a, b, c int) int {
	return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}

func minDistanceDp(word1, word2 string) int {
	m, n := len(word1), len(word2)
	// dp and init
	dp := make([][]int, m+1)
	for k, _ := range dp {
		arr := make([]int, n+1)
		dp[k] = arr
	}

	// base case
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	// 循环迭代
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[m][n]
}

// 最大子数组和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。子数组 是数组中的一个连续部分
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp 为i结尾的最大元素值
	dp := make([]int, len(nums))

	// base case
	dp[0] = nums[0]

	// case traverse
	for i := 1; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]+nums[i]), float64(nums[i])))
	}

	//求最值
	res := 0
	for _, v := range dp {
		res = int(math.Max(float64(res), float64(v)))
	}

	return res
}

/**
 * 空间压缩
 */
func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// base case
	dpPrefix := nums[0]
	dpCurrent := 0
	res := 0

	// case traverse
	for i := 1; i < len(nums); i++ {
		dpCurrent = int(math.Max(float64(dpPrefix+nums[i]), float64(nums[i])))
		dpPrefix = dpCurrent
		res = int(math.Max(float64(res), float64(dpCurrent)))
	}

	return res
}

/**
 * 前缀和
 */
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	res := 0
	minVal := 1000
	for i := 0; i < len(nums); i++ {
		// 维护 minVal 是 preSum[0..i] 的最小值
		minVal = int(math.Min(float64(minVal), float64(preSum[i])))
		// 以 nums[i] 结尾的最大子数组和就是 preSum[i+1] - min(preSum[0..i])
		res = int(math.Max(float64(res), float64(preSum[i+1]-minVal)))
	}

	return res
}

// 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	// base case init
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		arr := make([]int, len(text2)+1)
		dp[i] = arr
	}

	// 循环迭代
	for m := 1; m < len(text1)+1; m++ {
		for n := 1; n < len(text2)+1; n++ {
			if text1[m-1] == text2[n-1] {
				dp[m][n] = 1 + dp[m-1][n-1]
			} else {
				dp[m][n] = int(math.Max(float64(dp[m-1][n]), float64(dp[m][n-1])))
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

// 两个字符串的删除操作
// 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数，仅删除操作
func minDistance1(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return max(len(word1), len(word2))
	}

	// init and base case
	dp := make([][]int, len(word1)+1)
	for k := range dp {
		dp[k] = make([]int, len(word2)+1)
		dp[k][0] = k
	}
	for k1 := range dp[0] {
		dp[0][k1] = k1
	}

	// 循环迭代
	for m := 1; m < len(word1)+1; m++ {
		for n := 1; n < len(word2)+1; n++ {
			if word1[m-1] == word2[n-1] {
				dp[m][n] = dp[m-1][n-1]
			} else {
				dp[m][n] = 1 + min(dp[m][n-1], dp[m-1][n], len(word1)+len(word2))
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

// 两个字符串的最小ASCII删除和
// 给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和
func minimumDeleteSum(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return max(len(word1), len(word2))
	}

	// init and base case
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(word1[i-1])
		}
	}
	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = dp[0][j-1] + int(word2[j-1])
	}

	// 循环迭代
	for m := 1; m < len(word1)+1; m++ {
		for n := 1; n < len(word2)+1; n++ {
			if int(word1[m-1]) == int(word2[n-1]) {
				dp[m][n] = dp[m-1][n-1]
			} else {
				dp[m][n] = min1(dp[m][n-1]+int(word2[n-1]), dp[m-1][n]+int(word1[m-1]))
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min1(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 让字符串成为回文串的最少插入次数
// 给你一个字符串s，每一次操作你都可以在字符串的任意位置插入任意字符。返回让s成为回文串的最少操作次数
func minInsertions(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	// init and base case
	dp := make([][]int, length)
	for k := range dp {
		dp[k] = make([]int, length)
	}

	// span代表子串长度
	for span := 2; span <= length; span++ {
		for i := 0; i <= length-span; i++ {
			j := i + span - 1
			dp[i][j] = min1(dp[i+1][j]+1, dp[i][j-1]+1)
			if s[i] == s[j] {
				dp[i][j] = min1(dp[i+1][j-1], dp[i][j])
			}
		}
	}

	return dp[0][length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println("1.编辑距离 - 递归：")
	word1 := "horse"
	word2 := "ros"
	word11 := "intention"
	word22 := "execution"
	fmt.Println("原词和变化后的词为：", word1, ",", word2)
	distance := minDistance(word1, word2)
	fmt.Println("word1 -> word2 需要的最少操作数为：", distance)

	line.SplitDottedLine()

	fmt.Println("1.编辑距离 - dp：")
	fmt.Println("原词和变化后的词为：", word11, ",", word22)
	distance1 := minDistanceDp(word11, word22)
	fmt.Println("word1 -> word2 需要的最少操作数为：", distance1)

	line.SplitLine()

	fmt.Println("2. 最大子数组和 - dp：")
	nums := []int{-3, 1, 3, -1, 2, -4, 2}
	sum := maxSubArray(nums)
	fmt.Println("原数组为：", nums)
	fmt.Println("最大子数组和为：", sum)

	line.SplitDottedLine()

	fmt.Println("2. 最大子数组和 - dp - 空间压缩：")
	sum = maxSubArray1(nums)
	fmt.Println("原数组为：", nums)
	fmt.Println("最大子数组和为：", sum)

	line.SplitDottedLine()

	fmt.Println("2. 最大子数组和 - dp - 前缀和：")
	sum = maxSubArray2(nums)
	fmt.Println("原数组为：", nums)
	fmt.Println("最大子数组和为：", sum)

	line.SplitLine()

	fmt.Println("3. 最长公共子序列 - dp：")
	s1 := "zabcde"
	s2 := "acez"
	fmt.Println("原字符串是s1和s2分别为：", s1, ",", s2)
	length := longestCommonSubsequence(s1, s2)
	fmt.Println("最长公共自序列的长度为：", length)

	line.SplitLine()

	fmt.Println("4. 两个字符串的删除操作 - dp:")
	word111 := "sea"
	word222 := "eat"
	fmt.Println("原两个字符串分别为：", word111, ",", word222)
	minStep := minDistance1(word111, word222)
	fmt.Println("删除至相同需要的最小步数为：", minStep)

	line.SplitLine()

	fmt.Println("5. 删除至相同需要的最小Ascii和 - dp:")
	fmt.Println("原两个字符串分别为：", word111, ",", word222)
	minAsciiSum := minimumDeleteSum(word111, word222)
	fmt.Println("删除至相同需要的最小Ascii和为：", minAsciiSum)

	line.SplitLine()

	fmt.Println("6. 让字符串成为回文串的最少插入次数 - dp:")
	str := "leetcode"
	fmt.Println("原字符串为：", str)
	minSteps := minInsertions(str)
	fmt.Println("使其成为回文字符串的最少操作次数为：", minSteps)
}
