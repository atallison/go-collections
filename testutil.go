package collection

import (
	"errors"
	"reflect"
	"testing"
)

// MustEqual let the test fail if the given values are not the same.
func MustEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v but actual: %v", expected, actual)
	}
}

// MustBeNil let the test fail if the given value is non nil.
func MustBeNil(t *testing.T, v interface{}) {
	t.Helper()
	if v != nil {
		t.Errorf("nil is expected but actual %v", v)
	}
}

// MustBeErr let the test fail if an underlying error of the given error is not the same with the expected error.
func MustBeErr(t *testing.T, expected, actual error) {
	t.Helper()
	if !errors.Is(actual, expected) {
		t.Errorf("expected error: %v but actual %v", expected, actual)
	}
}

// collectionMustEqual makes sure the given list is logically the same as the given slice.
// For example, a linkedlist 1 -> 2 -> 3 -> nil is considered to be the same as [1, 2, 3].
func collectionMustEqual[T any](t *testing.T, values []T, c collection[T]) {
	t.Helper()
	buff := []T{}
	i := c.Iterator()
	for i.Next() {
		buff = append(buff, i.Value())
	}
	MustEqual(t, values, buff)
	MustEqual(t, len(values), c.Len())
}
