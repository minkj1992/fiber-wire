// https://go.dev/blog/defer-panic-and-recover
package main

import (
	"testing"
)

func TestPanicAndRecovery(t *testing.T) {
	PanicAndRecovery()

	// Calling Callee.
	// Printing in Callee 0
	// Printing in Callee 1
	// Printing in Callee 2
	// Printing in Callee 3
	// Panicking!
	// Defer in Callee 3
	// Defer in Callee 2
	// Defer in Callee 1
	// Defer in Callee 0
	// Recovered in Caller 4
	// Returned normally from Caller.
}
