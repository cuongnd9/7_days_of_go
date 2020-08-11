package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/xerrors"
)

func foo() error {
	return xerrors.Errorf("foo failed, %w", sql.ErrTxDone)
	// %w: the returned error will still implement an Unwrap method returning the operand, but the error's Format method will not return the wrapped error.
}

func main() {
	if err := foo(); xerrors.Is(err, sql.ErrTxDone) {
		fmt.Printf("🐛 got err, %+v\n", err)
	}
}
