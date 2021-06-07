package sort

import (
	"fmt"
	"reflect"
)

/*
Conceptually, a merge sort works as follows:
    Divide the unsorted list into n sublists, each containing one element (a list of one element is considered sorted).
    Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining. This will be the sorted list.
*/

type Cloneable interface {
	Clone(inter interface{}) interface{}
}

type MergeInterface interface {
	Interface
	IndexInterface
	Cloneable
}

// Clone slices
//func Clone(from MergeInterface) interface{} {
func Clone(from MergeInterface) MergeInterface {
	getType := reflect.TypeOf(from)
	getValue := reflect.ValueOf(from)
	dc := reflect.New(getType)
	for i := 0; i < getType.NumField(); i++ {
		f := getType.Field(i)
		if f.PkgPath != "" {
			continue
		}
		x := getValue.Field(i).Interface()
		fmt.Printf("Field(%d)", i, x)

		item, err := copySlice(x)
		if err != nil {
			fmt.Errorf("failed to copy the field %v in the struct %#v: %v", getType.Field(i).Name, x, err)
		}

		fmt.Println("item ", item)
		dc.Elem().Field(i).Set(reflect.ValueOf(item))
	}
	//fmt.Println("dc -> elem -> interface", dc.Elem().Interface())
	return dc.Elem().Interface()
}

func copySlice(x interface{}) (interface{}, error) {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("must pass a value with kind of Slice; got %v", v.Kind())
	}

	size := v.Len()
	t := reflect.TypeOf(x)
	dc := reflect.MakeSlice(t, size, size)

	for i := 0; i < size; i++ {
		//      kind := reflect.ValueOf(x).Kind()
		//      if kind == reflect.Array || kind == reflect.Chan || kind == reflect.Func || kind == reflect.Interface || kind == reflect.Map || kind == reflect.Ptr || kind == reflect.Slice || kind == reflect.Struct || kind == reflect.UnsafePointer {
		//          return nil, fmt.Errorf("unable to copy %v (a %v) as a primitive", x, kind)
		//      }
		iv := reflect.ValueOf(v.Index(i).Interface())
		if iv.IsValid() {
			dc.Index(i).Set(iv)
		}
	}

	return dc.Interface(), nil
}

func mergeSort(data MergeInterface) {
	cp := Clone(data)
	//cp := data.Clone(data).(*IntSlice)
	//cp := &IntSlice{[]int{2, 9, 3, 5, 1, 7}}
	//fmt.Println("cp == data", cp == data)
	mergesort_array(cp, data, 0, data.Len())
}

func mergesort_array(data MergeInterface, result MergeInterface, start, end int) {
	if end-start < 2 {
		return
	}

	if end-start == 2 {
		if result.Less(start+1, start) {
			result.Swap(start, start+1)
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
			result.IndexOrSet(idx, data.IndexOrSet(i, nil))
			i++
		} else {
			//result[idx] = data[j]
			result.IndexOrSet(idx, data.IndexOrSet(j, nil))
			j++
		}
		idx++
	}
}
