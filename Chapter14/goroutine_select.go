package main

import (
	"fmt"
	"os"
)

func tel_select(ch chan int, quit chan bool) {
	for i:=0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

func main() {
	var ok = true
	ch := make(chan int)
	quit := make(chan bool)

	go tel_select(ch, quit)
	for ok {
		select {
		case i:= <-ch:
			fmt.Printf("The counter is at %d\n", i)
		case <-quit:
			os.Exit(0)
		}
	}
}
