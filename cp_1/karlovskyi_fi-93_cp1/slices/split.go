// Package slices implements simple function to split any slice by separator
package slices

func Split[T comparable](slice []T, sep T) [][]T {
	res := make([][]T, 0)
	start := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == sep && start != i {
			res = append(res, slice[start:i+1])
			start = i + 1
		}
	}
	if start != len(slice) {
		res = append(res, slice[start:])
	}
	return res
}
