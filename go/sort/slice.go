package sort

import (
	"fmt"
)

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

func (i *IntSlice) HashCode(v interface{}) uint64 {
	return countHash(v)
}

func (i *IntSlice) IndexOrSet(a int, val interface{}) interface{} {
	if val != nil {
		i.Slices[a] = val.(int)
	}
	return i.Slices[a]
}

func (i *IntSlice) Convert(slices []interface{}) []int {
	s := make([]int, len(slices))
	for i, v := range slices {
		s[i] = v.(int)
	}
	return s
}

/*
func (i *IntSlice) Clone(inter interface{}) interface{} {
	nInter := reflect.New(reflect.TypeOf(inter).Elem())

	//	fmt.Printf("reflect Elem() %v \n", nInter)
	val := reflect.ValueOf(inter).Elem()
	nVal := nInter.Elem()
	for i := 0; i < val.NumField(); i++ {
		nvField := nVal.Field(i)
		nvField.Set(val.Field(i))
	}

	return nInter.Interface()
}
*/
func (i *IntSlice) Clone() (interface{}, error) {
	cp := *i
	ns, err := copySlice(i.Slices)
	if err != nil {
		return nil, err
	}
	cp.Slices = ns.([]int)
	//fmt.Println("cp.Slices: ", &cp.Slices[0] == &i.Slices[0])
	//fmt.Println("cp ->: ", cp, *i)
	return &cp, nil
}

func (i *IntSlice) String() string {
	return fmt.Sprintf("%v", i.Slices)
}

type StrSlice struct {
	Slices []string
}

func (s *StrSlice) Len() int {
	return len(s.Slices)
}

func (s *StrSlice) Less(a, b int) bool {
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

func (s *StrSlice) Swap(a, b int) {
	s.Slices[a], s.Slices[b] = s.Slices[b], s.Slices[a]
}

func (s *StrSlice) HashCode(v interface{}) uint64 {
	return countHash(v)
}

func (s *StrSlice) IndexOrSet(a int, val interface{}) interface{} {
	if val != nil {
		s.Slices[a] = val.(string)
		//s.Slices[a] = val.(string)
	}
	return s.Slices[a]
}

func (s *StrSlice) Convert(slices []interface{}) []string {
	t := make([]string, len(slices))
	for i, v := range slices {
		t[i] = v.(string)
	}
	return t
}

func (s *StrSlice) Clone() (interface{}, error) {
	cp := *s
	ns, err := copySlice(s.Slices)
	if err != nil {
		return nil, err
	}
	cp.Slices = ns.([]string)
	fmt.Println("cp.Slices: ", &cp.Slices[0] == &s.Slices[0])
	fmt.Println("cp ->: ", cp, *s)
	return cp, nil
}
