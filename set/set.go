package set

type Set[T comparable] struct {
	values map[T]struct{}
}

func New[T comparable]() *Set[T] {
	values := make(map[T]struct{})
	return &Set[T]{
		values: values,
	}
}

func (s *Set[T]) Len() int {
	return len(s.values)
}

func (s *Set[T]) Add(obj T) {
	s.values[obj] = struct{}{}
}

func (s *Set[T]) Contains(obj T) bool {
	_, ok := s.values[obj]
	return ok
}

func (s *Set[T]) Remove(obj T) {
	delete(s.values, obj)
}

func (s *Set[T]) Clear() {
	s.values = make(map[T]struct{})
}

func (s *Set[T]) Each(do func(i int, obj *T)) {
	i := 0
	for obj := range s.values {
		do(i, &obj)
		i++
	}
}

func (s *Set[T]) Any(match func(obj T) bool) bool {
	for obj := range s.values {
		if match(obj) {
			return true
		}
	}
	return false
}
