package main

import "fmt"

// START OMIT
func main() {

	c := make(chan int, 1)
	c <- 1

	select {
	case i := <-c:
		fmt.Printf("received : %d\n", i)
	default:
		fmt.Println("no value")
	}
}

// END OMIT
