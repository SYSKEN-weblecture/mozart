package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int { // 引数の型は後ろに書く。 返り値の型を明示する。
	return x + y
}

func sub(x, y int) int { // 引数の型はこのように書ける
	// 大小比較による実装
	// if x < y {
	// 	return y - x
	// }
	// return x - y

	// 絶対値取得による実装
	return int(math.Abs(float64(x - y)))
}

func swap(a, b string) (string, string) { // 複数のresultを返すことがでできるので、入れ替えも楽勝
	return b, a
}

func addsub(x, y int) (ad, su int) { // 予め返す変数を指定することができる。読みにくくなるので短い関数でのみ使う
	ad = add(x, y)
	su = sub(x, y)

	return
}

func main() {
	fmt.Println("add:", add(8, 6))
	fmt.Println("sub:", sub(6, 8))

	fmt.Println(addsub(6, 8))

	a, b := "hello", "world"
	fmt.Println(a + " " + b)
	fmt.Println(swap(a, b))
}
