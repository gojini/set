package set

import "sync"

type concurrentMap struct {
	lock    *sync.RWMutex
	entries Map
}

func NewConcurrentMap(capacity int) Map {
	return &concurrentMap{
		lock:    &sync.RWMutex{},
		entries: NewMap(capacity),
	}
}

func NewConcurrentSet(capacity int) Set {
	return NewConcurrentMap(capacity)
}

func (m *concurrentMap) Capacity() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.entries.Capacity()
}

func (m *concurrentMap) Size() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.entries.Size()
}

func (m *concurrentMap) Has(key any) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.entries.Has(key)
}

func (m *concurrentMap) Add(key any) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.entries.Add(key)
}

func (m *concurrentMap) AddKV(key, value any) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.entries.AddKV(key, value)
}

func (m *concurrentMap) Get(key any) any {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.entries.Get(key)
}

func (m *concurrentMap) Iterate() <-chan any {
	keys := m.copyEntries(true)

	return makeChan(keys)
}

func (m *concurrentMap) IterateValues() <-chan any {
	values := m.copyEntries(false)

	return makeChan(values)
}

func (m *concurrentMap) copyEntries(keys bool) []any {
	// snapshot all the values. this o(N) operation
	// for both memory and time
	m.lock.RLock()
	defer m.lock.RUnlock()

	items := make([]any, 0, m.entries.Size())
	if keys {
		for k := range m.entries.Iterate() {
			items = append(items, k)
		}
	} else {
		for v := range m.entries.IterateValues() {
			items = append(items, v)
		}
	}

	return items
}

func makeChan(items []any) <-chan any {
	ch := make(chan any)
	go func() {
		for _, v := range items {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func (m *concurrentMap) Remove(key any) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.entries.Remove(key)
}

func (m *concurrentMap) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.entries.Clear()
}
