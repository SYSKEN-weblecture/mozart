package main

import (
	"fmt"
)

var hoge, fuga bool // 複数の同一型変数は←の感じで宣言できる。boolのゼロ値はfalse
// hogefuga := "you" //関数の外では`:=`を使えない

func main() {
	var piyo int       // intはゼロ値は0
	var hiyoko int = 3 // 初期値を与えることができる

	piyoko := "hello world"      // `:=`を使うと型を宣言しなくても変数宣言できる。この場合、piyokoの型はstring
	a, b, c := false, "f*ck", 42 // `:=`を使うと、複数の変数に別々の型を与えて宣言することもできる。

	fmt.Println(hoge, fuga, piyo, hiyoko, piyoko, a, b, c)
}
