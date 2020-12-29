package unionFindPrac

import (
	wquf "alg/v1/weightedQuickUnion"
)

type SocCon struct {
	uf *wquf.WeightedQuickUnionUF
}

func NewSocCon(n int) *SocCon {
	s := &SocCon{
		uf: wquf.NewWeightedQuickUnionUF(n),
	}

	return s
}

func (s *SocCon) AllConnected() bool {
	return s.uf.Count() == 1
}
