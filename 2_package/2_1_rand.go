package main

import (
	"fmt"
	"math/rand"
)

// ↑と↓は同義だが、↑の書き方が推奨される。`$ go fmt`とかやると↑になる。
// import "fmt"
// import "math/rand"

func main() {
	fmt.Println("random number: ", rand.Intn(100))
	// `fmt.Println()`や`rand.Intn()`のように、外部に公開されたmethodは頭文字を大文字にする。||大文字になっている。
}
