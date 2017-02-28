package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	c := make(chan int, 10) // int 型の長さ10の buffer付きchannelを作る
	go func() {             // goroutineを使い別スレッドで実行
		for i := 0; i < 20; i++ {
			c <- i // channel に int型の値を送る
			fmt.Println("send")
		}
	}()
	for i := 0; i < 20; i++ {
		i := <-c // channelに値が入ってきたら i に代入
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}

// END OMIT
