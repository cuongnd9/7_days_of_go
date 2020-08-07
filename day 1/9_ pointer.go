package main

import "fmt"

func main() {
	i, j := 10, 50
	var a *int = &i
	fmt.Println(a, *a)
	*a = 17
	fmt.Println(a, *a, i)
	b := &j
	fmt.Println(b, *b)
	a = b
	fmt.Println(*a, *b)
	*b = 7
	fmt.Println(*a, *b)
}
