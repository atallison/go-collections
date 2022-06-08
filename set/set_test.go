package set

import (
	"strings"
	"testing"

	"github.com/hidetatz/collection/testutil"
)

func TestSet_Any(t *testing.T) {

	t.Run("empty", func(t *testing.T) {
		s := New[string]()
		testutil.MustEqual(t, 0, s.Len())
	})

	t.Run("non-empty", func(t *testing.T) {
		s := New[string]()
		s.Add("a")
		testutil.MustEqual(t, 1, s.Len())
	})

	t.Run("Cannot add duplicate", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		s.Add("test")
		testutil.MustEqual(t, 1, s.Len())
	})

	t.Run("Can remove", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		s.Remove("test")
		testutil.MustEqual(t, 0, s.Len())
	})

	t.Run("Remove an item that doesn't exist", func(t *testing.T) {
		s := New[string]()
		s.Remove("test")
		testutil.MustEqual(t, 0, s.Len())
	})

	t.Run("Run function on each", func(t *testing.T) {
		s := New[string]()
		s.Add("test_a")
		s.Add("test_b")
		contains := s.Any(func(obj string) bool { return strings.Contains(obj, "b") })
		testutil.MustEqual(t, true, contains)
	})
}
