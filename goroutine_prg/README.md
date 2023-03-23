## sync.WaitGroupとclose()の違い、それぞれの依存関係

### 「sync.WaitGroupとclose()の違い」

**[sync.WaitGroup](https://github.com/DaisukeKarasawa/go/tree/master/goroutine_prg/sync_wait)**

並列処理を知らせることで、その処理を終了するまで待機させることが出来る。

**[close()](https://github.com/DaisukeKarasawa/go/tree/master/goroutine_prg/range_close)**

チャネルを使用した範囲節のfor処理で、チャネルにこれ以上データが入らないことを知らせる。

### 「それぞれの「処理待機」に関する依存関係」

**sync.WaitGroup**

*「並列処理を待ってから次の処理を行わせるものなので、並列処理の終了を待っている.Wait()に依存する。」*

.Wait()で、.Add()した並列処理を終了まで待機させることができる。

次の処理に移行するには、.Done()で.Wait()に対して終了したことを知らせる必要がある。

**close()**

*「範囲節のforに対して、チャネルの送信の終了を知らせるものなので、チャネルを受信し続けるforに依存する。」*

何もしなければ、範囲節のforはチャネルに対して永遠に受信し続ける。

次の処理に移行するには、close()でforに対して送信の終了を知らせる必要がある。

#### 結論

main関数内にチャネルを使用した範囲節のforがあれば、必ずclose()が呼び出されるまでチャネルに対して受信し続けるので、

close()を呼び出す間はsync.WaitGroupを使用せずに並列処理を行うことが出来る。

##### sync.WaitGroupとclose()のコード*

[sync.WaitGroup](https://github.com/DaisukeKarasawa/go/blob/master/goroutine_prg/sync_wait/sync.go) / [close()](https://github.com/DaisukeKarasawa/go/blob/master/goroutine_prg/range_close/close.go)