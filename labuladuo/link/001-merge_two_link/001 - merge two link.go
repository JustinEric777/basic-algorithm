package main

import (
	"basic-algorithm/labuladuo/define/link"
)

// 合并两个有序链表
func mergeTwoSortedLink(link1, link2 *link.Node) *link.Node {
	var head, tail *link.Node

	if link1 == nil {
		return link2
	}
	if link2 == nil {
		return link1
	}

	if link1.Val < link2.Val {
		head, tail, link1 = link1, link1, link1.Next
	} else {
		head, tail, link2 = link2, link2, link2.Next
	}

	for link1 != nil && link2 != nil {
		if link1.Val < link2.Val {
			tail.Next, link1 = link1, link1.Next
		} else {
			tail.Next, link2 = link2, link2.Next
		}
		tail = tail.Next
	}

	if link1 != nil {
		tail.Next = link1
	}
	if link2 != nil {
		tail.Next = link2
	}

	return head
}

// 合并k个有序链表
func mergeKSortedLink(links ...*link.Node) *link.Node {
	//init
	head := links[0]

	for i := 1; i < len(links); i++ {
		head = mergeTwoSortedLink(head, links[i])
	}

	return head
}

func main() {
	// 1. 合并两个有序链表
	link1 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 9}}}
	link2 := &link.Node{Val: 5, Next: &link.Node{Val: 7, Next: &link.Node{Val: 11}}}
	sortedLink1 := mergeTwoSortedLink(link1, link2)
	link.EchoLink(sortedLink1)

	// 2. 合并k个有序队列
	link3 := &link.Node{Val: 2, Next: &link.Node{Val: 4, Next: &link.Node{Val: 8}}}
	link4 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 9}}}
	link5 := &link.Node{Val: 5, Next: &link.Node{Val: 7, Next: &link.Node{Val: 11}}}
	sortedLink2 := mergeKSortedLink(link3, link4, link5)
	link.EchoLink(sortedLink2)
}
