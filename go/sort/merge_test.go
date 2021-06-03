package sort

import (
	"reflect"
	"testing"
)

func TestMergeSortInt(t *testing.T) {
	tests, expected := &IntSlice{[]int{2, 9, 3, 5, 1, 7}}, &IntSlice{[]int{1, 2, 3, 5, 7, 9}}
	mergeSort(tests)
	if !reflect.DeepEqual(tests, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests)
	}
}
