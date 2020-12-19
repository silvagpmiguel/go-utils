package utils

import (
	"testing"
)

var intQueue = NewIntPQueue(true)
var stringQueue = NewStringPQueue(false)

func TestEmptyQueue(t *testing.T) {
	length := intQueue.Length()
	t.Logf("Length of empty int queue: %v\n", length)

	if !intQueue.IsEmpty() {
		t.Log("Length failed!")
		t.FailNow()
	}
}

func TestEnqueueIntAsc(t *testing.T) {
	intQueue.Enqueue(3, 2, 1)
	t.Logf("Add 3,2,1 into int queue: %v\n", intQueue.String())

	if intQueue.String() != "[1, 3, 2]" {
		t.Log("Create int queue failed!")
		t.FailNow()
	}
}

func TestEnqueueStringDesc(t *testing.T) {
	stringQueue.Enqueue("1", "2", "3")
	t.Logf("Add 3,2,1 into string queue: %v\n", stringQueue.String())

	if stringQueue.String() != "[3, 1, 2]" {
		t.Log("Create string queue failed!")
		t.FailNow()
	}
}

func TestQueueLength(t *testing.T) {
	length := intQueue.Length()
	t.Logf("Int Queue %v length: %v\n", intQueue, length)

	if length != 3 {
		t.Log("Queue length failed!")
		t.FailNow()
	}
}

func TestQueueClone(t *testing.T) {
	clone := intQueue.Clone()
	t.Logf("Clone int tree %v: %v\n", intQueue.String(), clone.String())

	if clone.String() != intQueue.String() {
		t.Log("Clone queue failed!")
		t.FailNow()
	}
}

func TestPeek(t *testing.T) {
	val, _ := intQueue.Peek()
	t.Logf("Peek Queue: %v", val)

	if val.(int) != 1 {
		t.Log("Peek queue failed!")
		t.FailNow()
	}
}

func TestDequeue(t *testing.T) {
	original := intQueue.String()
	_, err := intQueue.Dequeue()
	t.Logf("Dequeue %v: %v\n", original, intQueue)

	t.Logf("Error: %v", err)
	if intQueue.String() == "[2,3]" {
		t.Log("Dequeue failed!")
		t.FailNow()
	}
}
