package sort

import "reflect"

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type HashInterface interface {
	Interface
	IndexInterface
	HashCode(v interface{}) uint64
}

var _ HashInterface = (*StringSlice)(nil)
var _ HashInterface = (*IntSlice)(nil)

//var _ HashInterface = StringSlice{} // Verify that T implements I.
//var _ HashInterface = IntSlice{}    // Verify that T implements I.
//var _ I = (*T)(nil)                 // Verify that *T implements I.
func insertSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i - 1; j >= a && data.Less(j+1, j); j-- {
			data.Swap(j+1, j)
		}
	}
}

func quickSort(data Interface, a, b int) {
	if a >= b {
		return
	}
	p := partition(data, a, b)
	quickSort(data, a, p) // uppder boundary doesn't include p
	quickSort(data, p+1, b)
}

func selectPivot(data Interface, a, b int) int {
	return (a + b) / 2
}

func partition(data Interface, a, b int) int {
	pivotIdx := selectPivot(data, a, b)

	data.Swap(b-1, pivotIdx)

	store := a
	for i := a; i < b-1; i++ {
		if data.Less(i, b-1) {
			data.Swap(i, store)
			store++
		}
	}
	data.Swap(store, b-1)
	return store
}

func countSortInt(data []interface{}) {
	bucketSize := 100
	bucket := make([][]int, bucketSize)
	for i := 0; i < len(data); i++ {
		idx := int(countHash(data[i]) % uint64(bucketSize))
		if bucket[idx] == nil {
			bucket[idx] = make([]int, 0)
		}
		bucket[idx] = append(bucket[idx], data[i].(int))
	}
	extract(bucket, data)
}

func extract(bucket [][]int, data []interface{}) {
	idx := 0
	for i := 0; i < len(bucket); i++ {
		insertSort(&IntSlice{bucket[i]}, 0, len(bucket[i]))
		for _, v := range bucket[i] {
			data[idx] = v
			idx++
		}
	}
}

func numBuckets() int {
	return 26 * 26 * 26
}

func countSortInter(data HashInterface) {
	bucketSize := numBuckets()
	bucket := make([][]interface{}, bucketSize)
	for i := 0; i < data.Len(); i++ {
		idx := int(data.HashCode(data.IndexOrSet(i, nil)) % uint64(bucketSize))
		if bucket[idx] == nil {
			bucket[idx] = make([]interface{}, 0)
		}
		bucket[idx] = append(bucket[idx], data.IndexOrSet(i, nil))
	}
	extractInter(bucket, data)
}

func extractInter(bucket [][]interface{}, data HashInterface) {
	idx := 0
	var sls StringSlice
	var ils IntSlice

	for i := 0; i < len(bucket); i++ {
		slices := bucket[i]
		if len(slices) < 1 {
			continue
		}
		switch v := reflect.ValueOf(slices[0]); v.Kind() {
		case reflect.String:
			s := make([]string, len(slices))
			for i, v := range slices {
				s[i] = v.(string)
			}
			sls = StringSlice{Slices: s}
			insertSort(&sls, 0, len(bucket[i]))
			for _, v := range sls.Slices {
				data.IndexOrSet(idx, v)
				idx++
			}
		case reflect.Int:
			s := make([]int, len(slices))
			for i, v := range slices {
				s[i] = v.(int)
			}
			ils = IntSlice{Slices: s}
			insertSort(&ils, 0, len(bucket[i]))
			for _, v := range ils.Slices {
				data.IndexOrSet(idx, v)
				idx++
			}
		}
	}
}

func convertSlice(slices []interface{}) interface{} {
	switch v := reflect.ValueOf(slices[0]); v.Kind() {
	case reflect.String:
		s := make([]string, len(slices))
		for i, v := range slices {
			s[i] = v.(string)
		}
		return StringSlice{Slices: s}
	case reflect.Int:
		s := make([]int, len(slices))
		for i, v := range slices {
			s[i] = v.(int)
		}
		return IntSlice{Slices: s}
	}

	return nil
}
