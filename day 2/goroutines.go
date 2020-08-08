package main

import "fmt"

func main() {
  ch := make(chan string)

  go sayHi("Cuong", ch)
  go sayHi("Nhung", ch)
  go sayHi("Thu", ch)

  fmt.Println(<-ch, <-ch, <-ch)
}

func sayHi(name string, ch chan string) {
  ch <- "👋 Hi " + name
}
