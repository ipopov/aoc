package main

import "fmt"
import "aoc/util"
import "golang.org/x/exp/slices"

import "os"
import "bytes"
import "io"

func toSet(x []byte) []byte {
	slices.Sort(x)
	return slices.Compact(x)
}

func common(x []byte) byte {
	l := len(x)
	util.Check(l%2 == 0)
	mid := l / 2
	i := util.SetIntersect(toSet(x[:mid]), toSet(x[mid:]))
	util.Check(len(i) == 1)
	return i[0]
}

func score(b byte) int {
	if b <= 'Z' {
		util.Check(b >= 'A')
		return 27 + int(b-'A')
	}
	util.Check(b >= 'a' && b <= 'z')
	return 1 + int(b-'a')
}

func part1(in io.Reader) {
	fmt.Printf("%d\n",
		util.Sum(
			util.Map(
				func(x []byte) int { return score(common(x)) },
				util.AsLines(in))))
}

func part2(in io.Reader) {
	l := util.AsLines(in)
	util.Check(len(l)%3 == 0)
	sum := 0
	for i := 0; i < len(l); i += 3 {
		i := util.SetIntersect(util.Map(toSet, l[i:i+3])...)
		util.Check(len(i) == 1)
		sum += score(i[0])
	}
	fmt.Printf("%d\n", sum)
}

func main() {
	in := util.OrDie(io.ReadAll(os.Stdin))
	part1(bytes.NewReader(in))
	part2(bytes.NewReader(in))
}
