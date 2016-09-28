package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	who := "Alice "
	if len (os.Args) >1{
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("good morning, ", who)
}
