package main

import "fmt"

func main() {
	counter := make(map[string]int)

	for i := 1; i < 100; i++ {
		switch {
		case i%15 == 0:
			counter["FizzBuzz"]++
		case i%5 == 0:
			counter["Buzz"]++
		case i%3 == 0:
			counter["Fizz"]++
		default:
			counter["others"]++
		}
	}

	for k, v := range counter {
		fmt.Printf("%s %d\n", k, v)
	}
}
