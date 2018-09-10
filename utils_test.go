package bullshit

import (
	"testing"
)

var shuffleTests = [][]string{
	{"a", "b", "c", "d"},
	{"1", "2", "3", "4"},
}

func TestShuffle(t *testing.T) {
	for _, s := range shuffleTests {
		got := Shuffle(s)
		if len(got) != len(s) {
			t.Errorf("input length: %d, result length: %d\n", len(s), len(got))
		}
		if !containsAll(got, s) {
			t.Errorf("result %q doesn't contain all elements of input %q\n", got, s)
		}
	}
}

func containsAll(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	backup := make([]string, len(b))
	copy(backup, b)
	for i := range a {
		if i, ok := containsAt(backup, a[i]); ok {
			backup = append(backup[:i], backup[i+1:]...)
		} else {
			return false
		}
	}
	return true
}

func containsAt(a []string, s string) (int, bool) {
	for i := range a {
		if a[i] == s {
			return i, true
		}
	}
	return -1, false
}
