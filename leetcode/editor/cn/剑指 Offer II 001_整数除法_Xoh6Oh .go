//给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
//
//
//
// 注意：
//
//
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2³¹, 2³¹−1]。本题中，如果除法结果溢出，则返回 231 − 1
//
//
//
//
// 示例 1：
//
//
//输入：a = 15, b = 2
//输出：7
//解释：15/2 = truncate(7.5) = 7
//
//
// 示例 2：
//
//
//输入：a = 7, b = -3
//输出：-2
//解释：7/-3 = truncate(-2.33333..) = -2
//
// 示例 3：
//
//
//输入：a = 0, b = 1
//输出：0
//
// 示例 4：
//
//
//输入：a = 1, b = 1
//输出：1
//
//
//
// 提示:
//
//
// -2³¹ <= a, b <= 2³¹ - 1
// b != 0
//
//
//
//
// 注意：本题与主站 29 题相同：https://leetcode-cn.com/problems/divide-two-integers/
//
//
//
// Related Topics 位运算 数学 👍 250 👎 0

package main

import (
	"fmt"
	"math"
)

// leetcode submit region begin(Prohibit modification and deletion)
func divide(a int, b int) int {
	if a == math.MinInt32 {
		if b == 1 {
			return math.MinInt32
		}
		if b == -1 {
			return math.MaxInt32
		}
	}
	if b == math.MinInt32 {
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}
	if a == 0 {
		return 0
	}

	rev := false
	if a > 0 {
		a = -a
		rev = !rev
	}
	if b > 0 {
		b = -b
		rev = !rev
	}

	res := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1
		if quickAdd(b, mid, a) {
			res = mid
			if mid == math.MaxInt32 {
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	if rev {
		return -res
	}

	return res
}

// y *z >=x
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 {
		// 不能使用除法
		if z&1 > 0 { // 需要保证 result + add >= x
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 { // 需要保证 add + add >= x
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

func main() {
	i := divide(15, 2)
	fmt.Println(i)
}

//leetcode submit region end(Prohibit modification and deletion)
