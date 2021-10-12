package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	n := add(10, 20)
	fmt.Println(n)
}
