package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	// basic struct
	p := Point{6, 9}
	fmt.Println(p, p.x, p.y)
	p.x = 1
	p.y = 7
	fmt.Println(p, p.x, p.y)

	// struct with pointer
	var a *Point = &p
	fmt.Println(a, *a)
	fmt.Println(a.x, (*a).x) // the same result

	// init struct
	var (
		v1 = Point{1, 2}  // has type Point
		v2 = Point{x: 1}  // Y:0 is implicit
		v3 = Point{}      // X:0 and Y:0
		k  = &Point{1, 2} // has type *Point
	)
	fmt.Println(v1, v2, v3, k)
}
