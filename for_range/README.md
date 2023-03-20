### Goで範囲節によるfor

#### [Goで範囲節によるforを文字列に対して使った時、vに入るのはrune型](https://github.com/DaisukeKarasawa/go/blob/master/for_range/range.go)

文字列に対して範囲節のforを使用して、１文字ずつswitchやifで分岐処理を行った時

文字を比較するとエラーが発生した。
```
for _, v := range str {
  switch v {
  case "s":       // cannot convert "s" (untyped string constant) to type rune
    ...処理
  default:
    ...処理
  }

  if v == "s" {   // invalid operation: v == "s" (mismatched types rune and untyped string)
    ...処理
  }
}
```

**問題**

*cannot convert "文字" (untyped string constant) to type rune*

*invalid operation: v == "s" (mismatched types rune and untyped string)*

「文字列をrune型に変換できない」「型の不一致」と出ているので、string()でvをキャストする。
```
for _, v := range str {
  switch string(v) {
  case "s":
    ...処理   // 問題なく処理できる
  default:
    ...処理
  }

  if string(v) == "s" {
    ...処理   // 問題なく処理できる
  }
}
```
結果として、文字列に対して範囲節のforを使用した時にvに入るのは、文字ではなく*rune型*だということが分かった。

なので、string()でキャストをすることでstring型に変える必要があった。

##### rune型

Goにおける文字（Unicodeコードポイント）を表す型。

別名int32として定義されていて、rune型の値は32ビット符号付き整数と同じ。

（シングルクォートで囲って一文字のみ定義可能）
```
r1 := 'g'
r2 := 'go'          // more than one character in rune literal

fmt.Printf("%v", r) // 103
```