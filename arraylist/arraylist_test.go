package arraylist

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	al := New[int]()
	al.Add(1)
	al.Add(2)
	if !reflect.DeepEqual(al.values, []int{1, 2}) {
		t.Errorf("expected: %v but got %v", []int{1, 2}, al.values)
	}
}
