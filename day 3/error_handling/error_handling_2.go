package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func baz() error {
	return errors.Wrap(sql.ErrNoRows, "baz failed") // message + stack trace
}

func foo() error  {
	return errors.WithMessage(sql.ErrTxDone, "foo failed") // only message
}


func bar() error  {
	return errors.WithStack(sql.ErrTxDone) // no additional contextual text
}

func main() {
	if err := baz(); errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("🐛, %+v\n", err)
	}
	if err := foo(); errors.Cause(err) == sql.ErrTxDone {
		fmt.Printf("🐞, %v\n", err)
	}
	if err := bar(); errors.Cause(err) == sql.ErrTxDone {
		fmt.Printf("🦟, %+v\n", err)
	}
	// %+v: show full error, include stacktrace
	// %v: exclude stacktrace
}
