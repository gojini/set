package set

type valIndex struct {
	index int
	value any
}

type orderedMap struct {
	capacity int
	entries  map[any]valIndex
	index    []any
}

func NewSet(capacity int) Set {
	return NewMap(capacity)
}

func NewMap(capacity int) Map {
	return &orderedMap{
		capacity: capacity,
		entries:  make(map[any]valIndex, capacity),
		index:    make([]any, 0, capacity),
	}
}

func (m *orderedMap) Capacity() int {
	return cap(m.index)
}

func (m *orderedMap) Size() int {
	return len(m.index)
}

func (m *orderedMap) Has(key any) bool {
	_, ok := m.entries[key]
	return ok
}

func (m *orderedMap) Add(key any) {
	m.AddKV(key, nil)
}

func (m *orderedMap) AddKV(key, value any) {
	if m.Has(key) {
		// we already have this key, just overwrite the value
		v := m.entries[key]
		v.value = value
		m.entries[key] = v

		return // dont add duplicates
	}

	// add the new entry
	s := len(m.entries)
	m.entries[key] = valIndex{
		index: s,
		value: value,
	}

	// add the key to the index for ordering
	m.index = append(m.index, key)
}

func (m *orderedMap) Remove(key any) {
	_, ok := m.entries[key]
	if !ok {
		return
	}

	// Delete the entry
	delete(m.entries, key)

	// remove the key from the index
	nIndex := make([]any, 0, m.capacity)
	for _, k := range m.index {
		if k == key {
			continue
		}

		nIndex = append(nIndex, k)
	}

	m.index = nIndex
}

func (m *orderedMap) Iterate() <-chan any {
	ch := make(chan any)
	go func() {
		for _, key := range m.index {
			ch <- key
		}
		close(ch)
	}()
	return ch
}

func (m *orderedMap) IterateValues() <-chan any {
	ch := make(chan any)
	go func() {
		for _, key := range m.index {
			ch <- m.entries[key].value // if the key is present value will be set
		}
		close(ch)
	}()
	return ch
}

func (m *orderedMap) Get(key any) any {
	v, ok := m.entries[key]
	if !ok {
		return nil
	}

	return v.value
}

func (m *orderedMap) Clear() {
	m.entries = make(map[any]valIndex, m.capacity)
	m.index = make([]any, 0, m.capacity)
}
