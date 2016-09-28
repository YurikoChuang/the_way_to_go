package main

import (
	"fmt"
)

func main() {
	c := make (chan int)
	//consumer
	go func (){
		for {
			fmt.Print(<-c , " ")

		}
	}()

	// producer:
	for {
		select {    //如果有多个case可以执行  那么select 会随机select 一个case执行
		case c <- 0:
		case c <- 1:
		}


		close(c)
		break
	}



}