package linkedlist

import (
	"testing"

	"github.com/hidetatz/collection/testutil"
)

func pushBackAll[T any](t *testing.T, l *LinkedList[T], values []T) {
	t.Helper()
	for _, v := range values {
		l.PushBack(v)
	}
}

func assertLinkedListEq[T any](t *testing.T, values []T, l *LinkedList[T]) {
	t.Helper()
	testutil.MustEqual(t, len(values), l.length)

	// in case empty list is expected, (l.dummy == l.dummy.next == l.dummy.prev) is must
	if len(values) == 0 {
		testutil.MustEqual(t, l.dummy.next, l.dummy)
		testutil.MustEqual(t, l.dummy.prev, l.dummy)
		return
	}

	n := l.dummy.next
	p := l.dummy

	buff := []T{}
	for _ = range values {
		if p != n.prev {
			t.Errorf("Invalid LinkedList: %v is not the prev of %v\n", p.Value, n.prev.Value)
		}
		if p.next != n {
			t.Errorf("Invalid LinkedList: %v is not the next of %v\n", n.Value, p.next.Value)
		}
		buff = append(buff, n.Value)
		p = n
		n = n.next
	}

	if p != l.dummy.prev {
		t.Errorf("Invalid LinkedList: %v is not the prev of dummy\n", p.Value)
	}

	testutil.MustEqual(t, values, buff)
}

func TestLinkedList_PushBack(t *testing.T) {
	ll := New[int]()
	assertLinkedListEq(t, []int{}, ll)
	ll.PushBack(1)
	assertLinkedListEq(t, []int{1}, ll)
	ll.PushBack(2)
	assertLinkedListEq(t, []int{1, 2}, ll)
}

func TestLinkedList_PushFront(t *testing.T) {
	ll := New[int]()
	assertLinkedListEq(t, []int{}, ll)
	ll.PushFront(1)
	assertLinkedListEq(t, []int{1}, ll)
	ll.PushFront(2)
	assertLinkedListEq(t, []int{2, 1}, ll)
}

func TestLinkedList_Len(t *testing.T) {
	ll := New[int]()
	testutil.MustEqual(t, 0, ll.Len())

	ll.PushBack(1)
	testutil.MustEqual(t, 1, ll.Len())

	pushBackAll(t, ll, []int{2, 3, 4})
	testutil.MustEqual(t, 4, ll.Len())
}
