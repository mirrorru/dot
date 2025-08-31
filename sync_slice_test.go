package dot

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncSlice_InitSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		length   int
		capacity int
	}{
		{"zero length", 0, 0},
		{"non-zero length", 2, 5},
		{"len=cap", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var s SyncSlice[int]
			s.InitSize(tt.length, tt.capacity)

			assert.Len(t, len(s.values), tt.length)
			assert.Equal(t, tt.capacity, cap(s.values))
		})
	}
}

func TestSyncSlice_AppendAndLen(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"append one", []int{42}, 1},
		{"append multiple", []int{1, 2, 3}, 3},
		{"append empty", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var s SyncSlice[int]
			for _, v := range tt.input {
				s.Append(v)
			}
			assert.Equal(t, tt.expected, s.Len())
		})
	}
}

func TestSyncSlice_GetAndSet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		initial     []int
		index       int
		newVal      int
		expectedOld int
		expectedNew int
	}{
		{"set first element", []int{1, 2, 3}, 0, 10, 1, 10},
		{"set middle element", []int{1, 2, 3}, 1, 20, 2, 20},
		{"set last element", []int{1, 2, 3}, 2, 30, 3, 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := SyncSlice[int]{values: append([]int{}, tt.initial...)}
			old := s.Get(tt.index)
			assert.Equal(t, tt.expectedOld, old)

			s.Set(tt.index, tt.newVal)
			got := s.Get(tt.index)
			assert.Equal(t, tt.expectedNew, got)
		})
	}
}

func TestSyncSlice_ConcurrentAppend(t *testing.T) {
	t.Parallel()

	const goroutines = 50
	const perGoroutine = 100

	var s SyncSlice[int]
	var wg sync.WaitGroup

	wg.Add(goroutines)
	for g := range goroutines {
		go func(base int) {
			defer wg.Done()
			for i := range perGoroutine {
				s.Append(base*perGoroutine + i)
			}
		}(g)
	}
	wg.Wait()

	assert.Equal(t, goroutines*perGoroutine, s.Len())
}
