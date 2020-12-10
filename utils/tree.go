package utils

import (
	"fmt"
)

// Node of a tree
type Node struct {
	left  *Node
	right *Node
	val   interface{}
}

// Tree struct
type Tree struct {
	comparator Comparator
	node       *Node
	length     int
}

// NewIntTree creates a new int tree
func NewIntTree(asc bool) *Tree {
	comp := intDescComparator

	if asc {
		comp = intAscComparator
	}

	return &Tree{comparator: comp}
}

// NewStringTree creates a new string tree
func NewStringTree(asc bool) *Tree {
	comp := stringDescComparator

	if asc {
		comp = stringAscComparator
	}

	return &Tree{comparator: comp}
}

// NewTree creates a new tree ordered by a user defined comparator
func NewTree(comp Comparator) *Tree {
	return &Tree{comparator: comp}
}

// Add a value to the tree
func (t *Tree) Add(values ...interface{}) *Tree {
	for _, val := range values {
		t.node = t.node.add(t.comparator, val)
		t.length++
	}
	return t
}

// Search a value of the tree
func (t *Tree) Search(value interface{}) *Node {
	return t.node.search(t.comparator, value)
}

/** Remove TODO
func (t *Tree) Remove(value interface{}) *Tree {
	return t
}*/

// Reverse tree
func (t *Tree) Reverse() *Tree {
	t.node.reverse()
	return t
}

// Height returns the height of the tree
func (t *Tree) Height() int {
	return t.node.height()
}

// Clone makes a copy of a tree
func (t *Tree) Clone() *Tree {
	cloned := &Tree{comparator: t.comparator}
	t.node.clone(t.comparator, cloned)
	return cloned
}

// Length tree
func (t *Tree) Length(value interface{}) int {
	return t.length
}

// InOrder returns a array of nodes in order
func (t *Tree) InOrder() []*Node {
	arr := []*Node{}
	t.node.inOrder(&arr)
	return arr
}

// PreOrder returns a array of nodes in pre order
func (t *Tree) PreOrder() []*Node {
	arr := []*Node{}
	t.node.preOrder(&arr)
	return arr
}

// PosOrder returns a array of nodes in pos order
func (t *Tree) PosOrder() []*Node {
	arr := []*Node{}
	t.node.posOrder(&arr)
	return arr
}

// String returns the string method of this type
func (t *Tree) String() string {
	if t.length == 0 {
		return "[]"
	}

	nodes := t.InOrder()
	str := nodes[0].String()

	for i := 1; i < len(nodes); i++ {
		str += ", " + nodes[i].String()
	}

	return "[" + str + "]"
}

// String returns the string method of this type
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.val)
}

/** Private Aux Methods **/

func (n *Node) add(comp Comparator, val interface{}) *Node {
	if n == nil {
		return &Node{val: val}
	}

	if comp(n.val, val) > 0 {
		n.left = n.left.add(comp, val)
	} else {
		n.right = n.right.add(comp, val)
	}

	return n
}

func (n *Node) reverse() {
	if n == nil {
		return
	}

	temp := n.right
	n.right = n.left
	n.left = temp
	n.left.reverse()
	n.right.reverse()
}

func (n *Node) clone(comp Comparator, t *Tree) {
	if n == nil {
		return
	}

	n.left.clone(comp, t)
	t.node = t.node.add(comp, n.val)
	n.right.clone(comp, t)
}

func (n *Node) search(comp Comparator, val interface{}) *Node {
	if comp(n.val, val) == 0 {
		return n
	} else if comp(n.val, val) > 0 {
		return n.left.search(comp, val)
	} else {
		return n.right.search(comp, val)
	}
}

func (n *Node) height() int {
	if n == nil {
		return 0
	}

	lheight := n.left.height()
	height := n.right.height()

	if lheight > height {
		height = lheight
	}

	return lheight + 1
}

func (n *Node) inOrder(arr *[]*Node) {
	if n == nil {
		return
	}

	n.left.inOrder(arr)
	*arr = append(*arr, n)
	n.right.inOrder(arr)
}

func (n *Node) preOrder(arr *[]*Node) {
	if n == nil {
		return
	}

	*arr = append(*arr, n)
	n.left.preOrder(arr)
	n.right.preOrder(arr)
}

func (n *Node) posOrder(arr *[]*Node) {
	if n == nil {
		return
	}

	n.left.posOrder(arr)
	n.right.posOrder(arr)
	*arr = append(*arr, n)
}
