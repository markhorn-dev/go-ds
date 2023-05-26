package main

import (
	"fmt"

	"github.com/markhorn-dev/go-ds/maxheap"
	"github.com/markhorn-dev/go-ds/minheap"
)

func main() {

	fmt.Println("maxheap")
	maxh := maxheap.New()
	maxh.Heapify([]int{75, 50, 99, 1, 20})

	for maxh.Size() != 0 {
		max, err := maxh.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(max)
	}

	fmt.Println("minheap")
	minh := minheap.New()
	minh.Heapify([]int{75, 50, 99, 1, 20})

	for minh.Size() != 0 {
		min, err := minh.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(min)
	}

}
