package main

import (
	"fmt"
	"iter"
	"slices"
)

type Iterator[V any] struct {
	seq iter.Seq[V]
}

func From[V any](src []V) *Iterator[V] {
	return &Iterator[V]{seq: slices.Values(src)}
}

func (it *Iterator[V]) Filter(pred func(V) bool) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for v := range it.seq {
				if !pred(v) {
					continue
				}
				if !yield(v) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Skip(n int) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			count := 0
			for v := range it.seq {
				if count < n {
					count++
					continue
				}
				if !yield(v) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Take(n int) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			if n <= 0 {
				return
			}
			count := 0
			for v := range it.seq {
				if !yield(v) {
					return
				}
				count++
				if count >= n {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Reverse() *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			items := slices.Collect(it.seq)
			for i := len(items) - 1; i >= 0; i-- {
				if !yield(items[i]) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) Map(f func(V) V) *Iterator[V] {
	return &Iterator[V]{
		seq: func(yield func(V) bool) {
			for v := range it.seq {
				if !yield(f(v)) {
					return
				}
			}
		},
	}
}

func (it *Iterator[V]) ForEach(f func(V)) {
	for v := range it.seq {
		f(v)
	}
}

func (it *Iterator[V]) Collect() []V {
	return slices.Collect(it.seq)
}

func MapTo[V, U any](it *Iterator[V], f func(V) U) *Iterator[U] {
	return &Iterator[U]{
		seq: func(yield func(U) bool) {
			for v := range it.seq {
				if !yield(f(v)) {
					return
				}
			}
		},
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := From(nums).
		Filter(func(v int) bool { return v%2 == 0 }).
		Map(func(v int) int { return v * 10 }).
		Take(2).
		Collect()
	fmt.Println(result) // [20 40]

	reversed := From(nums).Reverse().Collect()
	fmt.Println(reversed) // [10 9 8 ... 1]

	From(nums).
		Filter(func(v int) bool { return v > 3 }).
		Map(func(v int) int { return v * 10 }).
		Skip(2).
		Take(2).
		ForEach(func(v int) { fmt.Printf("%d ", v) }) // 60 70
	fmt.Println()

	strs := MapTo(From(nums), func(x int) string {
		return fmt.Sprintf("%d", x)
	}).Collect()
	fmt.Println(strs) // [1 2 3 ... 10]
}
