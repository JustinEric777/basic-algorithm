package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

const MaxInt = int(^uint(0) >> 1)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 最小路径和
// 现在给你输入一个二维数组 grid，其中的元素都是非负整数，
// 现在你站在左上角，只能向右或者向下移动，需要到达右下角。现在请你计算，经过的路径和最小是多少？
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	// base case and init
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[0][j] = dp[0][j-1] + grid[0][j]
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

// 加权最短路径
// 有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [from, to, price] ，表示该航班都从城市 from 开始，以价格 price 抵达 toi。
// 现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的 价格最便宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1
// n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, K = 1
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// 所有数据整理为入度方向的数据结构
	hashMap := make(map[int][][]int)
	for i := 0; i < len(flights); i++ {
		from := flights[i][0]
		to := flights[i][1]
		price := flights[i][2]
		if _, ok := hashMap[to]; ok {
			hashMap[to] = append(hashMap[to], []int{from, price})
		} else {
			hashMap[to] = [][]int{{from, price}}
		}
	}

	return dpFunc(hashMap, src, dst, k)
}

func dpFunc(hashMap map[int][][]int, src, dst, k int) int {
	if src == dst {
		return 0
	}
	if k < 0 {
		return -1
	}

	// dp 目的仅能避免重复计算
	res := MaxInt
	if _, ok := hashMap[dst]; ok {
		for _, v := range hashMap[dst] {
			from := v[0]
			price := v[1]
			subProblem := dpFunc(hashMap, src, from, k-1)
			if subProblem != -1 {
				res = min(res, subProblem+price)
			}
		}
	}

	if res == MaxInt {
		return -1
	}

	return res
}

// 地下城游戏
// 一些恶魔抓住了公主（P）并将她关在了地下城的右下角。
// 地下城是由 M x N 个房间组成的二维网格。我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。
// 骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。
// 有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
// 为了尽快到达公主，骑士决定每次只向右或向下移动一步。
// 编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数
func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	dp := make([]int, n+1)

	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			// base case
			if i == m || j == n {
				dp[j] = MaxInt
				continue
			}
			if i == m-1 && j == n-1 {
				if dungeon[i][j] < 0 {
					dp[j] = -dungeon[i][j] + 1
				} else {
					dp[j] = 1
				}
				continue
			}

			// 迭代
			res := min(dp[j], dp[j+1]) - dungeon[i][j]
			if res <= 0 {
				dp[j] = 1
			} else {
				dp[j] = res
			}
		}
	}

	return dp[0]
}

// 自由之路
// 电子游戏“辐射4”中，任务 “通向自由” 要求玩家到达名为 “Freedom Trail Ring” 的金属表盘，并使用表盘拼写特定关键词才能开门。
// 给定一个字符串 ring ，表示刻在外环上的编码；给定另一个字符串 key ，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
// 最初，ring 的第一个字符与 12:00 方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00 方向对齐，然后按下中心按钮，以此逐个拼写完 key 中的所有字符。
// 旋转 ring 拼出 key 字符 key[i] 的阶段中：
// 您可以将 ring 顺时针或逆时针旋转 一个位置 ，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00 方向对齐，并且这个字符必须等于字符 key[i] 。
// 如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key 的下一个字符（下一阶段）, 直至完成所有拼写。
// 输入: ring = "godding", key = "gd"
// 输出: 4
func findRotateSteps(ring string, key string) int {
	if len(key) == 0 {
		return 0
	}

	var res int
	for k, v := range key {
		step := min(dp(ring, string(v), true), dp(ring, string(v), false))
		res = res + step + 1
		ring = ring[k:] + ring[:k]
	}

	return res
}

func dp(ring string, word string, isRing bool) int {
	var count int
	if isRing {
		for i := 0; i < len(ring); i++ {
			if string(ring[i]) == word {
				break
			}
			count++
		}
	} else {
		for i := len(ring) - 1; i > 0; i-- {
			if string(ring[i]) == word {
				break
			}
			count++
		}
	}

	return count
}

