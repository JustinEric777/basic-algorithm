package main

import (
	"basic-algorithm/labuladuo/define/link"
	"fmt"
)

// 1. 删除链表的倒数第k个节点
func deleteKNodeFromLinkByDesc(k int, link *link.Node) *link.Node {
	if link == nil {
		return link
	}

	count := 0
	fast, slow := link, link
	for fast != nil {
		count++
		fast = fast.Next
		if count == k {
			break
		}
	}

	//找到慢节点的前节点
	head := slow
	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}
	slow.Next = slow.Next.Next

	return head
}

func main() {
	link1 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7, Next: &link.Node{Val: 9}}}}}
	fmt.Println("1. 删除链表的倒数第k个节点：")
	link.EchoLink(link1, "删除之前的链表结果为：")
	newLink := deleteKNodeFromLinkByDesc(3, link1)
	link.EchoLink(newLink, "删除之后的链表结果为：")
}
