package link

import "fmt"

type Node struct {
	Val  int   `json:"val"`
	Next *Node `json:"next"`
}

func EchoLink(link *Node) {
	fmt.Println("链表的输出结果为：")
	for link != nil {
		fmt.Printf("%d\t", link.Val)
		link = link.Next
	}
	fmt.Printf("\n")
}
