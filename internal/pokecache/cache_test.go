package pokecache_test

import (
	"testing"
	"time"

	"github.com/hettiger/pokedexcli/internal/pokecache"
	"github.com/stretchr/testify/require"
)

const (
	fakeTTL time.Duration = 5 * time.Millisecond
	maxTTL  time.Duration = 1<<63 - 1
)

func TestCache_Add(t *testing.T) {
	t.Run("it adds entries to the cache", func(t *testing.T) {
		t.Parallel()

		c := pokecache.NewCache(maxTTL)

		c.Add("foo", []byte("bar"))
		c.Add("bar", []byte("baz"))

		foo, ok := c.Get("foo")
		require.True(t, ok)
		require.Equal(t, "bar", string(foo))

		bar, ok := c.Get("bar")
		require.True(t, ok)
		require.Equal(t, "baz", string(bar))
	})
}

func TestCache_Get(t *testing.T) {
	t.Run("it returns empty val and false on unknown key", func(t *testing.T) {
		t.Parallel()

		c := pokecache.NewCache(maxTTL)

		val, ok := c.Get("foo")
		require.False(t, ok)
		require.Equal(t, []byte{}, val)
	})

	t.Run("it does not return reaped entries", func(t *testing.T) {
		t.Parallel()

		c := pokecache.NewCache(fakeTTL)

		c.Add("foo", []byte("bar"))

		val, ok := c.Get("foo")
		require.True(t, ok)
		require.Equal(t, []byte("bar"), val)

		time.Sleep(2 * fakeTTL)

		val, ok = c.Get("foo")
		require.False(t, ok)
		require.Equal(t, []byte{}, val)
	})
}
