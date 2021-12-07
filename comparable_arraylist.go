package collection

type ComparableArrayList[T comparable] struct {
	*ArrayList[T]
}

// NewComparableArrayList returns an ArrayList based on the specified type which is comparable.
func NewComparableArrayList[T comparable]() *ComparableArrayList[T] {
	return &ComparableArrayList[T]{ArrayList: NewArrayList[T]()}
}

// Remove removes the same value with given v in the list.
// It uses == operator to make sure if the values are the same.
func (a *ComparableArrayList[T]) Remove(v T) {
	buff := []T{}
	for _, av := range a.ArrayList.values {
		if v == av {
			continue
		}
		buff = append(buff, av)
	}
	a.values = buff
}

// Contains returns if the given value is contained in the list.
func (a *ComparableArrayList[T]) Contains(v T) bool {
	for _, av := range a.ArrayList.values {
		if v == av {
			return true
		}
	}
	return false
}
