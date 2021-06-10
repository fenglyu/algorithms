package sort

import (
	"fmt"
	"reflect"
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
func (i *IntSlice) Clone() (interface{}, string, error) {
	cp := *i
	tname := reflect.TypeOf(cp)
	ns, err := copySlice(i.Slices)
	if err != nil {
		return nil, tname.Name(), err
	}
	cp.Slices = ns.([]int)
	//fmt.Println("cp.Slices: ", &cp.Slices[0] == &i.Slices[0])
	//fmt.Println("cp ->: ", cp, *i)

	//return &cp, tname, nil
	return &cp, tname.Name(), nil
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

func (s *StringSlice) HashCode(v interface{}) uint64 {
	return countHash(v)
}

func (s *StringSlice) IndexOrSet(a int, val interface{}) interface{} {
	if val != nil {
		s.Slices[a] = val.(string)
		//s.Slices[a] = val.(string)
	}
	return s.Slices[a]
}

func (s *StringSlice) Convert(slices []interface{}) []string {
	t := make([]string, len(slices))
	for i, v := range slices {
		t[i] = v.(string)
	}
	return t
}

func (s *StringSlice) Clone() (interface{}, string, error) {
	cp := *s
	tname := reflect.TypeOf(cp)
	ns, err := copySlice(s.Slices)
	if err != nil {
		return nil, tname.Name(), err
	}
	cp.Slices = ns.([]string)
	return &cp, tname.Name(), nil
}
