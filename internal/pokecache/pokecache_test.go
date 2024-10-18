package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	duration := time.Millisecond * 10
	cache := NewCache(duration)

	if cache.cache == nil {
		t.Errorf("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	duration := time.Millisecond * 10
	cache := NewCache(duration)

	cases := []struct {
		InputKey string
		InputVal []byte
	}{
		{
			InputKey: "key1",
			InputVal: []byte("blah"),
		},
		{
			InputKey: "key2",
			InputVal: []byte("val2"),
		},
	}
	for _, trial := range cases {

		cache.Add(trial.InputKey, trial.InputVal)
		actual, ok := cache.Get(trial.InputKey)
		if !ok {
			t.Errorf("error with Add and/or Get")
			continue
		}

		if string(actual) != string(trial.InputVal) {
			t.Errorf("return value does not match %v %v", actual, trial.InputVal)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	duration := time.Millisecond * 10
	cache := NewCache(duration)

	cases := []struct {
		InputKey string
		InputVal []byte
	}{
		{
			InputKey: "key1",
			InputVal: []byte("blah"),
		},
		{
			InputKey: "key2",
			InputVal: []byte("val2"),
		},
	}
	for _, trial := range cases {

		cache.Add(trial.InputKey, trial.InputVal)
		time.Sleep(duration + 5*time.Millisecond)
		_, ok := cache.Get(trial.InputKey)
		if ok {
			t.Errorf("error with reap loop")
			continue
		}

	}
}
