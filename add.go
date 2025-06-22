package add

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Add[T Number](first, second T) T {
	return first + second
}
