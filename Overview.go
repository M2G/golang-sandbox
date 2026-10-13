package main

import (
	"fmt"
	"uuid"
)

// GO 1.27

type Track struct {
	Title  string `json:"title"`
	Second int    `json:"second"`
}

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

	// uuid pkg
	fmt.Println(uuid.New())
	fmt.Println(uuid.NewV4())
	fmt.Println(uuid.NewV7())

	sessionId := uuid.MustParse("8be4df61-93ca-11d2-aa0d-00e098032b8c")

	fmt.Println("parsed:", sessionId)
	fmt.Println("nil:", uuid.Nil())
	fmt.Println("max:", uuid.Max())

	if sessionId == uuid.MustParse("8be4df61-93ca-11d2-aa0d-00e098032b8c") {
		fmt.Println("same session")
	}

	// JSON V2
	out, err := json.Marshal(Track{Title: "Hello World", Second: 123})
	fmt.Println(string(out), err)
}
