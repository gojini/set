package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gojini.dev/set"
)

func TestOrderedMap(t *testing.T) {
	t.Parallel()

	require := require.New(t)

	m := set.NewMap(10)
	require.NotNil(m)
	require.Equal(0, m.Size())
	require.Equal(10, m.Capacity())

	m.AddKV(1, 1)
	require.Equal(1, m.Size())
	require.Equal(1, m.Get(1))
	require.True(m.Has(1))
	require.False(m.Has(2))
	require.Equal(1, m.Get(1))
	require.Equal(nil, m.Get(2))

	// Add duplicate key
	m.AddKV(1, 2)
	require.Equal(1, m.Size())
	require.Equal(2, m.Get(1))

	m.AddKV(2, 2)
	require.Equal(2, m.Size())
	require.Equal(2, m.Get(2))
	require.True(m.Has(1))
	require.True(m.Has(2))

	// remove a key
	m.Remove(1)
	require.Equal(1, m.Size())
	require.False(m.Has(1))
	require.Equal(nil, m.Get(1))

	// remove a non-existent key
	m.Remove(1)
	require.Equal(1, m.Size())

	m.Clear()
	require.Equal(0, m.Size())

	m = set.NewMap(100)

	for i := 0; i < 100; i++ {
		m.AddKV(i, i)
	}
	require.Equal(100, m.Size())

	i := 0
	for k := range m.Iterate() {
		require.Equal(i, k)
		i++
	}

	i = 0
	for v := range m.IterateValues() {
		require.Equal(i, v)
		i++
	}
}

func TestOrderedSet(t *testing.T) {
	t.Parallel()

	require := require.New(t)
	s := set.NewSet(10)
	require.NotNil(s)
	require.Equal(0, s.Size())
	require.Equal(10, s.Capacity())

	s.Add(1)
	require.Equal(1, s.Size())
	require.True(s.Has(1))
	require.False(s.Has(2))

	// Add duplicate key
	s.Add(1)
	require.Equal(1, s.Size())

	s.Add(2)
	require.Equal(2, s.Size())
	require.True(s.Has(1))
	require.True(s.Has(2))

	// remove a key
	s.Remove(1)
	require.Equal(1, s.Size())
	require.False(s.Has(1))

	// remove a non-existent key
	s.Remove(1)
	require.Equal(1, s.Size())

	s.Clear()
	require.Equal(0, s.Size())

	s = set.NewSet(100)

	for i := 0; i < 100; i++ {
		s.Add(i)
	}
	require.Equal(100, s.Size())

	i := 0
	for k := range s.Iterate() {
		require.Equal(i, k)
		i++
	}
}
