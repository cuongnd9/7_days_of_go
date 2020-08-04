package main

import (
	"fmt"
	"runtime"
	"time"
)

func compare(x, y int) (z string) {
	if x > y {
		z = "x is greater than y"
	} else {
		z = "x is smaller than y"
	}
	return
}

func isEven(number int) bool {
	if num := number % 2; num == 0 {
		return true
	}
	return false
}

// switch case
func checkOS() string {
	switch os := runtime.GOOS; os {
	case "linux":
		return "🐧 Linux"
	case "darwin":
		return "💻 macOS"
	default:
		return os
	}
}

// switch case without condition
func showGreeter() string {
	h := time.Now().Hour()
	switch {
	case h < 12:
		return "good morning"
	case h < 17:
		return "good afternoon"
	default:
		return "good evening"
	}
}

func main() {
	t := compare(6, 7)
	fmt.Println(t)
	fmt.Println(isEven(7))
	fmt.Println(checkOS())
	fmt.Println(showGreeter())
}
