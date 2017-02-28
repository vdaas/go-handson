package main

import "fmt"

func main() {

	sample()

}

// START OMIT
func sample() {

	fmt.Println("start") //開始処理

	defer fmt.Println("defer") // メソッドを抜ける時に必ず処理したい

	fmt.Println("fin") //最終処理
}

// END OMIT
