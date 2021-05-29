package heap

import "sort"

type Heap struct {
	len int
	arr []interface{}
}

type HeapInterface interface {
	sort.Interface
}

func (h *Heap) buildHeap() {

}

func (h *Heap) heapfy() {

}
