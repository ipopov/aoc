package util

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
