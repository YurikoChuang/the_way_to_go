package main

import (
	"fmt"
	"net"
)

//make a tcp server
func main() {
	fmt.Println("starting the server ...")

	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	fmt.Println("starting ok!")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}

		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}

		fmt.Printf("received data: %v", string(buf))
	}
}
