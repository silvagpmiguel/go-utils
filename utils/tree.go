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
func NewIntTree() *Tree {
	return &Tree{comparator: intAscComparator}
}

// NewStringTree creates a new string tree
func NewStringTree() *Tree {
	return &Tree{comparator: stringAscComparator}
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

// ContainsOne check if tree contains a specific value
func (t *Tree) ContainsOne(val interface{}) bool {
	return t.node.contains(t.comparator, val)
}

// Contains check if tree contains a list of values
func (t *Tree) Contains(values ...interface{}) []bool {
	arr := make([]bool, len(values))
	comp := t.comparator

	for i, val := range values {
		arr[i] = t.node.contains(comp, val)
	}

	return arr
}

// RemoveMin value
func (t *Tree) RemoveMin() (*Tree, error) {
	if t.length == 0 {
		return nil, fmt.Errorf("tree is empty")
	}

	aux := t.Clone()
	aux.length--
	aux.node.removeMin()
	return aux, nil
}

// RemoveMax value
func (t *Tree) RemoveMax() (*Tree, error) {
	if t.length == 0 {
		return nil, fmt.Errorf("tree is empty")
	}

	aux := t.Clone()
	aux.length--
	aux.node.removeMax()
	return aux, nil
}

// Remove a value from the tree
func (t *Tree) Remove(value interface{}) *Tree {
	t.node.remove(value)
	return t
}

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
	toClone := &Tree{comparator: t.comparator, length: t.length}
	toClone.node = toClone.node.clone(t.node)
	return toClone
}

// IsEmpty tree
func (t *Tree) IsEmpty() bool {
	return t.length == 0
}

// Length tree
func (t *Tree) Length() int {
	return t.length
}

// InOrder returns a array of nodes in order
func (t *Tree) InOrder() []interface{} {
	var arr []interface{}
	t.node.inOrder(&arr)
	return arr
}

// PreOrder returns a array of nodes in pre order
func (t *Tree) PreOrder() []interface{} {
	var arr []interface{}
	t.node.preOrder(&arr)
	return arr
}

// PosOrder returns a array of nodes in pos order
func (t *Tree) PosOrder() []interface{} {
	var arr []interface{}
	t.node.posOrder(&arr)
	return arr
}

// String returns the string method of this type
func (t *Tree) String() string {
	if t.length == 0 {
		return "[]"
	}

	nodes := t.InOrder()
	str := fmt.Sprintf("%v", nodes[0])

	for i := 1; i < len(nodes); i++ {
		str += fmt.Sprintf(", %v", nodes[i])
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

	if comp(val, n.val) > 0 {
		n.right = n.right.add(comp, val)
	} else {
		n.left = n.left.add(comp, val)
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

func (n *Node) contains(comp Comparator, val interface{}) bool {
	if n == nil {
		return false
	}

	flag := false

	if n.val == val {
		return true
	} else if comp(val, n.val) > 0 {
		flag = n.right.contains(comp, val)
	} else {
		flag = n.left.contains(comp, val)
	}

	return flag
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

// TODO
func (n *Node) remove(value interface{}) *Node {
	return n
}

func (n *Node) removeMin() *Node {
	if n.left == nil {
		if n.right == nil {
			return nil
		}
		return n.remove(n.right.val)
	}

	for ; n.left.left != nil; n = n.left {
	}
	n.left = nil

	return n
}

func (n *Node) removeMax() *Node {
	if n.right == nil {
		if n.left == nil {
			return nil
		}
		return n.remove(n.right.val)
	}

	for ; n.right.right != nil; n = n.right {
	}
	n.right = nil

	return n
}

func (n *Node) clone(root *Node) *Node {
	if root == nil {
		return nil
	}

	n = &Node{val: root.val}
	n.left = n.clone(root.left)
	n.right = n.clone(root.right)

	return n
}

func (n *Node) inOrder(arr *[]interface{}) {
	if n == nil {
		return
	}

	n.left.inOrder(arr)
	*arr = append(*arr, n.val)
	n.right.inOrder(arr)
}

func (n *Node) preOrder(arr *[]interface{}) {
	if n == nil {
		return
	}

	*arr = append(*arr, n.val)
	n.left.preOrder(arr)
	n.right.preOrder(arr)
}

func (n *Node) posOrder(arr *[]interface{}) {
	if n == nil {
		return
	}

	n.left.posOrder(arr)
	n.right.posOrder(arr)
	*arr = append(*arr, n.val)
}
