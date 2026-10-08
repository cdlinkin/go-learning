package stage5generics

import (
	"cmp"
)

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Number interface {
	~int | ~float64
}

func Sum[T Number](nums []T) T {
	var sum T
	for _, v := range nums {
		sum += v
	}
	return sum
}

func Contains[T comparable](s []T, target T) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func Filter[T comparable](s []T, pred func(T) bool) []T {
	var newSlice []T
	for _, v := range s {
		item := pred(v)
		if item {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}

type Set[T comparable] struct {
	num map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		num: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(v T) {
	s.num[v] = struct{}{}
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.num[v]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.num)
}
