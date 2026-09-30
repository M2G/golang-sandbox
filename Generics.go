package main

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
