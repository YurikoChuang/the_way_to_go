package main

import "os"

func cat(f *os.File){
	const NBUF = 512
	var buf [NBUF]byte

	for {
		switch nr, err := f.Read(buf[:]); true{
		case nr < 0 :

		}
	}
}


func main() {


}
