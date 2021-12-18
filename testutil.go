package collection

import (
	"errors"
	"reflect"
	"testing"
)

// MustEqual let the test fail if the given values are not the same.
func MustEqual(t *testing.T, expected, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected: %v but got: %v", expected, got)
	}
}

// MustBeNil let the test fail if the given value is non nil.
func MustBeNil(t *testing.T, v interface{}) {
	t.Helper()
	if v != nil {
		t.Errorf("nil is expected but got %v", v)
	}
}

// MustBeErr let the test fail if an underlying error of the given error is not the same with the expected error.
func MustBeErr(t *testing.T, expected, got error) {
	t.Helper()
	if !errors.Is(got, expected) {
		t.Errorf("expected error: %v but got %v", expected, got)
	}
}

// iteratorMustEqual makes sure the given list is logically the same as the given slice.
// For example, a linkedlist 1 -> 2 -> 3 -> nil is considered to be the same as [1, 2, 3].
func iteratorMustEqual[T any](t *testing.T, i Iterator[T], values []T) {
	t.Helper()
	buff := []T{}
	cnt := 0
	for i.Next() {
		buff = append(buff, i.Value())
		cnt++
	}
	MustEqual(t, values, buff)
	MustEqual(t, len(values), cnt)
}
