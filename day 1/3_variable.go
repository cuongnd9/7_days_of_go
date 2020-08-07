package main

import "fmt"

// global variables
var canFly bool
var age int = 21

func main() {
	// explict
	var name string = "Cuong Tran"
	// inferred
	var job, address = "Software Engineer", "Ho Chi Minh city"
	// shorthand
	gender := "male"
	// constant
	const homeTown = "Tam Ky city, Quang Nam province"
	fmt.Println(canFly, name, age, job, address, gender, homeTown)
}
