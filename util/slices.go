package util

type Integers interface {
	int | int64 | uint64
}

func Sum[K comparable, V Integers](slice []V) V {
	var sum V
	for _, v := range slice {
		sum += v
	}
	return sum
}
