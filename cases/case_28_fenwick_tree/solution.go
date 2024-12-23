package case_28_fenwick_tree

type Fenwick struct {
	fenw []int
}

func NewFenwick(size int) *Fenwick {
	return &Fenwick{fenw: make([]int, size+1)}
}

func (f *Fenwick) Update(i, delta int) {
	for i < len(f.fenw) {
		f.fenw[i] += delta
		i += i & (-i)
	}
}

func (f *Fenwick) PrefixSum(i int) int {
	sum := 0
	for i > 0 {
		sum += f.fenw[i]
		i -= i & (-i)
	}
	return sum
}
