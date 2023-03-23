package main

import (
	"fmt"
	"sync"
)

// firstチャネルに 1~9 の値を送信
func producer(first chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		first <- i
	}
	close(first)
}

// firstチャネルから受信した値を * 2 してsecondチャネルに送信
func multi2(first <-chan int, second chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range first {
		second <- i * 2
	}
	close(second)
}

// secondチャネルから受信した値を * 4 してthirdチャネルに送信
func multi4(second, third chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range second {
		third <- i * 4
	}
	close(third)
}

// thirdチャネルから受信した値を出力する
func outputResult(third chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range third {
		fmt.Println(result)
	}
}

func main() {
	var wg sync.WaitGroup
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	// 並列処理が４つ動いていることを知らせる
	wg.Add(4)

	go producer(first, &wg)
	go multi2(first, second, &wg)
	go multi4(second, third, &wg)
	go outputResult(third, &wg)

	// ４つの並列処理が終了したことが知らされるまで、待機する
	wg.Wait()
}
