package main

import "fmt"

func main() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal 10")
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal 20")
	default:
		fmt.Println("greater than 2-")
	}
}	
