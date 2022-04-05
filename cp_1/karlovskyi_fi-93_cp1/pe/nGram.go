package pe

import "math"

type NGramStat struct {
	Store map[string]int
	Total int
	N     int
}

func newStat(n int) *NGramStat {
	return &NGramStat{
		Store: map[string]int{},
		Total: 0,
		N:     n,
	}
}
func (stat *NGramStat) AddNGram(str string) {
	stat.Store[str]++
	stat.Total++
}
func (stat *NGramStat) ForEachInOrder(fn func(string, int)) {
	keys := make([]string, len(stat.Store))
	values := make([]int, len(stat.Store))
	i := 0
	for k, v := range stat.Store {
		keys[i] = k
		values[i] = v
		i++
	}
	var quickSort func(k []string, v []int, l, h int)
	quickSort = func(k []string, v []int, l, h int) {
		if l < h {
			pivot := v[h]
			i := l - 1

			for j := l; j < h; j++ {
				if v[j] < pivot {
					i++
					v[i], v[j] = v[j], v[i]
					k[i], k[j] = k[j], k[i]
				}
			}
			v[i+1], v[h] = v[h], v[i+1]
			k[i+1], k[h] = k[h], k[i+1]
			pi := i + 1
			quickSort(k, v, l, pi-1)
			quickSort(k, v, pi+1, h)
		}
	}

	quickSort(keys, values, 0, len(keys)-1)

	for i := 0; i < len(keys); i++ {
		fn(keys[i], values[i])
	}
}

func (stat *NGramStat) Entropy() float64 {
	ent := 0.0
	totFloat := float64(stat.Total)
	for _, val := range stat.Store {
		valFloat := float64(val)
		ent -= (valFloat / totFloat) * math.Log2(valFloat/totFloat)
	}
	return ent / float64(stat.N)
}
