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

func AllValues[K comparable, V Integers](slice []V, v V) bool {
	for _, i := range slice {
		if i != v {
			return false
		}
	}
	return true
}

func CalculateDiffs[K comparable, V Integers](h []V) (ans []V) {
	for i := 1; i < len(h); i++ {
		ans = append(ans, h[i]-h[i-1])
	}
	return ans
}
