package tree

type Interface interface {
	//Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Node struct {
	Val  interface{}
	l, r *node
}

//func (n *Node) Len() int
