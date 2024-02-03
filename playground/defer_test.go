package main

import (
	"testing"
)

func TestDeferSimple(t *testing.T) {
	DeferEvaluate()
	DeferOrder()
}

func TestDeferWithInnerFunction(t *testing.T) {
	DeferWithFunction()
	DeferWithFunctionWithoudNamedReturnValue()
	DeferWithNestedFunction()
}
