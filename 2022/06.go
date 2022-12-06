package main

import "fmt"
import "aoc/util"

import "os"
import "io"

func find(l int, in []byte) int {
	for i := l; i < len(in); i++ {
		if len(util.ToSet(in[i-l:i])) == l {
      return i
		}
	}
  panic("")
}

func main() {
	in := util.OrDie(io.ReadAll(os.Stdin))
  fmt.Printf("%d\n", find(4, in))
  fmt.Printf("%d\n", find(14, in))
}
