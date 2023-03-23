package main

import "fmt"

// firstチャネルに 1~9 の値を送信
func producer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

// firstチャネルから受信した値を * 2 してsecondチャネルに送信
func multi2(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

// secondチャネルから受信した値を * 4 してthirdチャネルに送信
func multi4(second, third chan int) {
	// main関数内のforに対して、チャネルの送信の終了を知らせる
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multi2(first, second)
	go multi4(second, third)

	// third のチャネルが送信終了するまで受信し続ける
	for result := range third {
		// thirdチャネルから受信した値を出力する
		fmt.Println(result)
	}
}
