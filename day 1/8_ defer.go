package main

import "fmt"

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func main() {
	fmt.Print("counting ")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Print("done ")
	// counting done 9 8 7 6 5 4 3 2 1 0

	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)
	// value of a before deferred function call 10
	// value of a in deferred function 5
}
