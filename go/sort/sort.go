package sort

import "fmt"

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func insertSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i - 1; j >= a && data.Less(j+1, j); j-- {
			data.Swap(j+1, j)
		}
	}
}

func quickSort(data Interface, a , b int) {
	if a >= b {
		return 
	}
	p := partition(data, a, b)
	quickSort(data, a, p)  // uppder boundary doesn't include p
	quickSort(data, p+1, b)
}

func selectPivot(data Interface, a, b int) int{
	return (a + b) / 2 
}

func partition(data Interface, a, b int) int{
	pivotIdx := selectPivot(data,a,b)

	data.Swap(b-1, pivotIdx)

	store := a
	for i:=a; i< b-1 ;i++{
		if data.Less(i, b-1) {
			data.Swap(i, store)
			store++
		}
	}
	data.Swap(store, b-1)
	return store
}

type IntSlice struct {
	Slices []int
}

func (i *IntSlice) Len() int {
	return len(i.Slices)
}

func (is *IntSlice) Less(i, j int) bool {
	return is.Slices[i] < is.Slices[j]
}

func (i *IntSlice) Swap(a, b int) {
	i.Slices[a], i.Slices[b] = i.Slices[b], i.Slices[a]
}

func (i *IntSlice) String() string {
	return fmt.Sprintf("%v", i.Slices)
}

type StringSlice struct {
	Slices []string
}

func (s *StringSlice) Len() int {
	return len(s.Slices)
}

func (s *StringSlice) Less(a, b int) bool {
	short, long := []byte(s.Slices[a]), []byte(s.Slices[b])
	for i, _ := range short {
		if i > len(long)-1 {
			// b has no more character, a is bigger
			return false
		}
		if short[i] == long[i]{
			continue
		}else if short[i] > long[i] {
			return false
		}else{
			return true
		}
	}

	return true
}

func (s *StringSlice) Swap(a, b int) {
	s.Slices[a], s.Slices[b] = s.Slices[b], s.Slices[a]
}
