package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	out <- 2
	go f1(out)          //main也是一个协程
	//在其他协程运行时让 main 程序无限阻塞的通常做法是在 main 函数的最后放置一个{}。

}