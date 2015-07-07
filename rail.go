package main

import (
	"flag"
	"fmt"

	"github.com/sopwithcamel/golang-practice/railfence"
)

func main() {
	offsetPtr := flag.Int("offset", 3, "Offset")
	inputPtr := flag.String("input", "", "String to encipher")
	flag.Parse()

	fmt.Println(railfence.Encode(*offsetPtr, *inputPtr))
}
