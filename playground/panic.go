// https://go.dev/blog/defer-panic-and-recover
package main

import "fmt"

func PanicAndRecovery() {
	caller()
	fmt.Println("Returned normally from Caller.")
}

// 1. Recover는 고루틴의 패닉을 다시 제어하는 데 사용되는 내장 함수입니다.
// 2. Recover는 지연 함수(deferred functions) 내에서만 유용합니다.
// 3. 일반적인 실행 중에 recover를 호출하면 nil을 반환하고 다른 효과가 없습니다.
// 4. 현재 고루틴이 패닉 상태인 경우 recover를 호출하면 panic에 전달된 값을 캡처하고 정상적인 실행을 재개합니다.
// Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
func caller() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Caller", r)
		}
	}()
	fmt.Println("Calling Callee.")
	callee(0)
	fmt.Println("Returned normally from Callee.")
}

// 1. `panic`이 발생하면 현재 함수(F)의 실행이 즉시 중단됩니다.
// 2. 이후, 현재 함수(F) 내에서 연기된(deferred) 함수들이 정상적으로 실행됩니다.
// 3. 그리고나서 현재 함수(F)는 호출한 곳(Caller)로 돌아갑니다.
// 4. 호출한 곳에서는 마치 `panic` 함수를 직접 호출한 것처럼 동작하게 됩니다.
// 5. 이러한 과정이 반복되어 call stack을 따라 올라가면서 남은 call stack안에 존재하는 함수들이 종료됩니다.
// 6. 만약 위 과정을 처리하는 동안 고루틴 내에서 `panic`을 처리하는 로직이 없다면(recovery), 해당 고루틴은 더 이상 실행을 계속할 수 없게 되고, 이는 main으로 전파되어 프로그램을 종료시킵니다.
// Panic is a built-in function that stops the ordinary flow of control and begins panicking.
// When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.
// To the caller, F then behaves like a call to panic.
// The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes.
// Panics can be initiated by invoking panic directly.
// They can also be caused by runtime errors, such as out-of-bounds array accesses.
func callee(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in Callee", i)
	fmt.Println("Printing in Callee", i)
	callee(i + 1)
}

// Real world recover 처리
// https://go.dev/src/encoding/json/encode.go
// func (e *encodeState) marshal(v any, opts encOpts) (err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			if je, ok := r.(jsonError); ok {
// 				err = je.error
// 			} else {
// 				panic(r)
// 			}
// 		}
// 	}()
// 	e.reflectValue(reflect.ValueOf(v), opts)
// 	return nil
// }
