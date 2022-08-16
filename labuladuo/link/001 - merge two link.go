package main

import "fmt"

type linkNode struct {
	Val  int       `json:"val"`
	Next *linkNode `json:"next"`
}

func echoLink(link *linkNode) {
	fmt.Println("链表的输出结果为：")
	for link != nil {
		fmt.Printf("%d\t", link.Val)
		link = link.Next
	}
	fmt.Printf("\n")
}

// 合并两个有序链表
func mergeTwoSortedLink(link1, link2 *linkNode) *linkNode {
	var head, tail *linkNode

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
func mergeKSortedLink(links ...*linkNode) *linkNode {
	//init
	head := links[0]

	for i := 1; i < len(links); i++ {
		head = mergeTwoSortedLink(head, links[i])
	}

	return head
}

func main() {
	// 1. 合并两个有序链表
	link1 := &linkNode{1, &linkNode{3, &linkNode{9, nil}}}
	link2 := &linkNode{5, &linkNode{7, &linkNode{11, nil}}}
	link := mergeTwoSortedLink(link1, link2)
	echoLink(link)

	// 2. 合并k个有序队列
	link3 := &linkNode{2, &linkNode{4, &linkNode{8, nil}}}
	link4 := &linkNode{1, &linkNode{3, &linkNode{9, nil}}}
	link5 := &linkNode{5, &linkNode{7, &linkNode{11, nil}}}
	sortedLink := mergeKSortedLink(link3, link4, link5)
	echoLink(sortedLink)
}
