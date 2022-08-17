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

// 3. k个一组转置链表
func reserveLinkGroupByK(k int, head *link.Node) *link.Node {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	return nil
}

func main() {
	// 1. 链表转置
	head1 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7, Next: &link.Node{Val: 9}}}}}
	fmt.Println("1. 链表转置")
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

	// 3，k个一组进行链表转置
	head3 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7, Next: &link.Node{Val: 9, Next: &link.Node{Val: 11}}}}}}
	fmt.Println(head3)
	fmt.Println("3. k个一组进行链表转置：")
	link.EchoLink(head3, "转置前链表为：")
	link.EchoLink(head3, "转置后链表为：")
}
