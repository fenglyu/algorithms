package sort

import "fmt"

type IntSliceIface interface {
	Interface
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
		if short[i] == long[i] {
			continue
		} else if short[i] > long[i] {
			return false
		} else {
			return true
		}
	}

	return true
}

func (s *StringSlice) Swap(a, b int) {
	s.Slices[a], s.Slices[b] = s.Slices[b], s.Slices[a]
}
