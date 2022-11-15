package main

import "fmt"

func main() {
	var i interface{} = 4
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("this will pring too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("i is another type")
	}
}
