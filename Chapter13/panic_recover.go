package main

import (
	"fmt"
)

func badCall (){
	panic("bad end")
}

func test(){
	defer func(){
		if e := recover(); e!= nil {
			fmt.Printf("panicing %s \r\n", e)
		}
	}()

	badCall()
	fmt.Printf("after bad call \r\n")
}
func main() {
	fmt.Printf("calling test \r\n")
	test()
	fmt.Printf("test completed\r\n")
}
