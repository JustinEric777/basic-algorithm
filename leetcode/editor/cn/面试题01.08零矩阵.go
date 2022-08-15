//编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
//
//
//
// 示例 1：
//
// 输入：
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//输出：
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
//
//
// 示例 2：
//
// 输入：
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//输出：
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]
//
// Related Topics 数组 哈希表 矩阵 👍 49 👎 0

package main

import (
	"fmt"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	var nums []string
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				nums = append(nums, fmt.Sprintf("%d%s%d", i, "_", j))
			}
		}
	}

	for k := 0; k < len(nums); k++ {
		arr := strings.Split(nums[k], "_")
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		setZero(matrix, x, y)
	}
}

func setZero(matrix [][]int, x, y int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[x][i] = 0
	}

	for j := 0; j < len(matrix); j++ {
		matrix[j][y] = 0
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println("pre nums :", nums)

	setZeroes(nums)

	fmt.Println("after nums :", nums)

	nums1 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	fmt.Println("pre nums :", nums1)

	setZeroes(nums1)

	fmt.Println("after nums :", nums1)
}
