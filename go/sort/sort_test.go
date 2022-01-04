package sort

import (
	"reflect"
	"sort"
	"testing"
)

func TestInsertSortInt(t *testing.T) {
	tests, expected := &IntSlice{[]int{2, 9, 3, 5, 1, 7}}, &IntSlice{[]int{1, 2, 3, 5, 7, 9}}
	insertSort(tests, 0, tests.Len())
	if !reflect.DeepEqual(tests, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests)
	}
}
func TestQuickSortInt(t *testing.T) {
	tests, expected := &IntSlice{[]int{2, 9, 3, 5, 1, 7}}, &IntSlice{[]int{1, 2, 3, 5, 7, 9}}
	quickSort(tests, 0, tests.Len())
	if !reflect.DeepEqual(tests, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests)
	}
}

func TestInsertSortStrings(t *testing.T) {
	tests, expected := &StringSlice{[]string{"azbcef", "abcde", "abccd", "aaa"}}, &StringSlice{[]string{"aaa", "abccd", "abcde", "azbcef"}}
	insertSort(tests, 0, tests.Len())
	if !reflect.DeepEqual(tests, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests)
	}
}

func TestQuickSortStrings(t *testing.T) {
	tests, expected := &StringSlice{[]string{"azbcef", "abcde", "abccd", "aaa"}}, &StringSlice{[]string{"aaa", "abccd", "abcde", "azbcef"}}
	quickSort(tests, 0, tests.Len())
	if !reflect.DeepEqual(tests, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests)
	}
}

func TestCountSortInt(t *testing.T) {
	//tests, expected := &IntSlice{[]int{2, 9, 3, 5, 1, 7}}, &IntSlice{[]int{1, 2, 3, 5, 7, 9}}
	//tests, expected := []int{2, 9, 3, 5, 1, 7}, []int{1, 2, 3, 5, 7, 9}
	tests, expected := []any{2, 9, 3, 5, 1, 7}, []any{1, 2, 3, 5, 7, 9}
	countSortInt(tests)
	if !reflect.DeepEqual(tests, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests)
	}
}

/*
func TestCountSortString(t *testing.T) {
	tests, expected := []interface{}{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}, []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(expected)
	countSort(tests)
	if !reflect.DeepEqual(tests, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests)
	}
}
*/

// hash algorithm not working for []string currently
func TestCountSortInterStr(t *testing.T) {
	tests, expected := &StringSlice{Slices: []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}}, []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(expected)
	countSortInter(tests)
	if !reflect.DeepEqual(tests.Slices, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests.Slices)
	}
}

func TestCountSortIterInt(t *testing.T) {
	tests, expected := &IntSlice{Slices: []int{2, 9, 3, 5, 1, 7}}, []int{1, 2, 3, 5, 7, 9}
	countSortInter(tests)
	if !reflect.DeepEqual(tests.Slices, expected) {
		t.Errorf("Expected %v, Actually result %v\n", expected, tests.Slices)
	}
}
