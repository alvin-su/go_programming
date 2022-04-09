package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["a"] = 1
	x["b"] = 2
	y := map[string]int{"a": 1, "b": 2}
	isEqual := equal(x, y)
	fmt.Println(isEqual)
	isEqual2 := equal(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Println(isEqual2)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
