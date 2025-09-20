package dot_test

import (
	"sync"
	"testing"

	"github.com/mirrorru/dot"
	"github.com/stretchr/testify/assert"
)

func TestSyncStore_Basic(t *testing.T) {
	t.Parallel()

	var s dot.SyncStore[string, int]
	s.Put("a", 1)
	s.Put("b", 2)

	v, ok := s.GetCurrent("a")
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = s.GetCurrent("b")
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	s.Del("a")
	_, ok = s.GetCurrent("a")
	assert.False(t, ok)
}

func TestSyncStore_Basic2(t *testing.T) {
	t.Parallel()

	var s dot.SyncStore[string, int]

	v := s.GetOrPut("a", func() int {
		return 1
	})
	assert.Equal(t, 1, v)
	v = s.GetOrPut("a", func() int {
		return -1
	})
	assert.Equal(t, 1, v)

	v = s.GetOrPut("b", func() int {
		return 2
	})
	assert.Equal(t, 2, v)
}

func TestSyncStore_ForEach(t *testing.T) {
	t.Parallel()

	var s dot.SyncStore[int, string]
	s.Put(1, "one")
	s.Put(2, "two")

	m := map[int]string{}
	s.ForEach(func(k int, v string) {
		m[k] = v
	})

	assert.Equal(t, "one", m[1])
	assert.Equal(t, "two", m[2])
}

func TestSyncStore_Iterator(t *testing.T) {
	t.Parallel()

	var s dot.SyncStore[int, string]
	s.Put(1, "one")
	s.Put(2, "two")

	result := map[int]string{}
	for pair := range s.Iterator() {
		result[pair.Key] = pair.Value
	}
	assert.Equal(t, "one", result[1])
	assert.Equal(t, "two", result[2])
}

func TestSyncStore_Seq(t *testing.T) {
	t.Parallel()

	var s dot.SyncStore[int, string]
	s.Put(1, "one")
	s.Put(2, "two")

	vals := map[string]bool{}
	s.Seq()(func(v string) bool {
		vals[v] = true
		return true
	})
	assert.True(t, vals["one"])
	assert.True(t, vals["two"])
}

func TestSyncStore_Seq2(t *testing.T) {
	t.Parallel()

	var s dot.SyncStore[int, string]
	s.Put(1, "one")
	s.Put(2, "two")

	vals := map[int]string{}
	s.Seq2()(func(k int, v string) bool {
		vals[k] = v
		return true
	})
	assert.Equal(t, "one", vals[1])
	assert.Equal(t, "two", vals[2])
}

func TestSyncStore_Concurrent(t *testing.T) {
	t.Parallel()

	var s dot.SyncStore[int, int]
	wg := sync.WaitGroup{}
	for i := range 100 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Put(i, i*i)
		}(i)
	}
	wg.Wait()

	count := 0
	s.ForEach(func(k, v int) {
		assert.Equal(t, k*k, v)
		count++
	})
	assert.Equal(t, 100, count)
}
