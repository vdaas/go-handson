package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	go func() { // goroutineを使い別スレッドで実行
		for i := 0; i < 20; i++ {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 300)
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(i * 100)
		time.Sleep(time.Millisecond * 600)
	}
}

// END OMIT
