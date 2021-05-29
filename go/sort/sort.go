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

type IntSlice struct {
	slices []int
}

func (i *IntSlice) Len() int {
	return len(i.slices)
}

func (is *IntSlice) Less(i, j int) bool {
	return is.slices[i] < is.slices[j]
}

func (i *IntSlice) Swap(a, b int) {
	i.slices[a], i.slices[b] = i.slices[b], i.slices[a]
}

func (i *IntSlice) String() string {
	return fmt.Sprintf("%v", i.slices)
}

type String struct {
	slices []string
}

func (s *String) Len() int {
	return len(s.slices)
}

func (s *String) Less(a, b int) bool {
	i, j := len(s.slices[a]), len(s.slices[b])
	if i == 0 && j == 0 {
		return false
	} else if i == 0  {
		
	}
}

func (s *String) Swap(a, b int) {

}
