package main

import "fmt"
import "sort"
import "aoc/util"

import "os"
import "io"

import p "github.com/alecthomas/participle/v2"
import l "github.com/alecthomas/participle/v2/lexer"

type Item struct {
	Calories int `@Int Newline?`
}

type Elf struct {
	I []*Item `@@+ Newline?`
}

type Input struct {
	E []*Elf `@@+`
}

func parse(r io.Reader) *Input {
	lexer := l.MustSimple(
		[]l.SimpleRule{
			{"Int", `-?\d+`},
			{"Newline", `\n`},
		})

	parser := p.MustBuild[Input](p.Lexer(lexer))

	in := util.OrDie(io.ReadAll(os.Stdin))
	x := util.OrDie(parser.ParseBytes("", in))
	return x
}

func sum(i []int) int {
	var ret int
	for _, x := range i {
		ret += x
	}
	return ret
}

func main() {
	in := parse(os.Stdin)
	var totals sort.IntSlice
	for _, elf := range in.E {
		total := 0
		for _, item := range elf.I {
			total += item.Calories
		}
		totals = append(totals, total)
	}
	sort.Sort(sort.Reverse(totals))
	fmt.Printf("%d\n", sum(totals[0:1]))
	fmt.Printf("%d\n", sum(totals[0:3]))
}
