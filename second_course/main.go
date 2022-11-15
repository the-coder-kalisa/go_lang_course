package main

import "fmt"

func main() {
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one")
	case 2, 4, 6:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}
