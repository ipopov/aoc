package util

import "golang.org/x/exp/constraints"

func OrDie[T any](t T, err error) T {
  if err != nil {
    panic(err)
  }
  return t
}

func Sum(i []int) int {
	var ret int
	for _, x := range i {
		ret += x
	}
	return ret
}

func Max[T constraints.Ordered](x, y T) T {
	if x >= y {
		return x
	} else {
		return y
	}
}

func Abs[T constraints.Signed](x T) T {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
