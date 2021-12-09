package collection

import (
	"errors"
	"reflect"
	"testing"
)

func MustEqual(t *testing.T, expected, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected: %v but got: %v", expected, got)
	}
}

func MustBeNil(t *testing.T, v interface{}) {
	t.Helper()
	if v != nil {
		t.Errorf("nil is expected but got %v", v)
	}
}

func MustBeErr(t *testing.T, expected, got error) {
	t.Helper()
	if !errors.Is(got, expected) {
		t.Errorf("expected error: %v but got %v", expected, got)
	}
}
