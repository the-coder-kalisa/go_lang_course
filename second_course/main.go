package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	panic("something bad happend")
	fmt.Println("end")
}
