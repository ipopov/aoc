package main

import "fmt"
import "aoc/util"
import "golang.org/x/exp/slices"

import "os"
import "bufio"

func common(x []byte) byte {
  l := len(x)
  util.Check(l % 2 == 0)
  mid := l/2
  a, b := x[:mid], x[mid:]
  slices.Sort(a)
  a = slices.Compact(a)
  slices.Sort(b)
  b = slices.Compact(b)

  i := util.SetIntersect(a, b)
  util.Check(len(i) == 1)
  return i[0]
}

func score(b byte) int {
  if b <= 'Z' {
    util.Check(b >= 'A')
    return 27 + int(b - 'A')
  }
  util.Check(b >= 'a' && b <= 'z')
  return 1 + int(b - 'a')
}

func main() {
  s := bufio.NewScanner(os.Stdin)
  sum := 0
  for s.Scan() {
    sum += score(common(s.Bytes()))
  }
  fmt.Printf("%d\n", sum)
}
