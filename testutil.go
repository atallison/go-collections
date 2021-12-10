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
