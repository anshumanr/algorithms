package weightedQuickUnion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeighted1(t *testing.T) {
	w := NewWeightedQuickUnionUF(10)

	got := w.Connected(1, 2)
	assert.False(t, got)

	w.Union(0, 2)
	t.Log(w.LogData())

	got = w.Connected(0, 2)
	assert.True(t, got)

	c := w.Count()
	assert.Equal(t, 9, c)

	w.Union(1, 3)
	w.Union(5, 6)
	w.Union(6, 0)
	t.Log(w.LogData())

	c = w.Count()
	assert.Equal(t, 6, c)
	got = w.Connected(5, 2)
	assert.True(t, got)
	got = w.Connected(6, 2)
	assert.True(t, got)
	got = w.Connected(2, 5)
	assert.True(t, got)

	w.Union(2, 5)
	c = w.Count()
	assert.Equal(t, 6, c)
}

func TestWeighted2(t *testing.T) {
	w := NewWeightedQuickUnionUF(10)
	t.Logf(w.LogData())

	w.Union(0, 1)
	w.Union(2, 4)
	w.Union(3, 5)
	w.Union(1, 6)
	w.Union(3, 4)

	t.Logf(w.LogData())
	got := w.Count()
	assert.Equal(t, 5, got)

	w.Union(7, 9)
	w.Union(8, 1)

	t.Logf(w.LogData())
	got = w.Count()
	assert.Equal(t, 3, got)

	w.Union(7, 8)
	w.Union(0, 4)

	t.Logf(w.LogData())
	got = w.Count()
	assert.Equal(t, 1, got)
}
