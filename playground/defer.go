// https://go.dev/blog/defer-panic-and-recover
package main

import "fmt"

func DeferEvaluate() {
	i := 0
	// A deferred function’s arguments are evaluated when the defer statement is evaluated.
	// 아래 statement가 evaulate될 때, i가 evaluate되고 이후에 defer이 실행되는 시점에 처리된다.
	// i = 0 (o)
	// i = 1 (x)
	defer fmt.Println(i) // 0
	i++
}

func DeferOrder() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i) // 3,2,1,0
	}
}

// Test whether expression in defer impacts outside variable: No
// Deferred functions may read and assign to the returning function’s named return values.
func DeferWithFunction() (i int) {
	defer func() { i++; fmt.Printf("In defer: %d\n", i) }() // 1
	fmt.Printf("In function: %d\n", i)                      // 0
	return i
}

func DeferWithFunctionWithoudNamedReturnValue() int {
	i := 0
	defer func() { i++; fmt.Printf("In defer: %d\n", i) }() // 1
	fmt.Printf("In function: %d\n", i)                      // 0
	return i
}

func DeferWithNestedFunction() (i, j int) {
	defer func() { i++; fmt.Printf("In defer i,j: %d, %d\n", i, j) }() //In defer i,j: 3, 1
	defer func() { i++; fmt.Printf("In defer i: %d\n", i) }()          // In defer i: 2
	defer func() { j++; fmt.Printf("In defer j: %d\n", i) }()          // In defer j: 1
	defer func() { i++; fmt.Printf("In defer i: %d\n", i) }()          // In defer i: 1
	return i, j
}
