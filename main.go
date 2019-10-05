package main

import (
	"fmt"
	"os"
)

// BUFSIZE | 読み込みバッファー
const BUFSIZE = 1024

func main() {
	file, err := os.Open(`./gopher.txt`)
	if err != nil {
		fmt.Errorf("file open error")
	}

	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)

		if n==0 {
			break
		}

		if err != nil {
			fmt.Errorf("file read error")
		}

		fmt.Println(string(buf[:n]))
	}

	fmt.Println("Hello, Gopher")
}
