package unionFindPrac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessor(t *testing.T) {
	s := NewSuccessorWD(10)
	t.Logf(s.uf.LogData())

	s.Delete(6)
	t.Logf(s.uf.LogData())
	want := 7
	got := s.Successor(6)
	assert.Equal(t, want, got)

	s.Delete(6)
	t.Logf(s.uf.LogData())
	want = 7
	got = s.Successor(6)
	assert.Equal(t, want, got)

	s.Delete(7)
	t.Logf(s.uf.LogData())
	want = 8
	got = s.Successor(7)
	assert.Equal(t, want, got)
}
