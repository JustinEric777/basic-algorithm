//给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
//
// 回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
//
// 回文串不一定是字典当中的单词。
//
//
//
// 示例1：
//
// 输入："tactcoa"
//输出：true（排列有"tacocat"、"atcocta"，等等）
//
//
//
// Related Topics 位运算 哈希表 字符串 👍 67 👎 0
package main

import (
	"fmt"
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func canPermutePalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	hashMap := make(map[string]string)
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if _, ok := hashMap[char]; !ok {
			hashMap[char] = char
		}
	}

	if float64(len(hashMap)) == math.Round(float64(len(s))/2) {
		return true
	}

	return false
}

// bit map
// 基数个数出现次数大于1即为false
func canPermutePalindrome1(s string) bool {
	if len(s) < 2 {
		return true
	}

	hashMap := make(map[string]string)
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if _, ok := hashMap[char]; ok {
			delete(hashMap, char)
		} else {
			hashMap[char] = char
		}
	}

	if len(hashMap) < 2 {
		return true
	} else {
		return false
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	str := "tactcoa"
	res := canPermutePalindrome(str)
	fmt.Println(res)
	res1 := canPermutePalindrome1(str)
	fmt.Println(res1)
}
