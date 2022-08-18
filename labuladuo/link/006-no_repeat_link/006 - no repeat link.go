package main

import (
	"basic-algorithm/labuladuo/define/link"
	line "basic-algorithm/labuladuo/define/untils"
	"fmt"
)

// 1. 删除排序链表中的重复元素 - 添加虚拟节点，因为头结点也可能被删除
func deleteRepeatNodeFromSortedLink(head *link.Node) *link.Node {
	if head == nil || head.Next == nil {
		return head
	}

	first := &link.Node{Val: -1, Next: head}
	p := first
	for p.Next != nil && p.Next.Next != nil {
		if p.Next.Val == p.Next.Next.Val {
			curVal := p.Next.Val
			for p.Next != nil && p.Next.Val == curVal {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}

	return first.Next
}

// 2. 删除排序链表中的重复元素 - 递归
func deleteRepeatNodeFromSortedLink1(head *link.Node) *link.Node {
	if head == nil || head.Next == nil {
		return head
	}

	if head != nil {
		p := head
		for p.Val == p.Next.Val && p.Next != nil {
			p = p.Next
		}

		//分条件递归调用
		if p == head {
			// 不删除投节点
			head.Next = deleteRepeatNodeFromSortedLink(head.Next)
		} else {
			// 删除头结点
			head = deleteRepeatNodeFromSortedLink(p.Next)
		}
	}

	return head
}

// 3. 排序链表去重
func getNoRepeatNodeForSortedLink(head *link.Node) *link.Node {
	if head == nil || head.Next == nil {
		return head
	}

	first, p := head, head

	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			curVal := p.Val
			for p.Next != nil && p.Next.Val == curVal {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}

	return first
}

func main() {
	fmt.Println("1. 删除有序链表中的重复元素：（单指针）")
	head1 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7}}}}}}
	link.EchoLink(head1, "删除重复元素之前的链表为：")
	deleteRepeatNodeFromSortedLink := deleteRepeatNodeFromSortedLink(head1)
	link.EchoLink(deleteRepeatNodeFromSortedLink, "删除重复元素之后的元素为：")

	line.SplitLine()

	fmt.Println("1. 删除有序链表中的重复元素：（双指针）")

	line.SplitLine()

	fmt.Println("1. 删除有序链表中的重复元素：（递归）")
	head13 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7}}}}}}
	link.EchoLink(head13, "删除重复元素之前的链表为：")
	deleteRepeatNodeFromSortedLink1 := deleteRepeatNodeFromSortedLink1(head13)
	link.EchoLink(deleteRepeatNodeFromSortedLink1, "删除重复元素之后的元素为：")

	line.SplitLine()

	fmt.Println("3. 排序链表去重：(单指针)")
	head2 := &link.Node{Val: 1, Next: &link.Node{Val: 3, Next: &link.Node{Val: 3, Next: &link.Node{Val: 5, Next: &link.Node{Val: 5, Next: &link.Node{Val: 7}}}}}}
	link.EchoLink(head2, "去重之前的链表为：")
	noRepeatSortedLink := getNoRepeatNodeForSortedLink(head2)
	link.EchoLink(noRepeatSortedLink, "去重之后的链表为：")
}
