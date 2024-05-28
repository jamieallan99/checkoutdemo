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
				t.Errorf("Incorrect value, expected: %v, got: %v", tr.expected, value)
			}
		})
	}
}

func TestDel(t *testing.T) {
	defer KillStore()
	store.datamap = map[string]any{
		"123": "some data",
	}

	testtable := []struct {
		name string
		key  string
		err  error
	}{
		{
			"Key exists",
			"123",
			nil,
		},
		{
			"Key not found",
			"1234",
			ErrKeyNotFound,
		},
	}
	for _, tr := range testtable {
		t.Run(tr.name, func(t *testing.T) {
			err := Del(tr.key)
			if (tr.err == nil) != (err == nil) {
				if !errors.Is(ErrKeyNotFound, err) {
					t.Errorf("Incorrect error status expected: %v, got: %v", tr.err, err)
				}
			}
		})
	}
}

func TestPut(t *testing.T) {
	defer KillStore()
	store.datamap = map[string]any{
		"123": "some data",
	}

	testtable := []struct {
		name     string
		key      string
		value    any
		expected any
	}{
		{
			"Key exists",
			"123",
			"new test data",
			"new test data",
		},
		{
			"Key does not exist",
			"1234",
			"also new test data",
			"also new test data",
		},
	}
	for _, tr := range testtable {
		t.Run(tr.name, func(t *testing.T) {
			Put(tr.key, tr.value)
			value, err := Get(tr.key)
			if err != nil {
				if !errors.Is(ErrKeyNotFound, err) {
					t.Errorf("Unexpcted error: %v", err)
				}
			}
			if value != tr.expected {
				t.Errorf("Incorrect value, expected: %v, got: %v", tr.expected, value)
			}
		})
	}
}
