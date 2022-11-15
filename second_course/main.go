package main

import (
	"fmt"
)

func main() {
	s := [3]int{1, 2, 4}
	for k, v := range s{
		fmt.Println(k, v)
	}
}
