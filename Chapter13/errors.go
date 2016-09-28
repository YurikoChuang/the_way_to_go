package main

import (
	"fmt"
	"errors"
	"os"
)

var errNotFound error = errors.New("not found error")
func main() {
	fmt.Printf("error: %v", errNotFound)
}
