package tree

type Interface interface {
	//Len() int
	Less(i, j int) bool
	//Swap(i, j int)
	Insert(r *TreeNode, v interface{}) interface{}
}

type TreeNode struct {
	Value interface{}
	l, r  *TreeNode
}

//func (n *Node) Len() int
