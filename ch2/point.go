package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	var pp = f()
	fmt.Println(*pp)
	p1 := 2
	var p2 = incr(&p1)
	fmt.Println(p2)

	c := new(int)
	fmt.Println(*c)
	*c = 2
	fmt.Println(*c)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
