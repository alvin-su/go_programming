package main

import "fmt"

func main() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr[0])
	fmt.Println(arr[len(arr)-1])
	// Print the indices and elements.
	for i, v := range arr {
		fmt.Printf("%d %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range arr {
		fmt.Printf("%d\n", v)
	}

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)
