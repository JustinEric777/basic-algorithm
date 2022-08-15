package main

import (
	"fmt"
)

/**
 * 1. 合并两个有序链表
 * 2. 合并k个有序链表
 * 3. 寻找单链表的倒数第k个节点
 * 4. 寻找单链表的中点
 * 5. 判断单链表是否有环并找出环起点
 * 6. 判断两个单链表是否相交并找出交点
 * 7. 删除链表的倒数第k个节点
 * 8. 链表转置
 * 9. k个节点起链表转置
 * 10. k个节点一起翻转链表
 * 11. 判断回文单链表
 * 12. 单链表去重
 */

type LinkNode struct {
	Val  int
	Next *LinkNode
}

/**
 * 1. 合并两个链表
 */
func mergeTwoSingleSortedLink(link1 *LinkNode, link2 *LinkNode) *LinkNode {
	var head, tail *LinkNode

	if link1 == nil {
		return link2
	}
	if link2 == nil {
		return link1
	}

	//确认头尾节点的初始化
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

		//平移尾节点
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

/**
 * 2. 合并k个有序链表
 * 最小堆的方法处理,根据最小堆排序
 */
func mergeKSingleSortedLink(links ...*LinkNode) *LinkNode {
	if len(links) == 0 {
		return &LinkNode{}
	}

	res := links[0]
	for i := 1; i < len(links); i++ {
		res = mergeTwoSingleSortedLink(res, links[i])
	}

	return res
}

/**
 * 3. 寻找单链表的倒数第k个节点
 */
func searchDescKNodeFromSingleLink(link *LinkNode, k int) *LinkNode {
	if link == nil {
		return nil
	}

	//初始化头尾节点
	var count int
	head, tail := link, link
	for tail != nil {
		count++
		tail = tail.Next
		if count == k {
			break
		}
	}

	for tail != nil {
		tail = tail.Next
		head = head.Next
	}

	return head
}

/**
 * 4. 寻找单链表的中点 - 快慢指针
 */
func searchMiddleNodeFromSingleLink(link *LinkNode) *LinkNode {
	if link == nil {
		return nil
	}

	fast, slow := link, link
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

/**
 * 5. 判断单链表是否有环，并寻找环的起点
 */
func searchCircleFirstNodeFromSingleLink(link *LinkNode) *LinkNode {
	if link == nil {
		return nil
	}

	fast, slow := link, link
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}

	//重新指向头结点
	slow = link
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

/**
 * 6. 判断两个单链表是否相交并找出交点
 */
func searchCommonNodeFromTwoSingleLink(link1 *LinkNode, link2 *LinkNode) *LinkNode {
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

/**
 * 7. 删除链表的倒数第k个节点
 */
func delDescKNodeFromSingleLink(link *LinkNode, k int) *LinkNode {
	if link == nil {
		return nil
	}

	var count int
	head, tail := link, link
	for tail != nil {
		count++
		tail = tail.Next
		if count == k {
			break
		}
	}

	prev := head
	for tail != nil {
		head = head.Next
		tail = tail.Next
	}
	head.Next = head.Next.Next

	return prev
}

/**
 * 8. 单链表转置
 */
func reverseFromSingleLink(link *LinkNode, isDiGui bool) *LinkNode {
	if link.Next == nil {
		return link
	}

	// 递归
	if isDiGui {
		newHead := reverseFromSingleLink(link.Next, true)
		link.Next.Next = link
		link.Next = nil
		return newHead
	}

	// 非递归
	var prev, next *LinkNode
	for link != nil {
		//暂存next避免翻转丢失然后翻转
		next = link.Next
		link.Next = prev

		//下一个需要翻转的接待赋值
		prev = link
		link = next
	}

	return prev
}

/**
 * 9. 从倒数第几个k开始翻转单链表
 */
func reverseDescKFromSingleLink(link *LinkNode, k int) *LinkNode {
	if link == nil {
		return nil
	}

	count := 0
	//快慢指针找到k节点
	fast, slow, prev := link, link, link
	for fast != nil && slow != nil {
		fast = fast.Next
		if count < k {
			count++
		} else {
			//暂存所有node- 非递归 - 入栈

			slow = slow.Next
		}
	}

	// 出栈

	//反转
	var prev1 *LinkNode
	prev1 = nil
	for slow != nil {
		next := slow.Next
		slow.Next = prev1

		prev1 = slow
		slow = next
	}

	prev.Next = prev1

	return prev
}

/**
 * 10. k个一组翻转单链表
 */
func reverseUnionKFromSingleLink(link *LinkNode, k int) *LinkNode {

	return nil
}

/**
 * 11. 判断链表是否是回文
 */
func judgeSingleLinkIsHuiWen(link *LinkNode) bool {

	return false
}

/**
 * 12. 单链表去重
 */
func printNoRepeatSingleLink(link *LinkNode) *LinkNode {

	return nil
}

func printLink(link *LinkNode) {
	for link != nil {
		fmt.Print(link.Val)
		link = link.Next
	}

	fmt.Printf("\n")
}

func main() {
	// 1. 合并两个有序链表
	fmt.Println(" 1. 合并两个有序链表:")
	link1 := &LinkNode{1, &LinkNode{2, &LinkNode{4, nil}}}
	link2 := &LinkNode{1, &LinkNode{3, &LinkNode{4, nil}}}
	linkTwo := mergeTwoSingleSortedLink(link1, link2)
	fmt.Printf("    link1 ：")
	printLink(link1)
	fmt.Printf("    link2 ：")
	printLink(link2)
	fmt.Printf("    执行结果是：")
	printLink(linkTwo)

	// 2. 合并k个有序链表
	fmt.Println(" 2. 合并k个有序链表:")
	link3 := &LinkNode{1, &LinkNode{2, &LinkNode{4, nil}}}
	link4 := &LinkNode{1, &LinkNode{3, &LinkNode{4, nil}}}
	link5 := &LinkNode{1, &LinkNode{3, &LinkNode{5, nil}}}
	linkK := mergeKSingleSortedLink(link3, link4, link5)
	fmt.Printf("    link1 ：")
	printLink(link3)
	fmt.Printf("    link2 ：")
	printLink(link4)
	fmt.Printf("    link3 ：")
	printLink(link5)
	fmt.Printf("    执行结果是：")
	printLink(linkK)

	// 3. 查找的单链表的倒数k个节点
	fmt.Println("3. 查找的单链表的倒数k个节点:")
	link6 := &LinkNode{1, &LinkNode{2, &LinkNode{4, &LinkNode{5, nil}}}}
	link := searchDescKNodeFromSingleLink(link6, 4)
	fmt.Println("    查找的结果是：", link)

	// 4. 查找的单链表的倒数k个节点
	fmt.Println("4. 查找的单链表的中点节点:")
	node := searchMiddleNodeFromSingleLink(link3)
	fmt.Println("   查找的中点节点是：", node)

	// 5. 查找的单链表环的起点
	node1 := &LinkNode{1, nil}
	node2 := &LinkNode{2, nil}
	node3 := &LinkNode{3, nil}
	node4 := &LinkNode{4, nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node3
	fmt.Println("5. 查找的单链表环的起点:")
	firstNode := searchCircleFirstNodeFromSingleLink(node1)
	fmt.Println("   查找的起点是：", firstNode)

	node11 := &LinkNode{1, nil}
	node22 := &LinkNode{2, nil}
	node33 := &LinkNode{3, nil}
	node44 := &LinkNode{4, nil}
	node55 := &LinkNode{4, nil}
	node11.Next = node22
	node22.Next = node33
	node33.Next = node44
	node55.Next = node33
	fmt.Println("6. 查找的单链表的交点:")
	commonNode := searchCommonNodeFromTwoSingleLink(node11, node55)
	fmt.Println("   查找的交点是：", commonNode)

	fmt.Println("7. 删除链表的倒数第k个节点:")
	link7 := &LinkNode{1, &LinkNode{2, &LinkNode{4, &LinkNode{5, nil}}}}
	newLink := delDescKNodeFromSingleLink(link7, 2)
	fmt.Print("   删除后的链表为：")
	printLink(newLink)

	fmt.Println("8. 链表转置:")
	link8 := &LinkNode{1, &LinkNode{2, &LinkNode{4, &LinkNode{5, nil}}}}
	fmt.Printf("   转置前的链表为：")
	printLink(link8)
	resLink := reverseFromSingleLink(link8, false)
	fmt.Printf("   转置后的链表为：")
	printLink(resLink)

	fmt.Println("9. 从倒数第k个反转链表:")
	link9 := &LinkNode{1, &LinkNode{2, &LinkNode{3, &LinkNode{4, &LinkNode{5, nil}}}}}
	fmt.Printf("   转置前的链表为：")
	printLink(link9)
	k := 3
	descKFromSingleLink := reverseDescKFromSingleLink(link9, k)
	fmt.Printf("   从倒数第%d转置后的链表为：", k)
	printLink(descKFromSingleLink)
}
