package assign1

import (
	wquf "alg/v1/weightedQuickUnion"
	"fmt"
)

type Percolation struct {
	size      int
	max       int
	first     int
	last      int
	openCount int
	open      []bool
	uf        *wquf.WeightedQuickUnionUF
}

func NewPercolation(s int) *Percolation {
	m := s*s + 2
	p := &Percolation{
		size:      s,
		max:       m,
		first:     m - 2,
		last:      m - 1,
		openCount: 0,
		open:      make([]bool, m),
		uf:        wquf.NewWeightedQuickUnionUF(m),
	}

	return p
}

func (p *Percolation) Open(row, col int) error {
	//input index starts from 1, arr index starts from 0
	row--
	col--

	if (row < 0 || row > p.size-1) || (col < 0 || col > p.size-1) {
		return fmt.Errorf("Input out of bounds: %d, %d", row, col)
	}

	if p.isOpen(row, col) {
		return nil
	}

	index := p.getIndex(row, col)
	//fmt.Println("Open: ", row, col, index)

	if row == 0 {
		p.uf.Union(p.first, index)
	}

	if row == p.size-1 {
		p.uf.Union(p.last, index)
		//p.uf2.Union(p.last, index)
	}

	//upper
	if row > 0 && p.isOpen(row-1, col) {
		p.uf.Union(index, p.getIndex(row-1, col))
	}

	//lower
	if row < p.size-1 && p.isOpen(row+1, col) {
		p.uf.Union(index, p.getIndex(row+1, col))
	}

	//left
	if col > 0 && p.isOpen(row, col-1) {
		p.uf.Union(index, p.getIndex(row, col-1))
	}

	//right
	if col < p.size-1 && p.isOpen(row, col+1) {
		p.uf.Union(index, p.getIndex(row, col+1))
	}

	p.open[index] = true
	p.openCount++

	return nil
}

func (p *Percolation) IsOpen(row, col int) bool {
	//input index starts from 1, arr index starts from 0
	row--
	col--

	return p.isOpen(row, col)
}

func (p *Percolation) Percolates() bool {
	return p.uf.Connected(p.first, p.last)
}

func (p *Percolation) IsFull(row, col int) bool {
	row--
	col--

	return p.isFull(row, col)
}

func (p *Percolation) NumberOfOpenSites() int {
	return p.openCount
}

func (p *Percolation) LogData() string {
	return fmt.Sprintf("%v", p.open)
}

func (p *Percolation) isFull(row, col int) bool {
	index := p.getIndex(row, col)
	f := p.uf.Find(index)
	if f >= p.size*(p.size-1) {
		return true
	}

	return false
}

func (p *Percolation) isOpen(row, col int) bool {
	if (row < 0 || row > p.size-1) || (col < 0 || col > p.size-1) {
		fmt.Println("Input out of bounds:", row, col)
		return false
	}

	return p.open[p.getIndex(row, col)]
}

func (p *Percolation) getIndex(row, col int) int {
	i := (row * p.size) + col
	//fmt.Println("Index: ", i)
	return i
}
