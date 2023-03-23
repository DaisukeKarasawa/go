### close()を使用してチャネルの受け渡しの終了を知らせる

#### チャネル

並列処理や通常処理の間で、データのやり取りをする。
```
func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine1(s, c)
	x := <-c
	fmt.Println(x)
}
```

**プロセス**

*１：スライスとチャネルの作成*
*２：ゴルーチンにスライスとチャネルを渡す*
*３：ゴルーチンで合計処理をする*
*４：チャネルに sum を送る*
*５：c が受信して値を取り出す*
*６：取り出しが完了したら、次の処理へ*

**ポイント**

*1：プロセス４で受信をするまで処理が待機する。*

sync.Waitを使用せずにゴルーチンの処理を待機させることが出来る。
```
x := <-ch // 待機
```

*２：makeで作成したチャネルは、キューのようになる。（アンバッファ）*

同じチャネルを使用してデータのやり取りを複数した場合、送られてきたものから順番にチャネルに入る。
```
c := make(chan int) // 値１, 値２

x := <-c            // 値１を取り出す
y := <-c            // 値２を取り出す
```

*３：バッファの数をしてして、チャネルの受信量を指定する。（バッファ）*

バッファの数以上のデータが入るとエラーになる
```
ch := make(chan int, 2)
ch <- 100
ch <- 200
ch <- 300   // fatal error: all goroutines are asleep - deadlock!
```
その前にチャネルから値を取り出せば、問題はない。
```
ch := make(chan int, 2)
ch <- 100
ch <- 200
x := <-ch
ch <- 300
```

#### forでチャネルから値を取り出す

範囲節のforを使用して、チャネルからデータを取り出す。
```
ch := make(chan int, 2)
ch <- 100
ch <- 200

for c := range ch {
  fmt.Println(c)   // fatal error: all goroutines are asleep - deadlock!
}
```

**問題**

１つ目、２つ目のチャネルを取り出して、３つ目のチャネルを取り出そうとしても取り出すところが無い。

それでも range で取り出し続けようとしている。

解決方法：*close()*を使用してチャネルをcloseする必要がある。
```
ch := make(chan int, 2)
ch <- 100
ch <- 200
close(ch)

for c := range ch {
  fmt.Println(c)  // 100 200
}
```

##### 余談

引数にチャネルを取る場合、分かりやすいように受信するチャネルなのか、送信するチャネルなのかを明示的に指定することが出来る。
```
func multi2(first <-chan int, second chan<- int) {
  ...処理
}
```