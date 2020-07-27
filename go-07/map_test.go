package go_07

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "aa": 2, "aaa": 3}
	t.Log(m["a"])
}

func TestTraveMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 4, 3: 9}

	for k, v := range m {
		t.Log(k, v)
	}
}
