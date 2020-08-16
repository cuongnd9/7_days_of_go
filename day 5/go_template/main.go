package main

import (
	"fmt"

	"github.com/103cuong/uid"
)

func main() {
	length := 24
	fmt.Println("💅" + uid.Uid(length))
}
