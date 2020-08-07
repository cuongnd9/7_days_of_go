package main

import "fmt"

func add(x int) func(int) int {
  return func(y int) int {
    return x + y
  }
}

func subtract(x float64) func(float64) float64  {
  return func(y float64) float64 {
    return x - y;
  }
}

func main() {
  value := add(3)(4)
  fmt.Println(value) // 7
  fmt.Println(subtract(9.0)(8.12))
}