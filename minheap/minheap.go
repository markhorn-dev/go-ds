package minheap

import "fmt"

type MinHeap struct {
	elements []int
}

func New() *MinHeap {
	return &MinHeap{}
}

// Time: O(logn)
// Push appends the a new value to the heap, then updates the ordering
func (h *MinHeap) Push(value int) {
	h.elements = append(h.elements, value)
	h.heapifyUp(len(h.elements) - 1)
}

// Time: O(logn)
// Pop returns the largest value in the heap, then updates the ordering
func (h *MinHeap) Pop() (int, error) {
	if h.Size() == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	min := h.elements[0]
	lastIdx := len(h.elements) - 1
	h.elements[0] = h.elements[lastIdx]
	h.elements = h.elements[:lastIdx]
	h.heapifyDown(0)

	return min, nil
}

// Time: O(1)
// Peek returns the largest value of the maxheap without popping it
func (h *MinHeap) Peek() (int, error) {
	if h.Size() == 0 {
		return 0, fmt.Errorf("cannot peek an empty maxheap")
	}
	return h.elements[0], nil
}

// Time: O(1)
// Size returns the number of elements in the heap
func (h *MinHeap) Size() int {
	return len(h.elements)
}

// Time: O(n)
// Heapify pushes a slice of values to the heap
func (h *MinHeap) Heapify(values []int) {
	for _, val := range values {
		h.Push(val)
	}
}

// Time: O(logn)
// heapifyUp moves the value up the tree by exchanging the value with its parent node
// until it is in the correct position
func (h *MinHeap) heapifyUp(idx int) {
	parentIdx := (idx - 1) / 2

	for idx > 0 && h.elements[idx] < h.elements[parentIdx] {
		h.elements[idx], h.elements[parentIdx] = h.elements[parentIdx], h.elements[idx]
		idx = parentIdx
		parentIdx = (idx - 1) / 2
	}
}

// Time: O(logn)
// heapifyUp moves the value down the tree by exchanging the value with its child node
// until it is in the correct position
func (h *MinHeap) heapifyDown(idx int) {
	size := len(h.elements)
	leftChildIdx := 2*idx + 1
	rightChildIdx := 2*idx + 2
	smallest := idx

	if leftChildIdx < size && h.elements[leftChildIdx] < h.elements[smallest] {
		smallest = leftChildIdx
	}

	if rightChildIdx < size && h.elements[rightChildIdx] < h.elements[smallest] {
		smallest = rightChildIdx
	}

	if smallest != idx {
		h.elements[idx], h.elements[smallest] = h.elements[smallest], h.elements[idx]
		h.heapifyDown(smallest)
	}
}
