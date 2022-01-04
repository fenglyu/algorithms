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

//var _ MergeInterface = (*StringSlice)(nil)
//var _ MergeInterface = (*IntSlice)(nil)

/*
// Clone slices
// func Clone(from MergeInterface) interface{} {
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

	// fmt.Println("dc -> elem -> interface", dc.Elem().Interface())
	indexSetFunc := func(in []reflect.Value) []reflect.Value {
		//if len(in) < 3 {
		// pointer, index, value is missing
		fieldP, flag := getType.FieldByName("Slices")
		if !flag {
			fmt.Println("Error ")
		}
		in[1].Elem()
		//}
		return s.Slices[a]
	}

	//makeSwap := func(fptr interface{}) {
	//	fn := reflect.ValueOf(fptr).Elem()
	//	v := reflect.MakeFunc(fn.Type(), swap)
	//	fn.Set(v)
	//}

	//	var intSwap func(int, int) (int, int)
	//	makeSwap(&intSwap)
	//	fmt.Println(intSwap(0, 1))
	//
	//	// Make and call a swap function for float64s.
	//	var floatSwap func(float64, float64) (float64, float64)
	//	makeSwap(&floatSwap)
	//	fmt.Println(floatSwap(2.72, 3.14))
	makeIndexOrSet := func(fptr interface{}) {

	}
	return dc.Elem().Interface()
}


func Clone(from MergeInterface) MergeInterface {
	var cloned MergeInterface
	getType := reflect.TypeOf(from).Elem()
	cloned = from
	if getType.NumField() < 1 {
		fmt.Errorf("Fields number is 0\n")
		return nil
	}
	fieldP, flag := getType.FieldByName("Slices") // 拿到 p 的字段信息。
	if !flag {
		fmt.Println("Error ")
	}
	tv := reflect.ValueOf(from).Elem() // 拿到 t 的反射值。
	p := tv.FieldByIndex(fieldP.Index) // 拿到 t.p 的反射值，虽然不让写，但是可以读。

	num1 := p.Len()
	c := p.Cap()
	clonedP := reflect.MakeSlice(fieldP.Type, num1, c) // 构造一个新的 slice。
	src := unsafe.Pointer(p.Pointer())                 // 拿到 p 数据指针。
	dst := unsafe.Pointer(clonedP.Pointer())           // 拿到 clonedP 数据指针。
	sz := int(p.Type().Elem().Size())                  // 计算出 []int 单个数组元素的长度，即 int 的长度。
	l := num1 * sz                                     // 得到 p 的数据真实内存字节数。
	cc := c * sz                                       // 得到 p 的 cap 真实内存字节数。

	// 直接无视类型进行内存拷贝，相当于 C 语言里面的 memcpy。
	copy((*[math.MaxInt32]byte)(dst)[:l:cc], (*[math.MaxInt32]byte)(src)[:l:cc])

	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(reflect.ValueOf(cloned).Pointer())) + fieldP.Offset) // 拿到 p 的真实内存位置。

	// 这里已知 p 是一个 slice，用 `SliceHeader` 进行强制拷贝，相当于做了 cloned.p = clonedP。
	*(*reflect.SliceHeader)(ptr) = reflect.SliceHeader{
		Data: uintptr(dst),
		Len:  num1,
		Cap:  c,
	}

	return cloned
}
*/

type Cloneable interface {
	Clone() (any, string, error)
}

type IndexInterface interface {
	IndexOrSet(i int, val any) any
}

type MergeInterface interface {
	Interface
	IndexInterface
	Cloneable
}

func copySlice(x any) (any, error) {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("must pass a value with kind of Slice; got %v", v.Kind())
	}

	size := v.Len()
	t := reflect.TypeOf(x)
	dc := reflect.MakeSlice(t, size, size)

	for i := 0; i < size; i++ {
		iv := reflect.ValueOf(v.Index(i).Interface())
		if iv.IsValid() {
			dc.Index(i).Set(iv)
		}
	}

	return dc.Interface(), nil
}

func mergeSort(data MergeInterface) {
	//cp, dataType, err := data.Clone()
	cp, _, err := data.Clone()
	if err != nil {
		fmt.Println("err: ", err)
	}
	mergesort_array(cp.(MergeInterface), data, 0, data.Len())
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
