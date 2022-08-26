package tree

type Node struct {
	Val   int   `json:"val"`
	Left  *Node `json:"left"`
	Right *Node `json:"right"`
}

func EchoTwoBranchTree(root *Node) {
	if root == nil {
		return
	}

}
