package main

import "fmt"

func main() {
	a := 1.05
	b := 7
	c := a + float64(b)
	d := int(a) + b
	fmt.Println(c, d)
}
