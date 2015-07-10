package set

var sval = struct{}{}

type set map[interface{}]struct{}

type Set struct {
	set set
}

func New() *Set {
	return &Set{make(map[interface{}]struct{})}
}

func (s *Set) Add(values ...interface{}) {
	if len(values) == 0 {
		return
	}
	for _, v := range values {
		s.set[v] = sval
	}
}

func (s *Set) Remove(values ...interface{}) {
	if len(values) == 0 {
		return
	}

	for _, v := range values {
		delete(s.set, v)
	}
}

func (s *Set) Has(value interface{}) bool {
	_, ok := s.set[value]
	return ok
}

func (s *Set) Size() int {
	return len(s.set)
}

func (s *Set) Iter(cb func(item interface{})) {
	for item, _ := range s.set {
		cb(item)
	}
}

func (s *Set) Copy() *Set {
	set := New()
	for item, _ := range s.set {
		set.Add(item)
	}
	return set
}

func (s *Set) Map() map[interface{}]struct{} {
	return s.set
}

func (s *Set) Merge(set *Set) {
	set.Iter(func(item interface{}) {
		s.Add(item)
	})
}
