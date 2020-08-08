package main

import "fmt"

func main() {
  ch := make(chan string)

  fmt.Println("start...")

  go sayHi("Cuong", ch)
  go sayHi("Nhung", ch)
  fmt.Println("🌲 data: ", <-ch, <-ch)

  fmt.Println("before closing...")
  close(ch)
  fmt.Println("after closing...")

  go sayHi("Thu", ch)

  fmt.Println("🌳 data: ", <-ch)

  fmt.Println("end...")

  //start...
  //🌲 data:  👋 Hi Nhung 👋 Hi Cuong
  //before closing...
  //after closing...
  //🌳 data:
  //end...

}

func sayHi(name string, ch chan string) {
  ch <- "👋 Hi " + name
}
