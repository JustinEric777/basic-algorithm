package main

import (
	"basic-algorithm/labuladuo/define/link"
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 1. 链表转置
func reserveLink(head *link.Node) *link.Node {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, cur, next *link.Node
	prev, cur = nil, head
	for cur != nil {
		next = cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}

// 2. 从倒数k个节点开始转置链表
func reserveLinkFromDescK(k int, head *link.Node) *link.Node {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	//search k node
	fast, slow := head, head
	count := 0
	for fast != nil && count < k {
		count++
		fast = fast.Next
	}

	newLink := slow
	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}

	var prev, cur, next *link.Node
	prev, cur = nil, slow.Next
	for cur != nil {
		next = cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}
	slow.Next = prev

	return newLink
}

// 3. 根据区间反转链表 - 区间为左闭右开
func reserveLinkByRangeAB(a, b, head *link.Node) *link.Node {
	// search a
	newLink := head
	for head.Next != a {
		head = head.Next
	}

	var prev, cur, next *link.Node
	prev, cur, next = b, a, a

	for cur != b {
		next = cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	newLink.Next = prev

	return newLink
}

//已a为头结点范围反转
func reserve(a, b *link.Node) *link.Node {
	var prev, cur, next *link.Node
	prev, cur, next = nil, a, a

	for cur != b {
		next = cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}

// 4. k个一组转置链表 - 递归 + 迭代
func reserveLinkGroupByK(k int, head *link.Node) *link.Node {
	if k <= 1 || head == nil || head.Next == nil {
		return head
	}

	a, b := head, head
	count := 0
	for b != nil && count < k {
		count++
		b = b.Next
	}
	newLink := reserve(a, b)
	a.Next = reserveLinkGroupByK(k, b)

	return newLink
}

func main() {
	// 1. 链表转置
	head1 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7, Next: &link.Node{Val: 9}}}}}
	fmt.Println("1. 链表转置：")
	link.EchoLink(head1, "转置前链表为：")
	reservedHead := reserveLink(head1)
	link.EchoLink(reservedHead, "转置后链表为：")

	line.SplitLine()

	// 2. 从倒数第k个节点开始转置
	head2 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7, Next: &link.Node{Val: 9}}}}}
	fmt.Println("2. 链表从倒数第k个节点开始转置：")
	link.EchoLink(head2, "转置前链表为：")
	reserveHeadByDescK := reserveLinkFromDescK(3, head2)
	link.EchoLink(reserveHeadByDescK, "转置后的链表为：")

	line.SplitLine()

	// 3. 根据特定区间反转链表
	node11 := &link.Node{Val: 11}
	node9 := &link.Node{Val: 9, Next: node11}
	node7 := &link.Node{Val: 7, Next: node9}
	node5 := &link.Node{Val: 5, Next: node7}
	node3 := &link.Node{Val: 3, Next: node5}
	node1 := &link.Node{Val: 1, Next: node3}
	fmt.Println("3. 根据区间反转链表：")
	link.EchoLink(node1, "反转前的链表为：")
	linkByRangeAB := reserveLinkByRangeAB(node3, node9, node1)
	link.EchoLink(linkByRangeAB, "反转之后的链表为：")

	line.SplitLine()

	// 4，k个一组进行链表转置
	head4 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7, Next: &link.Node{Val: 9, Next: &link.Node{Val: 11}}}}}}
	fmt.Println("4. k个一组进行链表转置：")
	link.EchoLink(head4, "转置前链表为：")
	linkGroupByK := reserveLinkGroupByK(2, head4)
	link.EchoLink(linkGroupByK, "转置后链表为：")
}
