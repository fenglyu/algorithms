package sort

import (
	"reflect"
	"testing"
)

func TestInsertSortInt(t *testing.T) {
	tests, expected := &IntSlice{[]int{2, 9, 3, 5, 1, 7}}, &IntSlice{[]int{1, 2, 3, 5, 7, 9}}
	insertSort(tests, 0, tests.Len())
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
