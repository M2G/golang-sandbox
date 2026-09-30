package main

type Pair struct {
	A, B int
}

func (p Pair) Use[R any](f func(int, int) R) R {
	return f(p.A, p.B)
}
