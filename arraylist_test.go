package collection

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	al := NewArrayList[int]()
	al.Add(1)
	al.Add(2)
	if !reflect.DeepEqual(al.values, []int{1, 2}) {
		t.Errorf("expected: %v but got %v", []int{1, 2}, al.values)
	}
}

func TestAddAll(t *testing.T) {
	al := NewArrayList[int]()
	al.Add(1)
	al.AddAll([]int{2, 3, 4})
	al.AddAll([]int{5})
	if !reflect.DeepEqual(al.values, []int{1, 2, 3, 4, 5}) {
		t.Errorf("expected: %v but got %v", []int{1, 2, 3, 4, 5}, al.values)
	}
}

func TestAddAt(t *testing.T) {
	al := NewArrayList[int]()
	if err := al.AddAt(0, 1); err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	if err := al.AddAt(3, 2); err == nil {
		t.Errorf("err should be returned")
	}

	if err := al.AddAt(2, 2); err == nil {
		t.Errorf("err should be returned")
	}

	if err := al.AddAt(1, 2); err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	if err := al.AddAt(1, 3); err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	if err := al.AddAt(0, 4); err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	if !reflect.DeepEqual(al.values, []int{4, 1, 3, 2}) {
		t.Errorf("expected: %v but got %v", []int{4, 1, 3, 2}, al.values)
	}
}

func TestAddAllAt(t *testing.T) {
	al := NewArrayList[int]()
	if err := al.AddAllAt(0, []int{1}); err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	if !reflect.DeepEqual(al.values, []int{1}) {
		t.Errorf("expected: %v but got %v", []int{1}, al.values)
	}
}
