package main

import (
	"basic-algorithm/labuladuo/define/link"
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 1. 寻找单链表的倒数第k个节点
func searchDescKNodeFromLink(k int, link *link.Node) *link.Node {
	fast, slow := link, link

	count := 0
	for fast != nil && count < k {
		fast = fast.Next
		count++
	}

	for fast != nil && slow != nil {
		fast, slow = fast.Next, slow.Next
	}

	return slow
}

// 2. 寻找单链表的中点 - 快慢指针
func searchMiddleNodeFromLink(link *link.Node) *link.Node {
	fast, slow := link, link

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// 3. 判断单链表是否有环并找出环起点 - 从相遇点在走k-m下一个相遇点就是环起点
func searchCircleNodeFromLink(link *link.Node) *link.Node {
	fast, slow := link, link

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = link
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}

	return slow
}

// 4. 判断两个单链表是否相交并找出交点 - 交点意味着相等
func searchCommonNodeFromLink(link1, link2 *link.Node) *link.Node {
	if link1 == nil || link2 == nil {
		return nil
	}

	p1, p2 := link1, link2
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = link2
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = link1
		}
	}

	return p1
}

func main() {
	// 1. 查找链表的倒数第k个节点
	link1 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7, Next: &link.Node{Val: 9}}}}}
	fmt.Println("1. 查找链表的倒数第k个节点：")
	link.EchoLink(link1, "原链表为：")
	nodeFromLink := searchDescKNodeFromLink(4, link1)
	link.EchoLink(nodeFromLink, "查找的链表结果为：")

	line.SplitLine()

	// 2. 寻找链表的中点
	link2 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7, Next: &link.Node{Val: 9, Next: &link.Node{Val: 11}}}}}}
	fmt.Println("2. 寻找链表的中点：")
	link.EchoLink(link2, "原链表为：")
	middleNodeFromLink := searchMiddleNodeFromLink(link2)
	link.EchoLink(middleNodeFromLink, "查找之后的链表结果为：")

	line.SplitLine()

	// 3. 找出单链表是否有环，并输出环起点
	node1 := &link.Node{Val: 1}
	node2 := &link.Node{Val: 2}
	node3 := &link.Node{Val: 3}
	node4 := &link.Node{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node3
	fmt.Println("3. 找出单链表是否有环，并输出环起点：")
	fmt.Println("原链表为： node1 -> node2 -> node3 -> <-node4")
	circleNodeFromLink := searchCircleNodeFromLink(node1)
	fmt.Printf("链表相遇的环起点为：\n%d\n", circleNodeFromLink.Val)

	line.SplitLine()

	// 4. 判断两个链表是否相交，并找出交点
	node11 := &link.Node{Val: 1}
	node22 := &link.Node{Val: 2}
	node33 := &link.Node{Val: 3}
	node44 := &link.Node{Val: 4}
	node55 := &link.Node{Val: 4}
	node11.Next = node22
	node22.Next = node33
	node33.Next = node44
	node55.Next = node33
	fmt.Println("4. 判断两个链表是否相交，并找出交点：")
	fmt.Println("原链表1为：node11 -> node22 -> node33")
	fmt.Println("原链表2为：node55 -> node33")
	commonNodeFromLink := searchCommonNodeFromLink(node11, node55)
	link.EchoLink(commonNodeFromLink, "两链表的交点为：")
}
