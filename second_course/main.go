package main

import "fmt"

func main() {
	var r = complex(3, 2)
	fmt.Printf("%v %T\n", real(r), r)
	fmt.Printf("%v %T", imag(r), r)
}
