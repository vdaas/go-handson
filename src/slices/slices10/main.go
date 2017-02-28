package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	fmt.Printf("a: %v\n", a)
	b := a[0:5]
	c := make([]int, len(a))
	copy(c, a)
	for i := 0; i < len(b); i++ {
		b[i] = b[i] * -1
	}
	fmt.Printf("b: %v\n", b)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("c: %v\n", c)
}
