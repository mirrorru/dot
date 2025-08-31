package dot

import (
	"iter"
	"sync"
)

type SyncSlice[T any] struct {
	mx     sync.Mutex
	values []T
}

func (s *SyncSlice[T]) InitSize(length, capacity int) {
	if s.values == nil {
		s.values = make([]T, length, capacity)
	}
}

func (s *SyncSlice[T]) Len() int {
	s.mx.Lock()
	defer s.mx.Unlock()
	return len(s.values)
}

func (s *SyncSlice[T]) Append(val T) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.values = append(s.values, val)
}

func (s *SyncSlice[T]) Get(index int) (val T) {
	s.mx.Lock()
	defer s.mx.Unlock()

	return s.values[index]
}

func (s *SyncSlice[T]) Set(index int, val T) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.values[index] = val
}

func (s *SyncSlice[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		s.mx.Lock()
		snapshot := make([]T, len(s.values))
		copy(snapshot, s.values)
		s.mx.Unlock()

		for _, v := range snapshot {
			if !yield(v) {
				return
			}
		}
	}
}

func (s *SyncSlice[T]) Seq2() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		s.mx.Lock()
		snapshot := make([]T, len(s.values))
		copy(snapshot, s.values)
		s.mx.Unlock()

		for i, v := range snapshot {
			if !yield(i, v) {
				return
			}
		}
	}
}
