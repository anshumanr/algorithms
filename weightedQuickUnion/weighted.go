package weightedQuickUnion

type WeightedQuickUnionUF struct {
	size    int
	arr     []int
	arrSize []int

	sets int
}

func NewWeightedQuickUnionUF(n int) WeightedQuickUnionUF {
	uf := WeightedQuickUnionUF{
		size: n,
		sets: n,
	}

	uf.arr = make([]int, n)
	uf.arrSize = make([]int, n)

	for i := 0; i < n; i++ {
		uf.arr[i] = i
		uf.arrSize[i] = 1
	}

	return uf
}

func (w WeightedQuickUnionUF) Union(p, q int) {
	rootp := w.root(p)
	rootq := w.root(q)

	if rootp == rootq {
		return
	}

	if w.arrSize[rootp] >= w.arrSize[rootq] {
		w.arr[q] = rootp
		w.arrSize[p] += w.arrSize[q]
	} else {
		w.arr[p] = rootq
		w.arrSize[q] += w.arrSize[p]
	}

	w.sets--
}

func (w WeightedQuickUnionUF) Connected(p, q int) bool {
	return w.root(p) == w.root(q)
}

func (w WeightedQuickUnionUF) Count() int {
	return w.sets
}

func (w WeightedQuickUnionUF) root(node int) int {
	t := node
	for {
		if t == w.arr[t] {
			break
		}

		w.arr[t] = w.arr[w.arr[t]]
		t = w.arr[t]
	}

	return t
}
