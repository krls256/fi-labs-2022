package stat

import "sort"

type iterator struct {
	keys      []int
	values    []int
	index     int
	isForward bool
}

func (it *iterator) Next() {
	it.index++
}

func (it *iterator) Valid() bool {
	return len(it.keys) > it.index
}

func (it *iterator) Key() int {
	return it.keys[it.index]
}

func (it *iterator) Value() int {
	return it.values[it.index]
}

func (it *iterator) Len() int {
	return len(it.keys)
}

func (it *iterator) Less(i, j int) bool {
	if it.isForward {
		return it.less(i, j)
	}
	return it.less(j, i)
}

func (it *iterator) Swap(i, j int) {
	it.values[i], it.values[j] = it.values[j], it.values[i]
	it.keys[i], it.keys[j] = it.keys[j], it.keys[i]
}

func (it *iterator) less(i, j int) bool {
	return it.values[i] < it.values[j]
}

func (it *iterator) sort() {
	sort.Sort(it)
}
