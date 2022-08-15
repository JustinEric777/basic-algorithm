package main

type Node struct {
	Key int
	Val int
}

type DoubleLinkedList struct {
	Prev *Node
	Next *Node
}

var HashMap = make(map[int]*Node)

//在链表头部添加节点
func (dll *DoubleLinkedList) AddFirst(node *Node) {

}

//删除链表中的节点
func (dll *DoubleLinkedList) RemoveNode(node *Node) {

}

//删除链表中的最后一个节点
func (dll *DoubleLinkedList) RemoveLast() int {

	return 1
}

//返回链表的长度
func (dll *DoubleLinkedList) Size(node *Node) {

}

//get
func Get(key int) int {
	if val, ok := HashMap[key]; ok {
		return val.Val
	} else {
		return -1
	}
}

//put
func Put(key, val, len1 int) {
	list := DoubleLinkedList{}
	node := &Node{Key: key, Val: val}
	if val, ok := HashMap[key]; ok {
		list.RemoveNode(val)
		list.AddFirst(node)

		//更新map中的数据
		temp := HashMap[1]
		HashMap[1] = node
		HashMap[2] = temp
	} else {
		if len1 == len(HashMap) {
			last := list.RemoveLast()

			//更新map中的顺序
			delete(HashMap, last)
		}

		list.AddFirst(node)
		MapSort(node)
	}
}

func MapSort(node *Node) {
	for k, v := range HashMap {
		HashMap[k+1] = v
	}

	HashMap[1] = node
}
