package arraylist

type Iterator[T any] struct {
	list   *ArrayList[T]
	cursor int
}

func (i *Iterator[T]) Next() bool {
	return i.cursor < i.list.Len()
}

func (i *Iterator[T]) Value() T {
	// ignore error because i.cursor must be the valid value
	v, _ := i.list.GetAt(i.cursor)
	i.cursor++
	return v
}
