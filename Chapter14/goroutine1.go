package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("in main()")
	go longWait()
	go shortWait()
	fmt.Println("about to sleep in main()")
	time.Sleep(10 * 1e9)        //单位 纳秒ns
	fmt.Println("at the end of main()")

}

func longWait(){
	fmt.Println("beginning longWait()")
	time.Sleep(5 * time.Second);
	fmt.Println("end of longWait()")
}

func shortWait(){
	fmt.Println("beginning shortWait()")
	time.Sleep(2 * time.Second);
	fmt.Println("end of shortWait()")
}
