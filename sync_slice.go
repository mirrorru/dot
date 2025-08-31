package dot

import "sync"

type SyncSlice[T any] struct {
	mx     sync.RWMutex
	values []T
}

func (s *SyncSlice[T]) InitSize(length, capacity int) {
	if s.values == nil {
		s.values = make([]T, length, capacity)
	}
}

func (s *SyncSlice[T]) Values() []T {
	return s.values
}

func (s *SyncSlice[T]) Append(val T) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.values = append(s.values, val)
}

func (s *SyncSlice[T]) Get(index int) (val T) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	return s.values[index]
}

func (s *SyncSlice[T]) Set(index int, val T) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.values[index] = val
}
