package main

import (
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
	"math/rand"
	"time"
)

// O(1) 时间删除，插入，随机获取元素
// 底层维护数组解决随机获取O(1)
type RandomizedSet struct {
	ValToIndex map[int]int
	Nums       []int
}

func NewRandomizedSet() *RandomizedSet {
	valToIndex := make(map[int]int)

	return &RandomizedSet{
		ValToIndex: valToIndex,
	}
}

// 插入 - O (1)
func (randomizedSet *RandomizedSet) insert(num int) bool {
	length := len(randomizedSet.Nums)
	randomizedSet.ValToIndex[num] = length
	randomizedSet.Nums = append(randomizedSet.Nums, num)

	return true
}

// 删除 - O(1) - 移到末尾删除
func (randomizedSet *RandomizedSet) remove(num int) bool {
	length := len(randomizedSet.Nums)
	key := randomizedSet.ValToIndex[num]
	randomizedSet.Nums[key] = randomizedSet.Nums[length-1]
	randomizedSet.Nums = randomizedSet.Nums[:length-1]

	return true
}

func (randomizedSet *RandomizedSet) get() int {
	rand.Seed(time.Now().UnixNano())
	randomNum := int(rand.Int31())
	key := randomNum % len(randomizedSet.Nums)

	return randomizedSet.Nums[key]
}

// 避开黑名单中的随机数 - hashSet
// 输入一个正整数 N,代表左闭右开区间 【0，N), 再输入一个 blacklist 数组（在 0 -> N范围内）,等概率获取随机数且避开黑名单
type BlacklistRandom struct {
	Nums map[int]int
	Sz   int
}

// 主要是把黑名单后移
func NewBlacklistRandom(n int, blacklist []int) *BlacklistRandom {
	nums := make(map[int]int)
	sz := n - len(blacklist)
	for _, val := range blacklist {
		nums[val] = 1
	}

	last := n - 1
	for _, val := range blacklist {
		// 剔除最后一个不是黑名单
		val1 := nums[last]
		for val1 == 1 {
			last--
		}

		if val >= sz {
			continue
		}

		nums[val] = last
		last--
	}

	return &BlacklistRandom{Nums: nums, Sz: sz}
}

// 随机选择一个数
func (blacklistRandom *BlacklistRandom) pick() int {
	rand.Seed(time.Now().UnixNano())
	randomNum := int(rand.Int31())
	key := randomNum % blacklistRandom.Sz

	if blacklistRandom.Nums[key] > 0 {
		return blacklistRandom.Nums[key]
	}

	return key
}

func main() {
	fmt.Println("1. O(1) 复杂度插入、删除、随机获取元素：")
	randomizedSet := NewRandomizedSet()
	randomizedSet.insert(1)
	randomizedSet.insert(2)
	randomizedSet.insert(3)
	randomizedSet.remove(2)
	val := randomizedSet.get()
	fmt.Println("随机获取的元素值为：", val)

	line.SplitLine()

	fmt.Println("2. 避开黑名单生成随机数（黑名单后移）：")
	blacklistRandom := NewBlacklistRandom(12, []int{2, 5, 7})
	pickVal := blacklistRandom.pick()
	fmt.Println("随机生成的数是（避开2，5，7）：", pickVal)
}
