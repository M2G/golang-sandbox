package main

import "fmt"

type Pair struct {
	A, B int
}

func (p Pair) Use[R any](f func(int, int) R) R {
	return f(p.A, p.B)
}

func (p Pair) UseAll[R any, F ~func(int, int) R](ops []F) []R {
	out := make([]R, len(ops))
	for i, op := range ops {
		out[i] = p.Use(op)
	}
	return out
}

func add[T ~int](a, b T) T {
	return a + b
}

func multiply[T ~int](a, b T) T {
	return a * b
}

func join[T ~string](a T, b T) string {
	return fmt.Sprintf("%d+%d", a, b)
}

func main() {
	p := Pair{1, 2}

	fmt.Println(p.Use(add))
	fmt.Printf(p.Use(join))

	type BinOp func(int, int) int

	ops := []BinOp{add, multiply}

	fn := BinOp(add)
	fmt.Println(p.UseAll(ops))
	fmt.Println(fn(2, 3))

	ch := make(chan BinOp, 1)
	ch <- multiply
	fmt.Println(p.Use(<-ch))

	type Holder struct{ Fn BinOp }
	h := Holder{Fn: fn}
	fmt.Println(p.Use(h.Fn))

	named := map[string]BinOp{"add": add, "multiply": multiply}
	fmt.Println(named["add"])

}
