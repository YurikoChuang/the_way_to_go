package main

import (
	"time"
	"fmt"
)
func main() {
	ch:= make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1* time.Second)
}

func sendData(ch chan string){
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string){
	var input string
	for{
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
