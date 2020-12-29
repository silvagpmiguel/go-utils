package utils

import (
	"fmt"
)

// PQueue structure
type PQueue struct {
	comparator Comparator
	values     []interface{}
	length     int
}

// NewIntPQueue creates a new int priority queue
func NewIntPQueue(min bool) *PQueue {
	comp := intDescComparator

	if min {
		comp = intAscComparator
	}

	return &PQueue{
		comparator: comp,
		length:     0,
	}
}

// NewStringPQueue creates a new string priority queue
func NewStringPQueue(min bool) *PQueue {
	comp := stringDescComparator

	if min {
		comp = stringAscComparator
	}

	return &PQueue{
		comparator: comp,
		length:     0,
	}
}

// NewPQueue creates a new priority queue
func NewPQueue(comp Comparator) *PQueue {
	return &PQueue{
		comparator: comp,
		length:     0,
	}
}

// Enqueue element(s)
func (q *PQueue) Enqueue(val ...interface{}) *PQueue {
	for _, val := range val {
		q.values = append(q.values, val)
		q.bubbleUp(q.length)
		q.length++
	}

	return q
}

// Length of the queue
func (q *PQueue) Length() int {
	return q.length
}

// IsEmpty checks if the queue is empty
func (q *PQueue) IsEmpty() bool {
	return q.length == 0
}

// Peek element to queue
func (q *PQueue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("queue is empty")
	}

	return q.values[0], nil
}

// Dequeue element
func (q *PQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return -1, fmt.Errorf("queue is empty")
	}

	last := q.length - 1
	q.swap(0, last)
	toRemove := q.values[last]
	q.values = q.values[:last]
	q.length--
	q.bubbleDown(0, last)

	return toRemove, nil
}

// Clone queue
func (q *PQueue) Clone() *PQueue {
	return &PQueue{
		comparator: q.comparator,
		values:     q.values,
		length:     q.length,
	}
}

// String method
func (q *PQueue) String() string {
	length := q.length

	if length == 0 {
		return "[]"
	}

	str := fmt.Sprintf("%v", q.values[0])

	for i := 1; i < length; i++ {
		str += fmt.Sprintf(", %v", q.values[i])
	}

	return "[" + str + "]"
}

/* Private Aux Methods */

func (q *PQueue) bubbleDown(start, end int) {
	length := q.length

	for i := start; ; {
		child := q.leftChild(i)

		if child >= length {
			break
		}

		if rightChild := child + 1; rightChild < length && q.needSwap(rightChild, child) {
			child = rightChild
		}

		if !q.needSwap(child, i) {
			break
		}

		q.swap(i, child)
		i = child
	}
}

func (q *PQueue) bubbleUp(actual int) {
	for i := actual; i >= 0; {
		parent := q.parent(i)

		if !q.needSwap(i, parent) {
			break
		}

		q.swap(i, parent)
		i = parent
	}
}

func (q *PQueue) parent(i int) int {
	return (i - 1) / 2
}

func (q *PQueue) rightChild(i int) int {
	return 2*i + 2
}

func (q *PQueue) leftChild(i int) int {
	return 2*i + 1
}

func (q *PQueue) needSwap(i, j int) bool {
	if q.comparator(q.values[i], q.values[j]) < 0 {
		return true
	}
	return false
}

func (q *PQueue) swap(i, j int) {
	aux := q.values[i]
	q.values[i] = q.values[j]
	q.values[j] = aux
}
