package main

import (
	"fmt"
	"github.com/103cuong/go_grpc/pkg/cmd/server"
	"os"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
