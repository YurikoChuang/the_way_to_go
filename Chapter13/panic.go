package main

import "fmt"

func main() {
	fmt.Println("starting the program")
	panic("a severe error occurred: stopping the program")
	fmt.Println("ending the program")
}
