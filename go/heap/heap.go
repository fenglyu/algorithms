package heap

import (
	"github.com/fenglyu/algorithms/go/sort"
)

type Interface interface {
	sort.Interface
}

func heapSort(data Interface) {
	buildHeap(data)
	for i := data.Len() - 1; i > 0; i-- {
		data.Swap(0, i)
		heapify(data, 0, i)
	}
}

func buildHeap(data Interface) {
	for i := data.Len()/2 - 1; i >= 0; i-- {
		heapify(data, i, data.Len()-1)
	}
}

func heapify(data Interface, idx, max int) {
	l, r := 2*idx+1, 2*idx+2
	largest := idx

	if l < max && data.Less(largest, l) {
		largest = l
	}
	if r < max && data.Less(largest, r) {
		largest = r
	}
	if largest != idx {
		data.Swap(largest, idx)
		heapify(data, largest, max)
	}
}

// siftDown implements the heap property on data[lo:hi].
// first is an offset into the array where the root of the heap lies.
func siftDown(data Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

func heapSortForLoop(data Interface, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}
