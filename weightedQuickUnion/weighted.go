package weightedQuickUnion

import (
	"fmt"
)

type WeightedQuickUnionUF struct {
	size    int
	arr     []int
	arrSize []int
	sets    int
	setMap  map[int][]int
}

func NewWeightedQuickUnionUF(n int) *WeightedQuickUnionUF {
	uf := &WeightedQuickUnionUF{
		size: n,
		sets: n,
	}

	uf.arr = make([]int, n)
	uf.arrSize = make([]int, n)
	uf.setMap = make(map[int][]int, n)

	for i := 0; i < n; i++ {
		uf.arr[i] = i
		uf.arrSize[i] = 1
		uf.setMap[i] = []int{i} //append(uf.setMap[i], i)
	}

	return uf
}

func (w *WeightedQuickUnionUF) Union(p, q int) {
	rootp := w.root(p)
	rootq := w.root(q)

	if rootp == rootq {
		return
	}

	// fmt.Println("|Items: ", p, q, "|Arr: ", w.arr[p], w.arr[q])
	// fmt.Println("|Roots: ", rootp, rootq, "|ArrSize: ", w.arrSize[rootp], w.arrSize[rootq])
	// fmt.Println("|Sets: ", w.setMap)

	//fmt.Println("1: ", w.setMap[rootp], w.setMap[rootq], w.arr)
	if w.arrSize[rootp] >= w.arrSize[rootq] {
		w.arr[rootq] = rootp
		w.arrSize[rootp] += w.arrSize[rootq]
		w.setMap[rootp] = append(w.setMap[rootp], w.setMap[rootq]...)
		delete(w.setMap, rootq)
	} else {
		w.arr[rootp] = rootq
		w.arrSize[rootq] += w.arrSize[rootp]
		w.setMap[rootq] = append(w.setMap[rootq], w.setMap[rootp]...)
		delete(w.setMap, rootp)
	}
	//fmt.Println("2: ", w.setMap[rootp], w.setMap[rootq], w.arr)

	w.sets--
}

func (w *WeightedQuickUnionUF) Connected(p, q int) bool {
	return w.root(p) == w.root(q)
}

func (w *WeightedQuickUnionUF) Count() int {
	return w.sets
}

func (w *WeightedQuickUnionUF) LogData() string {
	str := fmt.Sprintln("|Arr: ", w.arr, "|ArrSize", w.arrSize, "\n|Sets:", w.setMap)
	return str
}

func (w *WeightedQuickUnionUF) root(node int) int {
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
