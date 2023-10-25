package dag

type Set map[any]any

type Hashable interface {
	Hashcode() any
}

func hashcode(v any) any {
	if h, ok := v.(Hashable); ok {
		return h.Hashcode()
	}
	return v
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) Add(v any) {
	s[hashcode(v)] = v
}

func (s Set) Delete(v any) {
	delete(s, hashcode(v))
}

func (s Set) Include(v any) bool {
	_, ok := s[hashcode(v)]
	return ok
}
func (s Set) Intersection(other Set) Set {
	result := make(Set)
	if s == nil || other == nil {
		return result
	}

	if other.Len() < s.Len() {
		s, other = other, s
	}

	for _, v := range s {
		if other.Include(v) {
			result.Add(v)
		}
	}
	return result
}

func (s Set) Difference(other Set) Set {
	if other == nil || other.Len() == 0 {
		return s.Copy()
	}

	r := make(Set)
	for _, v := range s {
		if !other.Include(v) {
			r.Add(v)
		}
	}
	return r
}

func (s Set) Filter(cb func(any) bool) Set {
	if cb == nil {
		return s.Copy()
	}
	c := make(Set)
	for _, v := range s {
		if cb(v) {
			c.Add(v)
		}
	}
	return c
}

func (s Set) Copy() Set {
	c := make(Set, s.Len())
	for k, v := range s {
		c[k] = v
	}
	return c
}
