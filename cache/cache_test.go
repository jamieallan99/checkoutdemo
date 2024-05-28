package cache

import (
	"errors"
	"testing"
)

func TestGet(t *testing.T) {
	defer KillStore()
	store.datamap = map[string]any{
		"123": "some data",
	}

	testtable := []struct {
		name     string
		key      string
		expected any
		err      error
	}{
		{
			"Key exists",
			"123",
			"some data",
			nil,
		},
		{
			"Key not found",
			"1234",
			nil,
			ErrKeyNotFound,
		},
	}
	for _, tr := range testtable {
		t.Run(tr.name, func(t *testing.T) {
			value, err := Get(tr.key)
			if (tr.err == nil) != (err == nil) {
				if !errors.Is(ErrKeyNotFound, err) {
					t.Errorf("Incorrect error status expected: %v, got: %v", tr.err, err)
				}
			}
			if value != tr.expected {
				t.Errorf("Incorrect value expected: %v, got: %v", tr.expected, value)
			}
		})
	}
}
