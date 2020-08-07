package main

import "fmt"

func main() {
	// 1. basic slice
	a := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	a = a[:7]
	fmt.Println(a) // [1 2 3 5 6 7 8]

	// 2. slices are like references to arrays
	pets := []string{"cat", "dog", "bird", "fish"}
	fmt.Println(pets) // [cat dog bird fish]
	list := pets[0:2]
	fmt.Println(list) // [cat dog]
	list[1] = "pig"
	fmt.Println(list) // [cat pig]
	fmt.Println(pets) // [cat pig bird fish]

	// 3. slice length & capacity
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// len=6 cap=6 [2 3 5 7 11 13]

	// 4. slice the slice to give it zero length.
	s = s[:1]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// len=1 cap=6 [2]

	// Extend its length.
	s = s[:2]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// len=2 cap=6 [2 3]

	// Drop its first two values.
	s = s[2:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// len=0 cap=4 []

	s = s[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// len=4 cap=4 [5 7 11 13]

	// 5. nil slices
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var t []int
	fmt.Println(t, len(t), cap(t)) // [] 0 0
	if t == nil {
		fmt.Println("nil!") // nil!
	}

	// 6. Creating a slice with make
	x := make([]int, 10)
	fmt.Printf("%d %d %v \n", len(x), cap(x), x)
	// 10 10 [0 0 0 0 0 0 0 0 0 0]

	y := make([]int, 0, 5)
	fmt.Printf("%d %d %v \n", len(y), cap(y), y)
	// 0 5 []

	z := y[:5]
	fmt.Printf("%d %d %v \n", len(z), cap(z), z)
	// 5 5 [0 0 0 0 0]

	// 7. Appending to a slice
	// cap will auto scale up depend on the added elements
	m := []int{}
	fmt.Printf("len=%d cap=%d %v\n", len(m), cap(m), m)
	// len=0 cap=0 []

	m = append(m, 6, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(m), cap(m), m)
	// len=2 cap=2 [6 5]

	m = append(m, 1, 2, 5, 6, 8, 9, 3)
	fmt.Printf("len=%d cap=%d %v\n", len(m), cap(m), m)
	// len=9 cap=10 [6 5 1 2 5 6 8 9 3]

	// 8. range
	// contain i(index) & v(clone value)
	var numbers []int = []int{1, 3, 7, 8}
	for i, v := range numbers {
		fmt.Println(i, v)
	}
	// 0 1
	// 1 3
	// 2 7
	// 3 8

	// skip the index or value by assigning to _
}
