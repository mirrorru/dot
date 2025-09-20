package dot

import (
	"sync"
)

type SyncStore[K comparable, V any] struct {
	storage map[K]V
	mx      sync.RWMutex
}

// Preallocate - init internal map with specified size
func (s *SyncStore[K, V]) Preallocate(mapSize int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	if s.storage == nil {
		s.storage = make(map[K]V, mapSize)
	}
}

func (s *SyncStore[K, V]) Put(key K, val V) {
	s.mx.Lock()
	defer s.mx.Unlock()
	if s.storage == nil {
		s.storage = make(map[K]V)
	}
	s.storage[key] = val
}

func (s *SyncStore[K, V]) GetCurrent(key K) (val V, founded bool) {
	if s.storage != nil {
		s.mx.RLock()
		defer s.mx.RUnlock()
		val, founded = s.storage[key]
	}

	return val, founded
}

func (s *SyncStore[K, V]) GetOrPut(key K, maker func() V) V {
	val, founded := s.GetCurrent(key)
	if founded {
		return val
	}

	s.mx.Lock()
	defer s.mx.Unlock()

	if s.storage == nil {
		s.storage = make(map[K]V)
	}

	val, founded = s.storage[key]
	if founded {
		return val
	}

	val = maker()
	s.storage[key] = val

	return val
}

func (s *SyncStore[K, V]) Del(key K) {
	s.mx.Lock()
	defer s.mx.Unlock()
	delete(s.storage, key) // works with nil
}

// ForEach вызывает handler для каждой пары ключ-значение
func (s *SyncStore[K, V]) ForEach(handler func(key K, value V)) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	for k, v := range s.storage {
		handler(k, v)
	}
}

func (s *SyncStore[K, V]) Iterator() <-chan struct {
	Key   K
	Value V
} {
	ch := make(chan struct {
		Key   K
		Value V
	})
	go func() {
		s.mx.RLock()
		defer s.mx.RUnlock()
		for k, v := range s.storage {
			ch <- struct {
				Key   K
				Value V
			}{k, v}
		}
		close(ch)
	}()
	return ch
}

// Seq реализует iter.Seq (итерация по значениям)
func (s *SyncStore[K, V]) Seq() func(yield func(V) bool) {
	return func(yield func(V) bool) {
		s.mx.RLock()
		defer s.mx.RUnlock()
		for _, v := range s.storage {
			if !yield(v) {
				break
			}
		}
	}
}

// Seq2 реализует iter.Seq2 (итерация по ключу и значению)
func (s *SyncStore[K, V]) Seq2() func(yield func(K, V) bool) {
	return func(yield func(K, V) bool) {
		s.mx.RLock()
		defer s.mx.RUnlock()
		for k, v := range s.storage {
			if !yield(k, v) {
				break
			}
		}
	}
}
