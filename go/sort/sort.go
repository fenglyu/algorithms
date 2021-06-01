package sort

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// ?
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

func countSortInter(data Interface) {
	bucketSize := 100
	bucket := make([]interface{}, bucketSize)
	for i := 0; i < len(data); i++ {
		idx := int(data.HashCode(data[i]) % uint64(bucketSize))
		if bucket[idx] == nil {
			bucket[idx] = make([]interface{}, 0)
		}
		bucket[idx] = append(bucket[idx], data[i])
	}
}

func extractInter(bucket []interface{}, data Interface) {
	idx := 0
	for i := 0; i < len(bucket); i++ {
		insertSort(&IntSlice{bucket[i]}, 0, len(bucket[i]))
		for _, v := range bucket[i] {
			data[idx] = v
			idx++
		}
	}
}
