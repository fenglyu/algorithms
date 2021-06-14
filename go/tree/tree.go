package tree

type NodeInterface interface {
	//Len() int
	//Less(i,j int) bool
	Less(v interface{}) bool
	//Compare(i, j interface{}) int
	//Swap(i, j int)
	//Insert(r *BinaryNode, v interface{}) interface{}
}

//func (n *Node) Len() int

type BinaryTree struct {
	root *BinaryNode
}

// Insert value into proper location in Binary Tree
func (t *BinaryTree) Add(value interface{}) {
	if t.root == nil {
		t.root = &BinaryNode{Value: value}
	} else {
		t.root = t.root.Add(value)
	}
}

type BinaryNode struct {
	Value  interface{}
	l, r   *BinaryNode
	height int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (n *BinaryNode) Less(v interface{}) bool {
	switch n.Value.(type) {
	case int:
		return n.Value.(int) < v.(int)
	case float64:
		return n.Value.(float64) < v.(float64)
	}
	return false
}

// NewNode create new binary node
func NewNode(value interface{}) *BinaryNode {
	return &BinaryNode{
		Value:  value,
		l:      nil,
		r:      nil,
		height: 0,
	}
}

// computeHeight compute height of node in BST form children
func (n *BinaryNode) computeHeight() {
	height := 1
	if n.l != nil {
		height = max(height, n.l.height)
	}

	if n.r != nil {
		height = max(height, n.r.height)
	}

	n.height = height + 1
}

// heightDifference compute height difference of node's children in BST
func (n *BinaryNode) heightDifference() int {

	leftTarget, rightTarget := 0, 0
	if n.l != nil {
		leftTarget = 1 + n.l.height
	}

	if n.r != nil {
		rightTarget = 1 + n.r.height
	}

	return leftTarget - rightTarget
}

func (n *BinaryNode) Add(value interface{}) *BinaryNode {
	newRoot := n
	// if vale <= n.Value
	if !n.Less(value) {
		n.l = n.addToSubTree(n.l, value)
		if n.heightDifference() == 2 {
			if n.l.Less(val) {
				newRoot = n.rotateLeftRight()
			} else {
				newRoot = n.rotateRight()
			}
		}
	} else {
		n.r = n.addToSubTree(n.r, value)
		if n.heightDifference() == 02 {
			if n.r.Less(value) {
				newRoot = n.rotateLeft()
			} else {
				newRoot = n.rotateLeftRight()
			}
		}
	}
	newRoot.computeHeight()
	return newRoot
}

// Add value to parent subtree(if exists) and return root in case it has changed because of rotation
func (n *BinaryNode) addToSubTree(parent *BinaryNode, value interface{}) *BinaryNode {
	if parent == nil {
		return NewNode(value)
	}

	parent = parent.Add(value)
	return parent
}

// Perform left rotation on given node
func (n *BinaryNode) rotateLeft() *BinaryNode {
	newRoot := n.l
	grandson := newRoot.r
	n.l = grandson
	newRoot.r = n
	n.computeHeight()
	return newRoot
}

// Perform right rotation on given node
func (n *BinaryNode) rotateRight() *BinaryNode {
	newRoot := n.r
	grandson := newRoot.l
	n.r = grandson
	newRoot.l = n
	n.computeHeight()
	return newRoot

}

// Perform left-rigth rotation on given node

func (n *BinaryNode) rotateLeftRight() *BinaryNode {
	return nil

}

// Perform  rigth-left rotation on given node
func (n *BinaryNode) rotateRightLeft() *BinaryNode {
	return nil

}
