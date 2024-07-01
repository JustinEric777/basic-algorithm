package main

import (
	"basic-algorithm/labuladuo/define/link"
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 1. 判断链表是否是回文链表 - 转置之后判等
func palindromeLink(head *link.Node) bool {
	if head == nil || head.Next == nil {
		return false
	}

	reserveHead := reserveLink(head)
	for head != nil && reserveHead != nil {
		if head.Val != reserveHead.Val {
			return false
		}
		head, reserveHead = head.Next, reserveHead.Next
	}

	return true
}

func reserveLink(head *link.Node) *link.Node {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, cur, next *link.Node
	prev, cur, next = nil, head, head
	for cur != nil {
		next = cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}

// 2. 判断回文链表，根据递归的思想后续遍历
var left *link.Node

func palindromeLink1(head *link.Node) bool {
	left = head
	return reserve(head)
}

func reserve(right *link.Node) bool {
	if right == nil {
		return true
	}

	// 将递归当堆栈使用
	res := reserve(right.Next)
	res = res && (right.Val == left.Val)
	left = left.Next

	return res
}

// 3. 找到链表的中点，然后反转一半，然后比较
func palindromeLink2(head *link.Node) bool {
	// 找中点
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}
	//如果fast指针没有指向null，说明链表长度为奇数，slow还要再前进一步
	if fast != nil {
		slow = slow.Next
	}

	left := head
	right := reserveLink(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next
	}

	return true
}

func main() {
	node6 := &link.Node{Val: 1}
	node5 := &link.Node{Val: 5, Next: node6}
	node4 := &link.Node{Val: 7, Next: node5}
	node3 := &link.Node{Val: 7, Next: node4}
	node2 := &link.Node{Val: 5, Next: node3}
	node1 := &link.Node{Val: 1, Next: node2}
	fmt.Println("1. 判断链表是否为回文链表：")
	link.EchoLink(node1, "原链表为：")
	res := palindromeLink(node1)
	fmt.Println("回文链表的判责结果为：", res)

	line.SplitLine()

	node66 := &link.Node{Val: 1}
	node55 := &link.Node{Val: 5, Next: node66}
	node44 := &link.Node{Val: 7, Next: node55}
	node33 := &link.Node{Val: 7, Next: node44}
	node22 := &link.Node{Val: 5, Next: node33}
	node11 := &link.Node{Val: 1, Next: node22}
	fmt.Println("2. 判断链表是否为回文链表：")
	link.EchoLink(node11, "原链表为：")
	res1 := palindromeLink1(node11)
	fmt.Println("回文链表的判责结果为：", res1)

	line.SplitLine()

	node663 := &link.Node{Val: 1}
	node553 := &link.Node{Val: 5, Next: node663}
	node443 := &link.Node{Val: 7, Next: node553}
	node333 := &link.Node{Val: 7, Next: node443}
	node223 := &link.Node{Val: 5, Next: node333}
	node113 := &link.Node{Val: 1, Next: node223}
	fmt.Println("3. 判断链表是否为回文链表：")
	link.EchoLink(node11, "原链表为：")
	res2 := palindromeLink1(node113)
	fmt.Println("回文链表的判责结果为：", res2)
}
