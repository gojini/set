package tests

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"gojini.dev/set"
)

func TestConcurrentMap(t *testing.T) {
	t.Parallel()

	require := require.New(t)

	m := set.NewConcurrentMap(1000)
	require.NotNil(m)
	require.Equal(0, m.Size())
	require.Equal(1000, m.Capacity())

	waiter := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waiter.Add(1)
		go func() {
			defer waiter.Done()
			base := i * 100
			m.AddKV(base, base)

			require.Equal(base, m.Get(base))
			require.True(m.Has(base))
			m.Remove(base)
			require.False(m.Has(base))

			for i := base; i < base+100; i++ {
				m.AddKV(i, i)

				require.Equal(i, m.Get(i))
				require.True(m.Has(i))
			}
		}()
	}

	waiter.Wait()

	for k := range m.Iterate() {
		require.True(m.Has(k))
	}

	for v := range m.IterateValues() {
		require.True(m.Has(v))
	}

	m.Clear()
	require.Equal(0, m.Size())
	require.Equal(1000, m.Capacity())
}

func TestConcurrentSet(t *testing.T) {
	t.Parallel()

	require := require.New(t)

	s := set.NewConcurrentSet(1000)
	require.NotNil(s)
	require.Equal(0, s.Size())
	require.Equal(1000, s.Capacity())

	waiter := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waiter.Add(1)
		go func() {
			defer waiter.Done()
			base := i * 100
			s.Add(base)

			require.True(s.Has(base))
			s.Remove(base)
			require.False(s.Has(base))

			for i := base; i < base+100; i++ {
				s.Add(i)

				require.True(s.Has(i))
			}
		}()
	}

	waiter.Wait()

	for k := range s.Iterate() {
		require.True(s.Has(k))
	}

	s.Clear()
	require.Equal(0, s.Size())
	require.Equal(1000, s.Capacity())
}
