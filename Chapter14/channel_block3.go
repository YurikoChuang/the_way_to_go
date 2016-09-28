package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 50) //加了缓冲之后不会有 received   因为main线程已经停止(缓冲信道不阻塞)

	go func() {
		time.Sleep(15 * time.Second)
		x := <-c
		fmt.Println("Received ", x)
	}()

	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("send", 10)
}
