package stat

import "sort"

type Iterator struct {
	keys      []int
	values    []int
	index     int
	isForward bool
}

func (it *Iterator) Next() {
	it.index++
}

func (it *Iterator) Valid() bool {
	return len(it.keys) > it.index
}

func (it *Iterator) Key() int {
	return it.keys[it.index]
}

func (it *Iterator) Value() int {
	return it.values[it.index]
}

func (it *Iterator) Len() int {
	return len(it.keys)
}

func (it *Iterator) Less(i, j int) bool {
	if it.isForward {
		return it.less(i, j)
	}
	return it.less(j, i)
}

func (it *Iterator) Swap(i, j int) {
	it.values[i], it.values[j] = it.values[j], it.values[i]
	it.keys[i], it.keys[j] = it.keys[j], it.keys[i]
}

func (it *Iterator) less(i, j int) bool {
	return it.values[i] < it.values[j]
}

func (it *Iterator) sort() {
	sort.Sort(it)
}
