package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
	"math"
)

// 1. 最小覆盖子串 - 返回字符串 s 中 涵盖t所有子符的最小子串
func minCoverStr(s, t string) string {
	if len(s) < len(t) {
		return ""
	}
	windows := make(map[string]int)
	needle := make(map[string]int, len(t))
	for _, val := range t {
		needle[string(val)]++
	}

	left, right, need := 0, 0, 0
	start, length := 0, len(s)
	for right < len(s) {
		// 向右扩张囊括所有t
		iword := string(s[right])
		right++
		if _, res := needle[iword]; res {
			windows[iword]++
			if windows[iword] == needle[iword] {
				need++
			}
		}

		// 左边缩小窗口
		for need == len(needle) {
			// 更新最小覆盖串,找到长度最小的
			if right-left < length {
				start = left
				length = right - left
			}

			dword := string(s[left])
			left++
			if _, res := needle[dword]; res {
				if windows[dword] == needle[dword] {
					need--
				}
				windows[dword]--
			}
		}
	}

	return s[start : start+length]
}

// 2. 字符串排列 - 两个字符串 s1 和 s2, 判断 s2 是否包含 s1 的 排列
func arrangeStr(s1, s2 string) bool {
	windows := make(map[string]int)
	needle := make(map[string]int, len(s1))
	for _, val := range s1 {
		needle[string(val)]++
	}

	left, right, need := 0, 0, 0
	for right < len(s2) {
		//窗口向右扩展
		iword := string(s2[right])
		right++
		if _, res := needle[iword]; res {
			windows[iword]++
			if windows[iword] == needle[iword] {
				need++
			}
		}

		//窗口从左边开始收缩
		for need == len(needle) {
			if right-left == need {
				return true
			}
			oword := string(s2[left])
			left++
			if _, res := needle[oword]; res {
				if windows[oword] == needle[oword] {
					need--
				}
				windows[oword]--
			}
		}

	}

	return false
}

// 3. 找所有字母的异位词 - 两个字符串 s 和 p, 找到 s 中的所有 p 的异位词，并返回对应的索引， cba, bac 是 abc 的异位词
func findAnagramsFromStr(s, p string) []int {
	needle := make(map[string]int, len(p))
	windows := make(map[string]int)
	for _, val := range p {
		needle[string(val)]++
	}

	var resArr []int
	left, right, need := 0, 0, 0
	for right < len(s) {
		// 窗口向右扩展
		iword := string(s[right])
		right++
		if _, res := needle[iword]; res {
			windows[iword]++
			if windows[iword] == needle[iword] {
				need++
			}
		}

		// 左边收缩窗口
		for need == len(needle) {
			// 记录满足条件的异位词
			if right-left == len(needle) && s[left:right] != p {
				resArr = append(resArr, left)
			}

			oword := string(s[left])
			left++
			if _, res := needle[oword]; res {
				if windows[oword] == needle[oword] {
					need--
				}
				windows[oword]--
			}
		}
	}

	return resArr
}

// 4. 最长无重复子串
func maxLengthNoRepeatStr(s string) string {
	windows := make(map[string]int)
	left, right, start, length := 0, 0, 0, 0

	for right < len(s) {
		// 窗口右移扩充
		iword := string(s[right])
		right++
		windows[iword]++

		// 收缩窗口
		for windows[iword] > 1 {
			oword := string(s[left])
			left++
			windows[oword]--
		}

		if right-left > length {
			start = left
			length = right - left
		}
	}

	return s[start : start+length]
}

// 5. 寻找重复子序列 - DNA 序列由四种碱基 A, G, C, T 组成，现在给你输入一个只包含 A, G, C, T 四种字符的字符串 s 代表一个 DNA 序列，请你在 s 中找出所有重复出现的长度为 10 的子字符串
func findRepeatSubSequence(s string) []string {
	windows := make(map[string]int)
	left, right := 0, 0
	str := ""

	var result []string
	for right < len(s) {
		// 向右扩充
		iword := string(s[right])
		right++
		str += iword

		// 左边收缩窗口
		if right-left == 10 {
			if _, res := windows[str]; res {
				result = append(result, str)
			}

			left++
			windows[str]++
			// 移除对应的字符
			str = str[1:]
		}

	}

	return result
}

// 活动 hash 技巧 - 4 进制处理
func findRepeatSubSequence1(s string) []string {
	// 字符串转换为 4 进制数组
	nums := make([]int, len(s))
	for key, val := range s {
		switch string(val) {
		case "A":
			nums[key] = 0
		case "G":
			nums[key] = 1
		case "C":
			nums[key] = 2
		case "T":
			nums[key] = 3
		}
	}

	// 存储重复的 hash 结果
	seen := make(map[int]int)
	// 存储重复的 substr 结果
	var result []string

	// hash
	L, R, windowHash := 10, 4, 0
	RL := int(math.Pow(float64(R), float64(L-1)))

	// slide window
	left, right := 0, 0
	for right < len(nums) {
		// 向右扩大窗口
		windowHash = windowHash*R + nums[right]
		right++

		// 收缩窗口
		for right-left == L {
			if _, res := seen[windowHash]; res {
				result = append(result, s[left:right])
			} else {
				seen[windowHash]++
			}

			windowHash = windowHash - nums[left]*RL
			left++
		}
	}
	return result
}

// 6. Rabin-Karp 算法
func rabinKarpStr() {

}

// 7. kmp算法, 字符串匹配
func kmpStr() {

}

func main() {
	fmt.Println("1. 最小覆盖子串：")
	str, t := "ADOBECODEBANC", "ABC"
	fmt.Println("原字符串 s 和 t 分别为：", str, ",", t)
	coverStr := minCoverStr(str, t)
	fmt.Println("查找的结果为：", coverStr)

	line.SplitLine()

	fmt.Println("2. 字符串的排列（s2是否包含s1的排列）：")
	s1, s2 := "ab", "eidabooo"
	fmt.Println("原字符串 s1 和 s2 分别为：", s1, "，", s2)
	res := arrangeStr(s1, s2)
	fmt.Println("判责的结果为：", res)
	s11, s22 := "ab", "eidboaoo"
	fmt.Println("原字符串 s1 和 s2 分别为：", s11, ",", s22)
	res1 := arrangeStr(s11, s22)
	fmt.Println("判责的结果为：", res1)

	line.SplitLine()

	fmt.Println("3. 找所有字母的异位词：")
	s, p := "cbaebabacd", "abc"
	fmt.Println("原字符串 s 和目标字符串 p 为：", s, ",", p)
	resFromStr := findAnagramsFromStr(s, p)
	fmt.Println("查找异位词的结果为：", resFromStr)

	line.SplitLine()

	fmt.Println("4. 最长无重复子串：")
	str1 := "abcabcbb"
	fmt.Println("原字符串为：", str1)
	noRepeatMaxStr := maxLengthNoRepeatStr(str1)
	fmt.Println("查找最长的字符串结果为：", noRepeatMaxStr)

	line.SplitLine()

	fmt.Println("5. 高效寻找重复子序列：- 方法一（slide window）")
	str2 := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println("原字符串为：", str2)
	repeatSubSequence := findRepeatSubSequence(str2)
	fmt.Println("查找的结果为：", repeatSubSequence)

	fmt.Println("5. 高效寻找重复子序列：- 方法二 （hash slide window）")
	str3 := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println("原字符串为：", str3)
	repeatSubSequence1 := findRepeatSubSequence1(str3)
	fmt.Println("查找的结果为：", repeatSubSequence1)

	line.SplitLine()
}
