package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	canFly   bool
}

func main() {
	t := reflect.TypeOf(Animal{})
	filed, _ := t.FieldByName("Name")
	fmt.Println(filed.Tag)
}
