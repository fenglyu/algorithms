package sort

import (
	"reflect"
)

/*
Conceptually, a merge sort works as follows:
    Divide the unsorted list into n sublists, each containing one element (a list of one element is considered sorted).
    Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining. This will be the sorted list.
*/

type MergeInterface interface {
	Interface
	IndexInterface
}

func mergeSort(data MergeInterface) {
	copy := Clone(data)
	mergesort_array(copy, data, 0, data.Len())
}

func mergesort_array(data, result MergeInterface, start, end int) {
	if end-start < 2 {
		return
	}

	if end-start == 2 {
		if data.Less(start+1, start) {
			data.Swap(start, start+1)
		}
		return
	}

	mid := (end + start) / 2
	mergesort_array(result, data, start, mid)
	mergesort_array(result, data, mid, end)

	// merge A left- and right- side
	i, j, idx := start, mid, start
	for idx < end {
		if j >= end || (i < mid && data.Less(i, j)) {
			//result[i] = data[i]
			result.IndexOrSet(i, data.IndexOrSet(i, nil))
			i++
		} else {
			//result[idx] = data[j]
			result.IndexOrSet(idx, data.IndexOrSet(j, nil))
			j++
		}
		idx++
	}
}

// Clone slices
func Clone(from MergeInterface) MergeInterface {
	v := reflect.ValueOf(from)
	if v.Kind() != reflect.Slice || v.Kind() != reflect.Array {
		return nil
	}

	if v.Len() < 1 {
		return nil
	}

	switch k := v.Index(0); k.Kind() {
	case reflect.Int:
		//sType := reflect.SliceOf(reflect.TypeOf(int(0)))
		sType := reflect.TypeOf((*MergeInterface)(nil)) //.Elem()
		to := reflect.MakeSlice(sType, v.Len(), v.Len()*2)
		for i := 0; i < from.Len(); i++ {
			to.IndexOrSet(i, from.IndexOrSet(i, nil))
		}
		return to
	case reflect.String:
		sType := reflect.SliceOf(reflect.TypeOf(string("")))
		to := reflect.MakeSlice(sType, v.Len(), v.Len()*2)
		for i := 0; i < from.Len(); i++ {
			to.IndexOrSet(i, from.IndexOrSet(i, nil))
		}
		return to
	}
	return nil
}
