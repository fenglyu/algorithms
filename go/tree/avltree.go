package tree

import (
	"fmt"
)

type BinaryTree struct {
	root *BinaryNode
}

// Insert value into proper location in Binary Tree
func (t *BinaryTree) Add(value any) {
	if t.root == nil {
		t.root = &BinaryNode{
			Value:  value,
			l:      nil,
			r:      nil,
			height: 0}
	} else {
		t.root = t.root.Add(value)
	}
}

type BinaryNode struct {
	Value  any
	l, r   *BinaryNode
	height int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// In order traversal of elements in the tree.
func (n *BinaryNode) inorder() {
	if n.l != nil {
		n.l.inorder()
	}
	fmt.Printf("%d ", n.Value)

	if n.r != nil {
		n.r.inorder()
	}
}

func (n *BinaryNode) preorder() {
	fmt.Printf("%d ", n.Value)
	if n.l != nil {
		n.l.preorder()
	}
	if n.r != nil {
		n.r.preorder()
	}
}

func (n *BinaryNode) postorder() {
	if n.l != nil {
		n.l.postorder()
	}
	if n.r != nil {
		n.r.postorder()
	}
	fmt.Printf("%d ", n.Value)
}

func (n *BinaryNode) Less(v any) bool {
	switch n.Value.(type) {
	case int:
		return n.Value.(int) < v.(int)
	case float64:
		return n.Value.(float64) < v.(float64)
	}
	return false
}

// NewNode create new binary node
func NewNode(value any) *BinaryNode {
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

func (n *BinaryNode) Add(value any) *BinaryNode {
	//fmt.Printf("[Add] cur node %d, left %v, right %v\n", n.Value, n.l, n.r)
	newRoot := n
	// if vale <= n.Value
	if !n.Less(value) {
		n.l = n.addToSubTree(n.l, value)
		if n.heightDifference() == 2 {
			if n.l.Less(value) {
				newRoot = n.rotateLeftRight()
			} else {
				newRoot = n.rotateRight()
			}
		}
	} else {
		n.r = n.addToSubTree(n.r, value)
		if n.heightDifference() == -2 {
			if n.r.Less(value) {
				newRoot = n.rotateLeft()
			} else {
				newRoot = n.rotateRightLeft()
			}
		}
	}
	newRoot.computeHeight()
	return newRoot
}

// Add value to parent subtree(if exists) and return root in case it has changed because of rotation
func (n *BinaryNode) addToSubTree(parent *BinaryNode, value any) *BinaryNode {
	//fmt.Printf("[addToSubTree]cur node %d,  left %v, right %v\n", n.Value, n.l, n.r)
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

/*
	Perform left-right rotation on given node
	Addr: shorturl.at/nzGTX
            ┌───┐                                     ┌────┐
            │ 50│                                     │ 40 │\
            └───┘                                   //└────┘\\
           //                                      //        \\
       ┌───/                                   ┌────┐         \\────┐
       │ 30│                                  /│ 30 │          │ 50 │
      /└───┘\             ─────────►         //└────┘        //└────┘
┌───┐//     \\ ┌───┐                        //    \\        //
│30L│        \\│40 │                   ┌───┐/      \────┐ ┌─/──┐
└───┘         /└───\\                  │30L│       │ 40L│ │ 40R│
             //     \\                 └───┘       └────┘ └────┘
         ┌───/     ┌─\─┐
         │40L│     │40R│
         └───┘     └───┘
*/
func (n *BinaryNode) rotateLeftRight() *BinaryNode {
	child := n.l
	newRoot := child.r
	//fmt.Println("rotateLeftRight: ", newRoot, newRoot.l, newRoot.r)
	grand1, grand2 := newRoot.l, newRoot.r

	child.r = grand1
	newRoot.l = child
	newRoot.r = n

	n.l = grand2

	child.computeHeight()
	n.computeHeight()
	return newRoot
}

/*
	Perform right-left rotation on given node
	Addr: shorturl.at/nzGTX
           ┌───┐                                       ┌───┐
           │ 30│                                       │40 │
         //└───┘                                      /└───┘
        //     \\                                    //    \\
       //       \\                                  /       \\
  ┌────┐         \───┐                         ┌───┐         \────┐
  │ 30L│         │ 60│        ───────►        /│30 │         ┼ 60 │
  └────┘       //└───\                       //└───\\        /────┘
              //     \\                     //      \\      //   \\
          ┌───/       \────┐            ┌───/      ┌─\─┐ ┌──/│    \\┌────┐
         /│ 40│       │ 60R│            │30L│      │40L│ │40R│      │60R │
        //└───┘       └────┘            └───┘      └───┘ └───┘      └────┘
       //    \\
    ┌──/│     \┌───┐
    │40L│      │40R│
    └───┘      └───┘
*/
func (n *BinaryNode) rotateRightLeft() *BinaryNode {
	child := n.r
	newRoot := child.l

	grand1, grand2 := newRoot.l, newRoot.r

	child.l = grand2
	newRoot.l = n
	newRoot.r = child
	n.r = grand1

	child.computeHeight()
	n.computeHeight()
	return newRoot
}

// https://www.educative.io/edpresso/how-to-invert-a-binary-tree
/*
Max Howell
@mxcl
Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so fuck off.
*/

func (n *BinaryNode) invertion() *BinaryNode {
	return nil
}

/*
  // https://asciiflow.com/#/share/eJylVUtuwjAQvYrlNZIpX5UzoJ7AG9q43aRBgiCBEBLKuosuoqpnqThNTtJAHH%2FGY4jjkSUbJvGb53l%2BOdJs9SnoItul6YCmq4PY0AU9crrndPE8mw04PdSr0Xxer3Kxz%2BsfnBIVVXkBg%2FPMShtTv6jKr6o8G%2BMbQjwNoyEKOcGtbWZNxMA0pdowTcTuXpWleUxchYNkUDHp9QUm96H0%2FLJOBMnX5FWQRKQiF0lHYoUu1yMCA206jNKClALeJIbJPU4QhEyGKBhjzRypCwZ0oaSh4QAjpo5P0gP58BIAE4BNxub1vS5m4fcZmARr29iqx%2B8gzNKouQguoZBUNLSkYnSXMage81RdfYUVwQz5EmthFiEbEus4zNKWOuXfFhjqrk7VJTh2fmaKO2ZLEYozBIecC%2Bi2G2gzrPHzp5%2F1ssLcS2rwWs1oCqsaT0O15%2B26bDvaC5QkFGRneECjaGk4xnbrTR%2Fp3SXhfPuU3hzsnvA%2BChoC8VHDShBd95C230v99hXazAtw0rPeEfFS204bAY9BFRNp6Eux%2BRDbnKTiPX9YBWRjKboT6URs30SWrLIHYGjZty%2Fz0oHzO%2FTdZLfoyBoEC0DoQsO9ZshVUuAMT2D%2FI7cDfaid8WTTlR5vuoPTEz39A6Y7Dpk%3D)
          ┌────┐
          │    │                                                          ┌────┐
          │10  │                                                          │    │
          └────\\                                                         │10  │
                \\                                                        └────\\
                │\\───┐                                                         \\
                │     │  Node to be deleted                                     │\\───┐
                │  50 │                                                         │     │
               /└─────\                                                         │  40 │
              //      \\                                                       /└─────\
         ┌────/│      │\────┐                                                 //      \\
         │ 30  │      │ 60  │                                            ┌────/│      │\────┐
        /│     │      │     │                                            │ 30  │      │ 60  │
       //└───\─┘      └─────┘                                           /│     │      │     │
      //     \\                                                        //└───\─┘      └─────┘
┌─────/      │\\────┐                                                 //     \\
│     │      │      │                       ───────────►        ┌─────/      │\\────┐
│ 25  │      │  35  │                                           │     │      │      │
└─────┘      └────\─┘                                           │ 25  │      │  35  │
             //   \\                                            └─────┘      └────\─┘
            //      \\                                                       //   \\
       ┌────/─┐     │\─────┐                                                //      \\
       │      │     │      │                                           ┌────/─┐     │\─────┐
       │  33  │     │  40  │ Lergest left                              │      │     │      │
       │      │     │      │ descendant                                │  33  │     │  40L │
       └──────┘     └──────┘                                           │      │     │      │
                      /                                                └──────┘     └──────┘
                    //
                   //
               ┌───/─┐
               │     │
               │ 40L │
               │     │
			   └─────┘

*/
// Remove value from Binary Tree. Works in conjunction with
// remove method in Binary Tree.
func (n *BinaryNode) remove(value any) *BinaryNode {

	newRoot := n

	if value == n.Value {
		// if the target node containing the value to be removed has no left child
		// you can simply "lift" up its right child node
		if n.l == nil {
			return n.r
		}
		// Otherwise, find the node with the largest value in the tree rooted at the left child.
		// We can swap the largest value into the target node.
		child := n.l
		for child.r != nil {
			child = child.r
		}

		childKey := child.Value
		n.l = n.removeFromParent(n.l, childKey)
		n.Value = childKey

		if n.heightDifference() == -2 {
			if n.r.heightDifference() <= 0 {
				newRoot = n.rotateLeft()
			} else {
				newRoot = n.rotateRightLeft()
			}
		}
	} else if n.Less(value) {
		//n.Value < value
		n.r = n.removeFromParent(n.r, value)
		if n.heightDifference() == -2 {
			if n.l.heightDifference() <= 0 {
				newRoot = n.rotateRight()
			} else {
				newRoot = n.rotateLeftRight()
			}
		}
	} else {
		//value < n.Value
		n.l = n.removeFromParent(n.l, value)
		if n.heightDifference() == 2 {
			if n.r.heightDifference() >= 0 {
				newRoot = n.rotateLeft()
			} else {
				newRoot = n.rotateRight()
			}
		}
	}

	newRoot.computeHeight()
	return newRoot
}

// Helper method for remove, Ensures proper behaivor when
// removing node that has children.
func (n *BinaryNode) removeFromParent(parent *BinaryNode, value any) *BinaryNode {
	if parent != nil {
		return parent.remove(value)
	}

	return nil
}
