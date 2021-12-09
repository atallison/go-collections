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

// IndexOf returns the first index of an value which is the same as given v.
// It returns negative value if v is not found in the list.
func (a *ComparableArrayList[T]) IndexOf(v T) int {
	for i, av := range a.ArrayList.values {
		if v == av {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the last index of an value which is the same as given v.
// It returns negative value if v is not found in the list.
func (a *ComparableArrayList[T]) LastIndexOf(v T) int {
	for i := len(a.ArrayList.values) - 1; i >= 0; i-- {
		if v == a.ArrayList.values[i] {
			return i
		}
	}
	return -1
}
