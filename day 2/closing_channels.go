package main

func main()  {
  ch := make(chan string)

  go greet(ch)

  for {
   v, ok := <- ch
   if ok == true {
     println(v, ok)
   } else {
     println("loop broke", v, ok)
     break
   }
  }
  //
  //for v := range ch {
  //  println(v)
  //}
}

func greet(ch chan string)  {
  for i := 0; i < 7; i++ {
    ch <- "Hello 🙆‍♀️"
  }
  close(ch)
}
