package internal

import (
	"container/heap"
	"sync"
)

type Cache[T any] struct {
	mutex *sync.Mutex
	data  map[string]Item[T]

	evictionPolicy *lfuHeap
	size           int
}

func NewCache[T any]() *Cache[T] {
	return &Cache[T]{}
}

func (c *Cache[T]) Get(key string) (T, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, itemExists := c.data[key]
	if itemExists {
		item.frequency++
	}

	return item.value, itemExists
}

func (c *Cache[T]) Set(newData map[string]T) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for k, v := range newData {
		newItem := &Item[T]{
			key:       k,
			value:     v,
			frequency: 1,
		}

		c.data[k] = *newItem
		heap.Push(c.evictionPolicy, newItem)
	}
}

type lfuHeap []int

func (l lfuHeap) Len() int {
	return len(l)
}

func (l lfuHeap) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l lfuHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *lfuHeap) Push(x interface{}) {
	if x, xIsInt := x.(int); xIsInt {
		*l = append(*l, x)
	}
}

func (l *lfuHeap) Pop() interface{} {
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[:n-1]

	return x
}
