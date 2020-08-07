package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func swap(a, b int) (int, int) {
	return b, a
}

// named return values
func split(number int) (x, y int) {
	x = number / 10
	y = number % 10
	return
}

func main() {
	fmt.Println(add(6, 9))        // 15
	fmt.Println(subtract(17, 10)) // 7
	a, b := swap(10, 17)
	fmt.Println(a, b)      // 17 10
	fmt.Println(split(19)) // 1 9
}
