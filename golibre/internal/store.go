package internal

import (
	"context"
	"sync"
)

type Cacher[K comparable, T any] interface {
	Get(ctx context.Context, key K) (Item[T], bool)
	Set(newData map[K]Item[T])
}

func NewStore[Key comparable, T any]() *Store[Key, T] {
	return &Store[Key, T]{
		mutex: &sync.RWMutex{},
		items: map[Key]Item[T]{},
	}
}

type Store[Key comparable, T any] struct {
	mutex *sync.RWMutex
	items map[Key]Item[T]
}

func (s *Store[Key, T]) Set(newData map[Key]Item[T]) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.items = newData
}

func (s *Store[Key, T]) Get(ctx context.Context, key Key) (item Item[T], exists bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	item, exists = s.items[key]

	return item, exists
}

func (s *Store[Key, T]) GetMultiple(ctx context.Context, keys []Key) (map[Key]Item[T], []Key) {
	results := make(map[Key]Item[T])
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

type Item[T any] struct {
	key       string
	value     T
	frequency int
	index     int
}
