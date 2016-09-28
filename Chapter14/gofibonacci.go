package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	term := 25      //要计算多少阶斐波那契?
	i := 0
	c := make(chan int)
	start := time.Now()

	go fibnterms(term, c)
	for {
		if result, ok := <-c ; ok{
			fmt.Printf("fibonacci(%d) is: %d\n", i, result)
			i++
		}else{
			end := time.Now()
			delta := end.Sub(start)     //求出end 和 start的时间差
			fmt.Printf("longCalculation took this amount of time: %s\n", delta)
			os.Exit(0)
		}

	}
}


func fibnterms(term int , c chan int){
	for i := 0; i <= term; i++ {
		c <- fibonacci(i)
	}
	close(c)
}


func fibonacci(n int) (res int)  {         //计算斐波那契数列
	if n <= 1 {
		res = 1
	}else{
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}