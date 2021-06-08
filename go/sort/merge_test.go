package sort

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeSortInt(t *testing.T) {
	tests, expected := &IntSlice{[]int{2, 9, 3, 5, 1, 7}}, &IntSlice{[]int{1, 2, 3, 5, 7, 9}}
	mergeSort(tests)
	if !reflect.DeepEqual(tests, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests)
	}
}

/*
 */
func TestMergeSortString(t *testing.T) {
	tests, expected := &StrSlice{Slices: []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}}, []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(expected)
	mergeSort(tests)
	if !reflect.DeepEqual(tests.Slices, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests.Slices)
	}
}
