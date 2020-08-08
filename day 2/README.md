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

### WaitGroup

## documents

- [Go cheatsheet](https://devhints.io/go)
- [Anatomy of Channels in Go - Concurrency in Go](https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb)
