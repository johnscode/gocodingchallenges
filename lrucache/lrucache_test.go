package lrucache

import "testing"

const (
	first  = "first"
	second = "second"
	third  = "third"
	fourth = "fourth"
	fifth  = "fifth"
	sixth  = "sixth"
)

func fillCache() LRUCache[string] {
	c := NewLRUCache[string](5)
	c.Put(first, first)
	c.Put(second, second)
	c.Put(third, third)
	c.Put(fourth, fourth)
	c.Put(fifth, fifth)
	return c
}

func TestLruCache_Len(t *testing.T) {
	c := NewLRUCache[string](5)
	if c.Len() != 0 {
		t.Errorf("expected empty")
	}
}

func TestLruCache_Capacity(t *testing.T) {
	c := NewLRUCache[string](5)
	if c.Capacity() != 5 {
		t.Errorf("expected cap 5")
	}
}

func TestLruCache_Keys(t *testing.T) {
	c := fillCache()
	keys := c.Keys()
	if !containsInOrder(keys, []string{fifth, fourth, third, second, first}) {
		t.Errorf("not expected array")
	}
}

func TestLruCache_Put(t *testing.T) {
	t.Run("simple put", func(t *testing.T) {
		c := NewLRUCache[string](5)
		c.Put(first, first)
		if c.Len() != 1 {
			t.Errorf("expected 1 element")
		}
	})
	t.Run("put at capacity", func(t *testing.T) {
		c := fillCache()
		if c.Len() != 5 {
			t.Errorf("expected len 5")
		}
		c.Put(sixth, sixth)
		if c.Len() != 5 {
			t.Errorf("expected len still 5")
		}
	})
}

func TestLruCache_Get(t *testing.T) {
	t.Run("simple get", func(t *testing.T) {
		c := fillCache()
		if str, ok := c.Get(fifth); ok {
			if str != fifth {
				t.Errorf("expected fifth str")
			}
		}
	})
	t.Run("get after capacity", func(t *testing.T) {
		c := fillCache()
		c.Put(sixth, sixth)
		keys := c.Keys()
		if !containsInOrder(keys, []string{sixth, fifth, fourth, third, second}) {
			t.Errorf("not expected array")
		}
		if _, ok := c.Get(first); ok {
			t.Errorf("expected first to be ejected")
		}
	})
}

func TestLruCache_Remove(t *testing.T) {
	t.Run("remove existing", func(t *testing.T) {
		c := fillCache()
		if ok := c.Remove(third); !ok {
			t.Errorf("should have removed existing key")
		}
		if c.Len() != 4 {
			t.Errorf("expected len 4")
		}
		keys := c.Keys()
		if !containsInOrder(keys, []string{fifth, fourth, second, first}) {
			t.Errorf("not expected array")
		}
	})
	t.Run("remove most recent", func(t *testing.T) {
		c := fillCache()
		if ok := c.Remove(fifth); !ok {
			t.Errorf("should have removed existing key")
		}
		keys := c.Keys()
		if !containsInOrder(keys, []string{fourth, third, second, first}) {
			t.Errorf("not expected array")
		}
	})
	t.Run("remove least recent", func(t *testing.T) {
		c := fillCache()
		if ok := c.Remove(first); !ok {
			t.Errorf("should have removed existing key")
		}
		keys := c.Keys()
		if !containsInOrder(keys, []string{fifth, fourth, third, second}) {
			t.Errorf("not expected array")
		}
	})
	t.Run("remove nonexisting", func(t *testing.T) {
		c := fillCache()
		if ok := c.Remove(sixth); ok {
			t.Errorf("should not have removed nonexisting key")
		}
	})
}

func TestLruCache_Clear(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		c := fillCache()
		c.Clear()
		if c.Len() != 0 {
			t.Errorf("expected empty")
		}
	})
	t.Run("recover after clear", func(t *testing.T) {
		c := fillCache()
		c.Clear()
		c.Put(sixth, sixth)
		c.Put(first, first)
		if c.Len() != 2 {
			t.Errorf("expected len 2")
		}
		keys := c.Keys()
		if !containsInOrder(keys, []string{first, sixth}) {
			t.Errorf("unexpected order")
		}
	})
}

func containsInOrder(arr []string, sequence []string) bool {
	if len(sequence) == 0 {
		return true
	}
	if len(arr) < len(sequence) {
		return false
	}

	seqIndex := 0
	for _, item := range arr {
		if item == sequence[seqIndex] {
			seqIndex++
			if seqIndex == len(sequence) {
				return true
			}
		}
	}
	return false
}
