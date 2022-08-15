//给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
//
// 示例 1：
//
// 输入: s1 = "abc",
//输出: true s2 = "bca"
//
//
// 示例 2：
//
// 输入: s1 = "abc", s2 = "bad"
//输出: false
//
//
// 说明：
//
//
// 0 <= len(s1) <= 100
// 0 <= len(s2) <= 100
//
// Related Topics 哈希表 字符串 排序 👍 58 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
// hashMap
func CheckPermutation(s1 string, s2 string) bool {
	hashMap := make(map[string]string)
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		char := string(s1[i])
		hashMap[char] = char
	}

	for j := 0; j < len(s2); j++ {
		if _, ok := hashMap[string(s2[j])]; !ok {
			return false
		}
	}

	return true
}

//bitMap 异或交换律
func CheckPermutation1(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var res1, res2 uint8
	for i := 0; i < len(s1); i++ {
		res1 = s1[i] ^ res1
		res2 = s2[i] ^ res2
	}
	if res1 == res2 {
		return true
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	s1 := "abc"
	s2 := "bca"
	res := CheckPermutation1(s1, s2)
	fmt.Println(res)
}
