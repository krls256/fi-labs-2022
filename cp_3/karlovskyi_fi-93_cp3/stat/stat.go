package stat

type Stat struct {
	count map[int]int
	total int
}

func (stat *Stat) init() {
	if stat.count == nil {
		stat.count = map[int]int{}
	}
}

func (stat *Stat) GetAllCounts() map[int]int {
	stat.init()
	return stat.count
}

func (stat *Stat) GetCount(key int) int {
	stat.init()
	return stat.count[key]
}

func (stat *Stat) GetProbability(key int) float64 {
	stat.init()
	if stat.total == 0 {
		return 0
	}
	return float64(stat.count[key]) / float64(stat.total)
}

func (stat *Stat) GetTotal() int {
	return stat.total
}

func (stat *Stat) inc(key int) {
	stat.count[key]++
	stat.total++
}

func (stat *Stat) Inc(key int) {
	stat.init()
	stat.inc(key)
}

func (stat *Stat) Reset() {
	stat.init()
	stat.total = 0
	stat.count = map[int]int{}
}

func (stat *Stat) Append(arr []int) {
	stat.init()
	for _, i := range arr {
		stat.inc(i)
	}
}

func (stat *Stat) mapToDoubleArray() ([]int, []int) {
	keys := make([]int, 0, len(stat.count))
	values := make([]int, 0, len(stat.count))
	for k, v := range stat.count {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
func (stat *Stat) pureIterator() *Iterator {
	keys, values := stat.mapToDoubleArray()
	return &Iterator{keys: keys, values: values}
}
func (stat *Stat) ForwardIterator() *Iterator {
	it := stat.pureIterator()
	it.isForward = true
	it.sort()
	return it
}

func (stat *Stat) BackIterator() *Iterator {
	it := stat.pureIterator()
	it.isForward = false
	it.sort()
	return it
}
