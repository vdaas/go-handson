package main

import "fmt"

func sub(x, y int) int {
	return x - y
}

func main() {
	n := sub(20, 10)
	fmt.Println(n)
}
