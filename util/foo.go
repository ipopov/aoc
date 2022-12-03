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

func OrDie[T any](t T, err error) T {
  if err != nil {
    panic(err)
  }
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
  for s := bufio.NewScanner(r); s.Scan() ; {
    ret = append(ret, slices.Clone(s.Bytes()))
  }
  return ret
}
