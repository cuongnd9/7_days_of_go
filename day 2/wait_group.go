package main

import "sync"

func main()  {
  var wg sync.WaitGroup
  println("start...")
  wg.Add(2)
  go greet("Nhung", &wg)
  go greet("Thu", &wg)
  wg.Wait()
  println("finish...")
}

func greet(name string, wg *sync.WaitGroup)  {
  println("Hello " + name)
  wg.Done()
}
