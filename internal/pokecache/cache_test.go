package pokecache_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hettiger/pokedexcli/internal/pokecache"
	"github.com/stretchr/testify/require"
)

const (
	fakeTTL time.Duration = 5 * time.Millisecond
	maxTTL  time.Duration = 1<<63 - 1
)

func TestNewCache(t *testing.T) {
	tests := []time.Duration{-1, 0}
	for _, ttl := range tests {
		t.Run(fmt.Sprintf("it panics on non-positive interval: %v", ttl), func(t *testing.T) {
			t.Parallel()

			require.Panics(t, func() {
				pokecache.NewCache(ttl)
			})
		})
	}
}

func TestCache_Add(t *testing.T) {
	t.Run("it adds entries to the cache", func(t *testing.T) {
		t.Parallel()

		c := pokecache.NewCache(maxTTL)
		t.Cleanup(c.Stop)

		c.Add("foo", []byte("bar"))
		c.Add("bar", []byte("baz"))

		foo, ok := c.Get("foo")
		require.True(t, ok)
		require.Equal(t, "bar", string(foo))

		bar, ok := c.Get("bar")
		require.True(t, ok)
		require.Equal(t, "baz", string(bar))
	})

	t.Run("it copies the value", func(t *testing.T) {
		t.Parallel()

		c := pokecache.NewCache(maxTTL)
		t.Cleanup(c.Stop)

		val := []byte("bar")
		c.Add("foo", val)
		val[0] = 'X'

		got, _ := c.Get("foo")
		require.Equal(t, "bar", string(got))
	})
}

func TestCache_Get(t *testing.T) {
	t.Run("it returns empty val and false on unknown key", func(t *testing.T) {
		t.Parallel()

		c := pokecache.NewCache(maxTTL)
		t.Cleanup(c.Stop)

		val, ok := c.Get("foo")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("it does not return reaped entries", func(t *testing.T) {
		t.Parallel()

		c := pokecache.NewCache(fakeTTL)
		t.Cleanup(c.Stop)

		c.Add("foo", []byte("bar"))

		val, ok := c.Get("foo")
		require.True(t, ok)
		require.Equal(t, []byte("bar"), val)

		require.Eventually(t, func() bool {
			_, ok := c.Get("foo")
			return !ok
		}, time.Second, fakeTTL)
	})

	t.Run("it returns a copy", func(t *testing.T) {
		t.Parallel()

		c := pokecache.NewCache(maxTTL)
		t.Cleanup(c.Stop)

		c.Add("foo", []byte("bar"))
		got, _ := c.Get("foo")
		got[0] = 'X'

		again, _ := c.Get("foo")
		require.Equal(t, "bar", string(again))
	})
}

func TestCache_Stop(t *testing.T) {
	t.Run("it is idempotent", func(t *testing.T) {
		t.Parallel()

		c := pokecache.NewCache(time.Second)

		require.NotPanics(t, func() {
			c.Stop()
			c.Stop()
		})
	})
}
