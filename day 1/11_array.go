package main

import "fmt"

func main() {
	var list [7]int
	list[0] = 7
	fmt.Println(list, list[1], len(list)) // [7 0 0 0 0 0 0] 0 7

	for i, v := range list {
		fmt.Println(i, v)
	}

	oddNumbers := [3]int{3, 5}
	oddNumbers[2] = 7
	fmt.Println(oddNumbers) // [3 5 7]

	var fullName [4]string = [4]string{"Cuong", "Duy", "Nguyen"}
	fmt.Println(fullName) // [Cuong Duy Nguyen ]
}
