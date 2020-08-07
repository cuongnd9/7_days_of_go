package main

import "fmt"

func main() {
	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}
	// while
	sum := 0
	for sum < 777 {
		sum += 7
	}
	fmt.Println(sum)
	// forever
	for {
		fmt.Println("go")
	}
}
