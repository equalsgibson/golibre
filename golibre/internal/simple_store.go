package internal

import (
	"context"
	"sync"
)

func NewSimpleStore[Key comparable, T any]() *SimpleStore[Key, T] {
	return &SimpleStore[Key, T]{
		mutex: &sync.RWMutex{},
		items: map[Key]T{},
	}
}

type SimpleStore[Key comparable, T any] struct {
	mutex *sync.RWMutex
	items map[Key]T
}

func (s *SimpleStore[Key, T]) Set(newData map[Key]T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for key, data := range newData {
		s.items[key] = data
	}
}

func (s *SimpleStore[Key, T]) Evict(key Key) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.items, key)
}

func (s *SimpleStore[Key, T]) Get(ctx context.Context, key Key) (item T, exists bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	item, exists = s.items[key]

	return item, exists
}

func (s *SimpleStore[Key, T]) GetMultiple(ctx context.Context, keys []Key) (map[Key]T, []Key) {
	results := make(map[Key]T)
	misses := []Key{}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, key := range keys {
		item, itemExists := s.items[key]
		if !itemExists {
			misses = append(misses, key)

			continue
		}

		results[key] = item
	}

	return results, misses
}

func (s *SimpleStore[Key, T]) GetAll(ctx context.Context) map[Key]T {
	results := make(map[Key]T)

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for key, value := range s.items {
		results[key] = value
	}

	return results
}
