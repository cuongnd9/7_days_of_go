package main

import (
  "fmt"
  "runtime"
  "sync"
  "time"
)

func g1()  {
  fmt.Println("g1.....")
}
func g2()  {
  fmt.Println("g2.....")
  wg.Done()
}

func g3()  {
  fmt.Println("g3.....")
  wg.Done()
}

func sum(c chan int, s []int) {
 sum := 0
 for _, v := range s {
   sum += v
 }
 c <- sum
}

var wg sync.WaitGroup
func main() {
  go g1()
  time.Sleep(time.Second)
  gc := runtime.NumGoroutine()
  fmt.Println(gc)

  // synchronized goroutines
  fmt.Println("begin")
  wg.Add(2)

  go g2()
  go g3()

  wg.Wait()
  fmt.Println("end")

  s := []int{1, 5, 2, 8}
  c := make(chan int, 5) // initialize number of channels.
  go sum(c, s[:len(s) / 2])
  go sum(c, s[len(s) / 2:])
  x, y := <- c, <- c
  fmt.Println(x, y, x + y) // 10 6 16
}
