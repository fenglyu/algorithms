package heap

import "sort"


type Interface interface {
	sort.Interface
}

func heapSort(data Interface){
	buildHeap(data)
	for i:= data.Len() -1; i>0;i--{
		data.Swap(0, i)
		heapify(data, 0, i)
	}
}

func buildHeap(data Interface) {
	for i:= data.Len()/2 - 1; i>=0; i--{
		heapify(data, i, data.Len()-1)
	}
}

func heapify(data Interface, idx, max int) {
	l, r := 2 * idx +1, 2*idx+2
	largest := idx

	if l < max && data.Less(largest, l) {
		largest = l
	}
	if r < max  && data.Less(largest, r){
		largest = r 
	}
	if largest !=  idx{
		data.Swap(largest, idx)
		heapify(data, largest, max)
	}
}
