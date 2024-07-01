package link

import "fmt"

type Node struct {
	Val  int   `json:"val"`
	Next *Node `json:"next"`
}

func EchoLink(link *Node, desc string) {
	if desc == "" {
		desc = "链表的输出结果为："
	}

	fmt.Println(desc)
	if link == nil {
		fmt.Printf("链表内容为空")
	}

	for link != nil {
		fmt.Printf("%d\t", link.Val)
		link = link.Next
	}
	fmt.Printf("\n")
}
