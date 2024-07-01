package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 1. 顺/逆时针旋转矩阵 - n * n 的二阶矩阵，顺时针在原矩阵旋转90度，仅在原矩阵上操作
func rotateMatrix(matrix [][]int) [][]int {
	if len(matrix) < 2 && len(matrix[0]) < 2 {
		return matrix
	}

	// 矩阵转置
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 按行逆序
	for m := 0; m < len(matrix); m++ {
		for n := 0; n < len(matrix[m])/2; n++ {
			matrix[m][n], matrix[m][len(matrix[m])-1] = matrix[m][len(matrix[m])-1], matrix[m][n]
		}
	}

	return matrix
}

// 2. 矩阵的螺旋遍历 - n * n 的二维矩阵，顺时针螺旋输出所有元素
func spiralMatrix(matrix [][]int) []int {
	// 设定边界
	m, n := len(matrix), len(matrix[0])
	upperBound, lowerBound := 0, m-1
	leftBound, rightBound := 0, n-1

	var res []int
	for len(res) < m*n {
		//顶部 左 -> 右遍历
		if upperBound <= lowerBound {
			for j := leftBound; j <= rightBound; j++ {
				res = append(res, matrix[upperBound][j])
			}

			//上边界下移
			upperBound++
		}

		//右边， 上 -> 下遍历
		if leftBound <= rightBound {
			for i := upperBound; i <= lowerBound; i++ {
				res = append(res, matrix[i][rightBound])
			}

			// 右边界左移
			rightBound--
		}

		//下面 ， 右 -> 左 遍历
		if upperBound <= lowerBound {
			for m := rightBound; m >= leftBound; m-- {
				res = append(res, matrix[lowerBound][m])
			}

			//下边界上移
			lowerBound--
		}

		// 左边， 下 -> 上遍历
		if leftBound <= rightBound {
			for n := lowerBound; n >= upperBound; n-- {
				res = append(res, matrix[n][leftBound])
			}

			//左边界右移
			leftBound++
		}

	}

	return res
}

func main() {
	fmt.Println("1. 矩阵的顺时针/逆时针旋转90°：")
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("旋转之前的矩阵为：", matrix)
	matrix = rotateMatrix(matrix)
	fmt.Println("旋转之后的矩阵为：", matrix)

	line.SplitLine()

	fmt.Println("2. 矩阵的螺旋遍历输出：")
	matrix1 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("原矩阵为：", matrix1)
	res := spiralMatrix(matrix1)
	fmt.Println("遍历后为：", res)
}
