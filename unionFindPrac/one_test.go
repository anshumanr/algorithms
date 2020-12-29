package unionFindPrac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSocCon(t *testing.T) {
	s := NewSocCon(10)
	t.Logf(s.uf.LogData())

	s.uf.Union(0, 1)
	s.uf.Union(2, 4)
	s.uf.Union(3, 5)
	s.uf.Union(1, 6)
	s.uf.Union(3, 4)

	t.Logf(s.uf.LogData())
	got := s.uf.Count()
	assert.Equal(t, 5, got)

	s.uf.Union(7, 9)
	s.uf.Union(8, 1)

	t.Logf(s.uf.LogData())
	got = s.uf.Count()
	assert.Equal(t, 3, got)

	s.uf.Union(7, 8)
	s.uf.Union(0, 4)

	t.Logf(s.uf.LogData())
	got = s.uf.Count()
	assert.Equal(t, 1, got)

	want := s.AllConnected()
	assert.True(t, want)
}
