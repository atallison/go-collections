package collection

type ComparableArrayList[T comparable] struct {
	*ArrayList[T]
}

// NewComparableArrayList returns an ArrayList based on the specified type which is comparable.
func NewComparableArrayList[T comparable]() *ComparableArrayList[T] {
	return &ComparableArrayList[T]{ArrayList: NewArrayList[T]()}
}

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
