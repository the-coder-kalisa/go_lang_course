package main

import "fmt"

func main() {
	statePulations := make(map[string]int)
	statePulations = map[string]int{
		"Califon": 10,
		"Texas": 100,
	}
	statePulations["Georgia"] = 1000
	delete(statePulations, "Texas")
	if pop, ok := statePulations["Georgia"]; ok {
		fmt.Println(pop)
	} 
}
