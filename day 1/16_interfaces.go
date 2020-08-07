package main

import "fmt"

type Animal struct {
  name, color string
}

func (animal Animal) Miaow() {
  fmt.Printf("meow moew %s", animal.name)
}

type Cat interface {
  Miaow()
}

func main() {
  var cat Cat = Animal{"Miki", "white"}
  cat.Miaow()
}
