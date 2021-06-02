package sort

import (
	"reflect"
	s "testing"
)

func TestInsertSortInt(t *testing.T) {
	tests, expected := &IntSlice{[]int{2, 9, 3, 5, 1, 7}}, &IntSlice{[]int{1, 2, 3, 5, 7, 9}}
	mergeSort(tests)
	if !reflect.DeepEqual(tests, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests)
	}
}
