package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("/code/gocode/src/the_way_to_go/Chapter12/input.dat")
	if inputError != nil {
		fmt.Println("An error occurred on opening the inputfile\n" + "Does the file exist?\n" + "Have you got access to it?\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		fmt.Printf("the input was: %s", inputString)
	}
}
