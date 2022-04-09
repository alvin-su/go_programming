// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

/* func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
} */

//练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
func main() {

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
}
