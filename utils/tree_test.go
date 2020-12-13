package utils

import (
	"fmt"
	"testing"
)

var intTree = NewIntTree()
var stringTree = NewStringTree()

func TestLength(t *testing.T) {
	length := intTree.Length()
	t.Logf("Length of empty int tree: %v\n", length)
	if intTree.Length() != 0 {
		t.Log("Length failed!")
		t.FailNow()
	}
}

func TestAddInt(t *testing.T) {
	intTree.Add(3, 2, 1)
	t.Logf("Add 3,2,1 into int tree (show in order): %v\n", intTree.String())
	if intTree.String() != "[1, 2, 3]" {
		t.Log("Create int tree failed!")
		t.FailNow()
	}
}

func TestAddString(t *testing.T) {
	stringTree.Add("3", "2", "1")
	t.Logf("Add 3,2,1 into string tree (show in order): %v\n", stringTree.String())
	if stringTree.String() != "[1, 2, 3]" {
		t.Log("Create string tree failed!")
		t.FailNow()
	}
}

func TestReverse(t *testing.T) {
	reverseInt := intTree.Reverse()
	reverseString := stringTree.Reverse()
	t.Logf("Reverse int tree equals string tree: %v %v\n", reverseInt, reverseString)
	if reverseInt == reverseString {
		t.Log("Reverse failed!")
		t.FailNow()
	}
	intTree.Reverse()
	stringTree.Reverse()
}

func TestClone(t *testing.T) {
	clone := intTree.Clone()
	t.Logf("Clone int tree: %v %v\n", intTree.String(), clone.String())
	if clone.String() != intTree.String() {
		t.Log("Clone failed!")
		t.FailNow()
	}
}

func TestContainsValue(t *testing.T) {
	check := intTree.ContainsOne(1)
	check2 := stringTree.Contains("1", "2", "3")

	t.Logf("Int tree contains 1: %v", check)
	t.Logf("String tree contains '1' '2' '3': %v %v %v\n", check2[0], check2[1], check2[2])
	if !check || !check2[0] || !check2[1] {
		t.Log("Contains failed!")
		t.FailNow()
	}
}

func TestHeight(t *testing.T) {
	strTreeHeight := stringTree.Height()
	intTreeHeight := intTree.Add(4, 5).Height()
	t.Logf("Height of %v and %v: %v %v\n", stringTree, intTree, strTreeHeight, intTreeHeight)
	if strTreeHeight != intTreeHeight {
		t.Log("Height failed!")
		t.FailNow()
	}
}

func TestPreOrder(t *testing.T) {
	preOrder := intTree.PreOrder()
	t.Logf("Preorder of %v: %v\n", intTree, preOrder)
	if fmt.Sprintf("%v", preOrder) != "[3 2 1 4 5]" {
		t.Log("Height failed!")
		t.FailNow()
	}
}

func TestPosOrder(t *testing.T) {
	posOrder := intTree.PosOrder()
	t.Logf("PosOrder of %v: %v\n", intTree, posOrder)
	if fmt.Sprintf("%v", posOrder) != "[1 2 5 4 3]" {
		t.Log("Height failed!")
		t.FailNow()
	}
}