// 正则表达式
// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖整个 字符串 s的，而不是部分字符串
// 输入：s = "aa", p = "a"
// 输出：false
func regxIsMatch(s string, p string) bool {
	if len(p) == 0 {
		return false
	}

	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if string(p[j-1]) == "." {
			return true
		}

		return s[i-1] == p[j-1]
	}

	// init and base
	dp := make([][]bool, len(s)+1)
	for k := range dp {
		dp[k] = make([]bool, len(p)+1)
	}
	dp[0][0] = true

	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				// 匹配0
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if matches(i, j) {
					// 匹配多次
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if matches(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}

	return dp[len(s)][len(p)]
}

// 高楼扔鸡蛋
// 给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。
// 已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。
// 每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。
// 请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？
// 输入：k = 1, n = 2
// 输出：2
func superEggDrop(k int, n int) int {
	if k == 0 {
		return 0
	}

	// init and base case
	dp := make([][]int, k)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	dp[0][0] = 0

	for i := 1; i <= k; i++ {
		for j := 1; j <= n; j++ {
			//鸡蛋没碎
			res1 := 1 + superEggDrop(k, n-j)

			// 鸡蛋碎了
			res2 := 1 + superEggDrop(k-1, j)
			dp[i][j] = min(res1, res2)
		}
	}

	return dp[k][n]
}

// 戳气球
// 有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
// 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
// 求所能获得硬币的最大数量
// 输入：nums = [3,1,5,8] 输出： 167
func maxCoins(nums []int) int {
	// 填充变幻数组
	n := len(nums)
	newNums := make([]int, n+2)
	for k, v := range nums {
		newNums[k+1] = v
	}
	newNums[0], newNums[n+1] = 1, 1

	// init and base case
	dp := make([][]int, n+2)
	for k := range dp {
		dp[k] = make([]int, n+2)
		for j := 0; j < len(dp[k]); j++ {
			dp[k][j] = -1
		}
	}

	return traverMaxCoins(newNums, dp, 0, n+1)
}

