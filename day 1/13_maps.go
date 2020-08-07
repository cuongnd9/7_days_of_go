package main

import "fmt"

type Point struct {
  X, Y int
}

func main() {
  // 1. Basic
  var m map[string]Point = make(map[string]Point)
  fmt.Printf("%v \n", m) // map[]
  m["L"] = Point{6, 9}
  m["T"] = Point{0, 1}
  fmt.Printf("%v %v \n", m, m["T"]) // map[L:{6 9} T:{0 1}] {0 1}

  // 2. Map literals
  t := map[string]Point{
    "foo": Point{3, 1},
    "bar": Point {5, 6},
    "taz": {10, 20}, // shorthand Struct
  }
  fmt.Printf("%v \n", t)// map[bar:{5 6} foo:{3 1}]

  // 3. Mutating maps
  list := map[int]int {
    1: 10,
    2: 100,
    3: 1000,
  }
  fmt.Printf("%v\n", list) // map[1:10 2:100 3:1000]
  // insert an element
  list[4] = 10000
  fmt.Printf("%v\n", list) // map[1:10 2:100 3:1000 4:10000]
  // delete an element
  delete(list, 4)
  fmt.Printf("%v\n", list) // map[1:10 2:100 3:1000]
  // test that a key is present
  elem, ok := list[4]
  fmt.Println(elem, ok) // 0 false
}