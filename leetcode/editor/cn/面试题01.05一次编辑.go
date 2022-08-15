//字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，
//编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
//
//
// 示例 1:
//
// 输入:
//first = "pale"
//second = "ple"
//输出: True
//
//
//
// 示例 2:
//
// 输入:
//first = "pales"
//second = "pal"
//输出: False
//
// Related Topics 双指针 字符串 👍 106 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import "fmt"

func oneEditAway(first string, second string) bool {
	var temp string
	if len(first) < len(second) {
		temp, first = first, second
		second = temp
	}

	if len(first)-len(second) > 1 {
		return false
	}

	if len(first) == len(second) {
		sum := 0
		for i := 0; i < len(second); i++ {
			if first[i] != second[i] {
				sum++
			}
		}
		return sum <= 1
	} else {
		i, j := 0, len(second)-1
		for i <= j && first[i] == second[i] {
			i++
		}
		for i <= j && first[j+1] == second[j] {
			j--
		}

		return i > j
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	first := "pale"
	second := "ple"
	res := oneEditAway(first, second)
	fmt.Println(res)
}
