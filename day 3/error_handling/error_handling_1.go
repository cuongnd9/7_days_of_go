package main

import (
	"database/sql"
	"fmt"
)

func foo() error {
	return sql.ErrConnDone
}

func bar() error {
	return sql.ErrNoRows
}

func main() {
	if err := foo(); err != nil {
		fmt.Printf("🐛 got err, %v\n", err)
	}
	if err := bar(); err == sql.ErrNoRows {
		fmt.Printf("🐞 data not found, %v\n", err)
	}
}
