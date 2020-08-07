package main

import "fmt"

func main() {
	var (
		name string = "Cuong Duy Nguyen"
		age  int    = 21
	)
	fmt.Printf("type: %T, value: %v\n", name, name) // type: string, value: Cuong Duy Nguyen
	fmt.Printf("type: %T, value: %v\n", age, age)   // type: int, value: 21
}
