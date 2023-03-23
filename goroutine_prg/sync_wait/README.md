## sync.WaitGroupを使用して並列処理の処理の終了を待つ

### sync.WaitGroup

動いているゴルーチンの処理の終了を待ってから、次の処理へ移行することが出来る。
```
func goroutine(wg *sync.WaitGroup) {
  ...処理
  wg.Done()             // ゴルーチン内で処理が終了したことを知らせる
}

func nomal() {
  ...処理
}

func main() {
  var wg sync.WaitGroup // 宣言
  wg.Add(1)             // １つの並列処理があることを伝える

  go goroutine(&wg)     // ゴルーチンには、宣言したアドレスを渡す
  nomal()
  wg.Wait()             // .Done()が宣言されるまで待機する
}
```

**プロセス**

*１： var wg sync.WaitGroup で宣言する。*

*２： wg.Add(1) で並列処理を１つ追加する。*

*３： wg.Wait() で追加された並列処理１つが終了されるのを待つ。*

*４： wg.Done() で終了を知らせる。*

### .Wait()を置く場所

**例）**

*通常処理として、nomal関数で "hello" を５回出力させる。*

*goroutine関数で "world" を５回出力させるプログラムを並列で動かす。*

```
func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	wg.Done()
}

func nomal(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go goroutine("world", &wg)
	nomal("hello")
	wg.Wait()
}
```

この時に、main関数内の.Wait()の置く位置を変更して処理結果の変化を確認する。

**pattern1：.Add()の前**
```
wg.Wait()
wg.Add(1)
go goroutine("world", &wg)
```
結果：ゴルーチンの処理を待たずにプログラムが終了する。

**pattern2：.Add()の直後**
```
wg.Add(1)
wg.Wait()   // fatal error: all goroutines are asleep - deadlock!
go goroutine("world", &wg) 
```
結果：ゴルーチンを走らせる前に.Wait()しているので、永遠に.Done()を待ち続けてエラーが発生する。

**pattern3：go で並列処理を走らせた直後**
```
wg.Add(1)
go goroutine("world", &wg)
wg.Wait()
nomal("hello")
```
結果：goroutineの処理を終えてから、nomalの処理が実行される。（＝.Wait()以降の処理は.Done()が呼ばれるまで処理されない）

**pattern4：main関数の一番最後**
```
wg.Add(1)
go goroutine("world", &wg)
nomal("hello")
wg.Wait()
```
結果：nomalの処理結果が出てから、goroutineの処理結果が出る。


##### 余談

.Done()は必ず並列処理の最後に実行するので、deferで宣言すると最初に忘れず書いておくことが出来る。
```
func goroutine(wg *sync.WaitGroup) {
  defer wg.Done()   // 最後に実行
  ...処理
}
```
