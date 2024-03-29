package util

import "golang.org/x/exp/constraints"
import "golang.org/x/exp/slices"
import "io"
import "bufio"

func Check(b bool) {
	if !b {
		panic("assertion failure")
	}
}

func OkOrDie(err error) {
	Check(err == nil)
}

func OrDie[T any](t T, err error) T {
	OkOrDie(err)
	return t
}

func Sum[T constraints.Integer](i []T) T {
	var ret T
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

func setIntersect[E constraints.Ordered](s1, s2 []E) []E {
	var ret []E
	var i, j int
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			i++
			continue
		}
		if s1[i] == s2[j] {
			ret = append(ret, s1[i])
			i++
			j++
			continue
		}
		j++
	}
	return ret
}

func SetIntersect[E constraints.Ordered](xs ...[]E) []E {
	switch len(xs) {
	case 0:
		return []E{}
	case 1:
		return xs[0]
	}
	return setIntersect(xs[0], SetIntersect(xs[1:]...))
}

func AsLines(r io.Reader) [][]byte {
	var ret [][]byte
	for s := bufio.NewScanner(r); s.Scan(); {
		ret = append(ret, slices.Clone(s.Bytes()))
	}
	return ret
}

func Map[T, U any](f func(T) U, x []T) []U {
	var ret []U
	for _, i := range x {
		ret = append(ret, f(i))
	}
	return ret
}

func CountIf[T any](predicate func(T) bool, x []T) int {
	return Sum(Map(func(x T) int {
    if predicate(x) {
      return 1
    }
    return 0
	}, x))
}

type SetElement interface {
	comparable
	constraints.Ordered
}

func ToSet[T SetElement](x []T) []T {
	x = slices.Clone(x)
	slices.Sort(x)
	return slices.Compact(x)
}

