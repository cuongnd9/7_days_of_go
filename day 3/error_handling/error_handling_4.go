package main

import (
	"database/sql"
	"fmt"
	"errors"
)

func foo() error {
	return fmt.Errorf("foo failed: %w", sql.ErrNoRows)
}

// does not support call stack output (stacktrace)
func main() {
	if err := foo(); errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("🐛 got err, %+v\n", err)
	}
}
