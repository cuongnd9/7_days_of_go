package main

import (
  "fmt"
  "time"
)

type User struct {
  name string
  dob int
}

func (u User) age() int {
  return time.Now().Year() - u.dob
}

// Using pointer to modify struct's methods
func (u *User) changeName(name string) {
  u.name = name
}

func main() {
  user := User{"Cuong Duy Nguyen", 1998}
  fmt.Println(user.age()) // 22
  fmt.Println(user.name) // Cuong Duy Nguyen * not modify
  user.changeName("Nguyễn Duy Cương")
  fmt.Println(user.name) // Nguyễn Duy Cương * modify
}
