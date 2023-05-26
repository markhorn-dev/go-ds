package maxheap

import "fmt"

type MaxHeap struct {
	elements []int
}

func New() *MaxHeap {
	return &MaxHeap{}
}

// Time: O(logn)
// Push appends the a new value to the heap, then updates the ordering
func (h *MaxHeap) Push(value int) {
	h.elements = append(h.elements, value)
	h.heapifyUp(len(h.elements) - 1)
}

// Time: O(logn)
// Pop returns the largest value in the heap, then updates the ordering
func (h *MaxHeap) Pop() (int, error) {
	if h.Size() == 0 {
		return 0, fmt.Errorf("cannot pop from an empty maxheap")
	}

	res := h.elements[0]
	last := len(h.elements) - 1
	h.elements[0] = h.elements[last]
	h.elements = h.elements[:last]
	h.heapifyDown(0)

	return res, nil
}

// Time: O(1)
// Peek returns the largest value of the maxheap without popping it
func (h *MaxHeap) Peek() (int, error) {
	if h.Size() == 0 {
		return 0, fmt.Errorf("cannot peek an empty maxheap")
	}
	return h.elements[0], nil
}

// Time: O(1)
// Size returns the number of elements in the heap
func (h *MaxHeap) Size() int {
	return len(h.elements)
}

// Time: O(n)
// Heapify pushes a slice of values to the heap
func (h *MaxHeap) Heapify(values []int) {
	for _, val := range values {
		h.Push(val)
	}
}

// Time: O(logn)
// heapifyUp moves the value up the tree by exchanging the value with its parent node
// until it is in the correct position
func (h *MaxHeap) heapifyUp(idx int) {
	parentIdx := (idx - 1) / 2

	for idx > 0 && h.elements[idx] > h.elements[parentIdx] {
		h.elements[idx], h.elements[parentIdx] = h.elements[parentIdx], h.elements[idx]
		idx = parentIdx
		parentIdx = (idx - 1) / 2
	}
}

// Time: O(logn)
// heapifyUp moves the value up the tree by exchanging the value with its child node
// until it is in the correct position
func (h *MaxHeap) heapifyDown(idx int) {
	size := len(h.elements)
	leftChildIdx := 2*idx + 1
	rightChildIdx := 2*idx + 2
	largest := idx

	if leftChildIdx < size && h.elements[leftChildIdx] > h.elements[largest] {
		largest = leftChildIdx
	}

	if rightChildIdx < size && h.elements[rightChildIdx] > h.elements[largest] {
		largest = rightChildIdx
	}

	if largest != idx {
		h.elements[idx], h.elements[largest] = h.elements[largest], h.elements[idx]
		h.heapifyDown(largest)
	}
}
