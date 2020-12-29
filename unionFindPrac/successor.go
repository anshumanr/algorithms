package unionFindPrac

import (
	wquf "alg/v1/weightedQuickUnion"
)

type SuccessorWD struct {
	uf *wquf.WeightedQuickUnionUF
}

func NewSuccessorWD(n int) *SuccessorWD {
	s := &SuccessorWD{
		uf: wquf.NewWeightedQuickUnionUF(n),
	}

	return s
}

func (s *SuccessorWD) Delete(p int) {
	s.uf.Union(p, p+1)
}

func (s *SuccessorWD) Successor(p int) int {
	return s.uf.Find(p)
}
