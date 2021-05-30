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

type StringSlice struct {
	slices []string
}

func (s *StringSlice) Len() int {
	return len(s.slices)
}

func (s *StringSlice) Less(a, b int) bool {
	short, long := []byte(s.slices[a]), []byte(s.slices[b])

	for i, _ := range short {
		if i > len(long)-1 {
			// b has no more character, a is bigger
			return false
		}

		if short[i] > long[i] {
			return false
		}
	}

	return true
}

func (s *StringSlice) Swap(a, b int) {
	s.slices[a], s.slices[b] = s.slices[b], s.slices[a]
}
