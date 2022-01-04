package heap

import (
	"reflect"
	"testing"

	"github.com/fenglyu/algorithms/go/sort"
)

func TestHeapSortInt(t *testing.T) {
	tests, expected := &sort.IntSlice{Slices: []int{2, 9, 3, 5, 1, 7}}, &sort.IntSlice{Slices: []int{1, 2, 3, 5, 7, 9}}
	heapSort(tests)
	if !reflect.DeepEqual(tests, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests)
	}
}
