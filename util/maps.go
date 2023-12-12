package util

import "fmt"

func PrintStringMap(m map[string]string) {
	var maxLenKey int
	for k := range m {
		if len(k) > maxLenKey {
			maxLenKey = len(k)
		}
	}

	for k, v := range m {
		fmt.Printf("%*s: %s", maxLenKey, k, v)
	}
}

func PrintIntMap(m map[int]int) {
	for k, v := range m {
		fmt.Printf("[%d]: %d\n", k, v)
	}
}
