package set

import (
	"strings"
	"testing"

	"github.com/hidetatz/collection/testutil"
)

func TestNew(t *testing.T) {

	t.Run("New returns a set", func(t *testing.T) {
		s := New[string]()
		testutil.MustEqual(t, 0, s.Len())
	})

}

func TestSet_Clear(t *testing.T) {

	t.Run("Clear the set", func(t *testing.T) {
		s := New[string]()
		s.Add("test_a")
		s.Add("test_b")
		s.Clear()
		testutil.MustEqual(t, 0, s.Len())
	})

}

func TestSet_Add(t *testing.T) {

	t.Run("Add an item", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		testutil.MustEqual(t, 1, s.Len())
	})

	t.Run("Cannot add a duplicate", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		s.Add("test")
		testutil.MustEqual(t, 1, s.Len())
	})

	t.Run("Add an item with an empty value", func(t *testing.T) {
		s := New[string]()
		s.Add("")
		testutil.MustEqual(t, 1, s.Len())
	})

	t.Run("Cannot add the empty value twice", func(t *testing.T) {
		s := New[string]()
		s.Add("")
		s.Add("")
		testutil.MustEqual(t, 1, s.Len())
	})

}

func TestSet_Remove(t *testing.T) {

	t.Run("Remove an item", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		s.Remove("test")
		testutil.MustEqual(t, 0, s.Len())
	})

	t.Run("Remove an item that doesn't exist", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		s.Remove("missing")
		testutil.MustEqual(t, 1, s.Len())
	})

}

func TestSet_Contains(t *testing.T) {
	t.Run("Check if an item exists", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		contains := s.Contains("test")
		testutil.MustEqual(t, true, contains)
	})

	t.Run("Check if an item doesn't exist", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		contains := s.Contains("missing")
		testutil.MustEqual(t, false, contains)
	})

	t.Run("Check if the empty object exists", func(t *testing.T) {
		s := New[string]()
		contains := s.Contains("")
		testutil.MustEqual(t, false, contains)
	})
}

func TestSet_Any(t *testing.T) {

	t.Run("Run function on each", func(t *testing.T) {
		s := New[string]()
		s.Add("test_a")
		s.Add("test_b")
		contains := s.Any(func(obj string) bool { return strings.Contains(obj, "b") })
		testutil.MustEqual(t, true, contains)
	})

	t.Run("Any does not match an object in the set", func(t *testing.T) {
		s := New[string]()
		s.Add("test_a")
		s.Add("test_b")
		contains := s.Any(func(obj string) bool { return strings.Contains(obj, "c") })
		testutil.MustEqual(t, false, contains)
	})
}

func TestSet_Each(t *testing.T) {

	t.Run("Run function on each", func(t *testing.T) {
		s := New[string]()
		s.Add("test_a")
		s.Add("test_b")
		sb := strings.Builder{}
		s.Each(func(obj string) {
			sb.WriteString(obj)
		})
		contain := strings.Contains(sb.String(), "test_a") && strings.Contains(sb.String(), "test_b")

		testutil.MustEqual(t, true, contain)
	})
}

type testStruct struct {
	name string
	age  int
}

func TestSetStruct_Add(t *testing.T) {

	t.Run("Add an item", func(t *testing.T) {
		s := New[testStruct]()
		s.Add(testStruct{})
		testutil.MustEqual(t, 1, s.Len())
	})

	t.Run("Cannot add a duplicate", func(t *testing.T) {
		s := New[testStruct]()
		s.Add(testStruct{name: "Bob", age: 42})
		s.Add(testStruct{name: "Bob", age: 42})
		testutil.MustEqual(t, 1, s.Len())
	})
}

func TestSetStruct_Any(t *testing.T) {

	t.Run("Any matches a struct", func(t *testing.T) {
		s := New[testStruct]()
		s.Add(testStruct{name: "Bob", age: 42})
		s.Add(testStruct{name: "Alice", age: 34})
		contains := s.Any(func(obj testStruct) bool { return obj.name == "Bob" })
		testutil.MustEqual(t, true, contains)
	})
}
