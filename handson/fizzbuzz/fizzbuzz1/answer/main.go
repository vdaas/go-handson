package main

import "fmt"

// FizzBuzz
// 1以上、100未満の整数Nのうち
// ケース1: Nが3で割り切れて かつ 5で割り切れないときは "Fizz"
// ケース2: Nが5で割り切れて かつ 3で割り切れないときは "Buzz"
// ケース3: Nが3と5の両方で割り切れるときは "FizzBuzz"
// ケース4: Nが3と5のそのどちらでも割り切れない場合は N
// を標準出力に出力してください
func main() {
	for i := 1; i < 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}
}
