# day 2

## 🍼 Concurrency

### Goroutines

```go
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
```

### Buffered channels

> Buffered channels limit the amount of messages it can keep.

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
ch <- 3
// fatal error:
// all goroutines are asleep - deadlock!
```

### Closing channels

A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```go
v, ok := <-ch
```

`ok` is `false` if there are no more values to receive and the channel is closed.

```go
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
```

### WaitGroup

> A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. The goroutine calls wg.Done() when it finishes.

## documents

- [Go cheatsheet](https://devhints.io/go)
- [Understanding Golang and Goroutines](https://medium.com/technofunnel/understanding-golang-and-goroutines-72ac3c9a014d)
- [Anatomy of Channels in Go - Concurrency in Go](https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb)
- [Using Command Line Flags in Go](https://flaviocopes.com/go-command-line-flags/)
- [How To Use the Flag Package in Go](https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go)
- [Switch to Go Modules from Go Dep](https://levelup.gitconnected.com/switch-to-go-modules-from-go-dep-fcdd4aa41bd5)
- [Using Go Modules](https://blog.golang.org/using-go-modules)