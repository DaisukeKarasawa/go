## okイディオム

Goでは、エラー処理として ok を使用する。

ok は特定の文をテストするために使用され、意図していないことを実行するプログラムを最小に抑えることができる。

**基本的な使われ方**

1：関数の戻り値のテスト

2：キーと値がマップに存在するかどうかのテスト

3：インターフェース変数が特定の型かどうかのテスト（型アサーション）

4：チャネルが閉じているかのテスト

### 関数の戻り値のエラー

ok は基本複数の戻り値を返す関数で使用され、その戻り値の値の１つにエラーが含まれる。

Goには try catch のようなパターンが無いので、戻り値のエラーが通常の変数と同じように扱われる。

```
package main

import (
	"errors"
	"fmt"
)

// 第一引数の配列内に第二引数の要素が含まれているかどうかを確認する
func checkElement(slice []int, val int) (bool, error) {
	hashmap := make(map[int]int)

	for _, item := range slice {
		hashmap[item] = item
	}

	// 含まれていなければ、エラーメッセージを一緒に返す
	if _, ok := hashmap[val]; !ok {
		return false, errors.New("That element is not included.")
	}
	return true, nil
}

func main() {
	if ok, err := checkElement([]int{1, 2, 3, 4}, 4); !ok {
		fmt.Println(err)
	} else {
		fmt.Println("No error")			// No error
	}

	if ok, err := checkElement([]int{1, 2, 3, 4}, 5); !ok {
		fmt.Println(err)
	} else {
		fmt.Println("No error")			// That element is not included.
	}
}
```

### キーと値がマップに存在するか

マップ内にキー、もしくは値が存在するかどうかを確認するために使用可能。

キーが存在する場合は true、それ以外の場合は false を返す。

```
package main

import "fmt"

// 第一引数の配列内に第二引数の配列の要素が含まれているかどうかを確認する
func checkElement(slice []int, val []int) bool {
	hashmap := make(map[int]int)

	for _, item := range slice {
		hashmap[item] = item
	}

	for _, v := range val {
		// val のそれぞれの要素 v が hashmap(slice) に存在するか
		if _, ok := hashmap[v]; !ok {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 3, 5}
	arr3 := []int{2, 4, 6}
	fmt.Println(checkElement(arr1, arr2))		// true
	fmt.Println(checkElement(arr1, arr3))		// false
}
```

### インターフェース変数が特定の型か（型アサーション）

**型アサーション**

動的に変数の型をチェックする機能。（≒ 隠れていた型を明確にする）

全ての型に対応できるinterface{}型を使用した変数に対する、実体の型が何かを動的に確認する。
```
package main

import "fmt"

type Snack struct {
	price interface{}
}

func main() {
	var p interface{}
	p = 120

	// 型が int かどうかを確認する [変数名.(任意の型)]
	switch value := p.(type) {
	case int:
		snack := Snack{value}
		fmt.Println(snack.price)		// 120
	default:
		fmt.Println("Not a number")
	}
}
```
そして、型アサーションに失敗するとエラーが発生するので、２つ目の変数に代入させることで

エラーの発生によるプログラムの停止を防いでうえで、失敗を確認できる。

※１つ目の変数にはその型の初期値が入るので、必要がなければ _ で省略できる。
```
package main

import (
	"fmt"
)

type Snack struct {
	price interface{}
}

func main() {
	var p interface{}
	p = 120

	// 上のコードを ok を使用して書き換えた例
	// 型が int でない場合、特定の処理を行わせる
	if _, ok := p.(int); !ok {
		fmt.Println("Not a number")
	}
	snack := Snack{p}
	fmt.Println(snack.price)	// 120
}
```

### チャネルが閉じているか

[チャネル](https://github.com/DaisukeKarasawa/go/tree/master/goroutine_prg/range_close)を使用してゴルーチンを行った時に、動作の完了の通知を受け取ることができる。

```
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	numbers := []int{1, 2, 3, 4, 5}

  // それぞれの要素を倍にしてチャネルに送信する
	go func() {
		for _, v := range numbers {
			ch <- v * 2
		}
		close(ch)
	}()

  // チャネルの送信終了を受け取るまで、受信した値を出力する
	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				fmt.Println("Channel closed")
				return
			}
		}
	}()

	time.Sleep(1 * time.Second)
}

// 2
// 4
// 6
// 8
// 10
// Channel closed
```