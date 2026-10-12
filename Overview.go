package main

import "fmt"

// GO 1.27
// generic method
type Result[T any] struct {
	val T
	err error
}

func (r Result[T]) Then[R any](f func(T) R) Result[R] {
	if r.err != nil {
		return Result[R]{err: r.err}
	}
	return Result[R]{val: f(r.val)}
}

func main() {
	port := Result[int]{val: 8080}
	addr := port.Then(func(p int) string {
		return fmt.Sprintf("localhost:%d", p)
	})

	length := addr.Then(func(s string) int { return len(s) })
	fmt.Println(addr.val, length.val)
}