func traverMaxCoins(nums []int, dp [][]int, left, right int) int {
	if left >= right-1 {
		return 0
	}
	if dp[left][right] != -1 {
		return dp[left][right]
	}

	for i := left + 1; i < right; i++ {
		sum := nums[left] * nums[i] * nums[right]
		sum += traverMaxCoins(nums, dp, left, i) + traverMaxCoins(nums, dp, i, right)
		dp[left][right] = max(dp[left][right], sum)
	}

	return dp[left][right]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 博弈问题1：预测赢家
// 给你一个整数数组 nums 。玩家 1 和玩家 2 基于这个数组设计了一个游戏。
// 玩家 1 和玩家 2 轮流进行自己的回合，玩家 1 先手。开始时，两个玩家的初始分值都是 0 。
// 每一回合，玩家从数组的任意一端取一个数字（即，nums[0] 或 nums[nums.length - 1]），取到的数字将会从数组中移除（数组长度减 1 ）。
// 玩家选中的数字将会加到他的得分上。当数组中没有剩余数字可取时，游戏结束。
// 如果玩家 1 能成为赢家，返回 true 。如果两个玩家得分相等，同样认为玩家 1 是游戏的赢家，也返回 true 。
// 你可以假设每个玩家的玩法都会使他的分数最大化
func predictTheWinner(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	scoreA, scoreB, score := 0, 0, 0
	for {
		// A 先手
		nums, score = getMaxScore(nums)
		scoreA += score
		if len(nums) == 0 {
			break
		}

		// B
		nums, score = getMaxScore(nums)
		scoreB += score
		if len(nums) == 0 {
			break
		}
	}

	if scoreA >= scoreB {
		return true
	}

	return false
}

func getMaxScore(nums []int) ([]int, int) {
	maxVal := 0
	if nums[0] > nums[len(nums)-1] {
		maxVal = nums[0]
		nums = nums[1:]
	} else {
		maxVal = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
	}

	return nums, maxVal
}

// 博弈问题2：石子游戏
// Alice 和 Bob 用几堆石子在做游戏。一共有偶数堆石子，排成一行；每堆都有 正 整数颗石子，数目为 piles[i] 。
// 游戏以谁手中的石子最多来决出胜负。石子的 总数 是 奇数 ，所以没有平局。
// Alice 和 Bob 轮流进行，Alice 先开始 。 每回合，玩家从行的 开始 或 结束 处取走整堆石头。
// 这种情况一直持续到没有更多的石子堆为止，此时手中 石子最多 的玩家 获胜 。
// 假设 Alice 和 Bob 都发挥出最佳水平，当 Alice 赢得比赛时返回 true ，当 Bob 赢得比赛时返回 false
func stoneGame() {
	var aa string
	switch {
	case aa == "1":
		return
	}
}

// 4键键盘
// 假设你有一个特殊的键盘，上面只有四个键，它们分别是：
// 1、A 键：在屏幕上打印一个 A。
// 2、Ctrl-A 键：选中整个屏幕。
// 3、Ctrl-C 键：复制选中的区域到缓冲区。
// 4、Ctrl-V 键：将缓冲区的内容输入到光标所在的屏幕上。
// 这就和我们平时使用的全选复制粘贴功能完全相同嘛，只不过题目把 Ctrl 的组合键视为了一个键。
// 现在要求你只能进行 N 次操作，请你计算屏幕上最多能显示多少个 A？
func fourKeyBoards() {

}

// 打家劫舍1
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额
func rob1(nums []int) int {

	return 0
}

// 打家劫舍2
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额
func rob2(nums []int) int {

	return 0
}

// 打家劫舍3
// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额
func rob3(nums []int) int {

	return 0
}

// 房屋偷盗
// 一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响小偷偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组 nums ，请计算 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额
func rob4(nums []int) int {

	return 0
}

// 环形房屋
// 一个专业的小偷，计划偷窃一个环形街道上沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
// 给定一个代表每个房屋存放金额的非负整数数组 nums ，请计算 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额
func rob5(nums []int) int {

	return 0
}

// 股票买卖
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0
func maxProfit1(prices []int) int {

	return 0
}

// 股票买卖2
// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//返回 你能获得的 最大 利润
func maxProfit2(prices []int) int {

	return 0
}

// 股票买卖3
// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）
func maxProfit3(prices []int) int {

	return 0
}

// 股票买卖4
// 给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
func maxProfit4(prices []int) int {

	return 0
}

// 最佳买卖股票时机含冷冻期
// 给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）
func maxProfit5(prices []int) int {

	return 0
}

// 买卖股票的最佳时机含手续费
// 给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
// 返回获得利润的最大值
// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费
func maxProfit6(prices []int, fee int) int {

	return 0
}

// 股票的最大利润
// 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
func maxProfit7(prices []int, fee int) int {

	return 0
}

// 有限状态机之KMP
func kmp() {

}

func main() {
	fmt.Println("1. 最小路径和：- dp")
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println("原网格为：", grid)
	minPath := minPathSum(grid)
	fmt.Println("最后路径和为：", minPath)

	line.SplitLine()

	fmt.Println("2. 加权最短路径：")
	n := 3
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src := 0
	dst := 2
	k := 1
	cheapestPricePath := findCheapestPrice(n, flights, src, dst, k)
	fmt.Println("最便宜的路径为：", cheapestPricePath)

	line.SplitLine()

	fmt.Println("3. 地下城游戏：")
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	minimumHP := calculateMinimumHP(dungeon)
	fmt.Println("需要的最小生命值为：", minimumHP)

	line.SplitLine()

	fmt.Println("4. 自由之路：")
	ring := "godding"
	key := "gd"
	minSteps := findRotateSteps(ring, key)
	fmt.Println("最少的步骤为：", minSteps)

	line.SplitLine()

	fmt.Println("5. 正则匹配：")
	str := "aa"
	p := "a.*"
	isMatch := regxIsMatch(str, p)
	fmt.Println("匹配的结果为：", isMatch)

	line.SplitLine()

	fmt.Println("6. 硬币的最大值：")
	nums := []int{3, 1, 5, 8}
	maxCoinVal := maxCoins(nums)
	fmt.Println("获取到的硬币最大值为：", maxCoinVal)

	line.SplitLine()

	fmt.Println("7. 预测赢家：")
	nums1 := []int{1, 5, 2}
	fmt.Println("原始数组为：", nums1)
	theWinner := predictTheWinner(nums1)
	fmt.Println("预测结果为：", theWinner)

	line.SplitLine()
}
