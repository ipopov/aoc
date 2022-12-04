package main

import "fmt"
import "aoc/util"

import "os"
import "io"

import p "github.com/alecthomas/participle/v2"
import l "github.com/alecthomas/participle/v2/lexer"

type Range struct {
	Start int `@Number Dash`
	End   int `@Number`
}

type RangePair struct {
	A *Range `@@ Comma`
	B *Range `@@ Newline`
}

type Input struct {
	R []*RangePair `@@+`
}

func parse(r io.Reader) *Input {
	lexer := l.MustSimple(
		[]l.SimpleRule{
			{"Number", `[0-9]+`},
			{"Dash", `-`},
			{"Newline", `\n`},
			{"Comma", `,`},
		})

	parser := p.MustBuild[Input](p.Lexer(lexer))

	in := util.OrDie(io.ReadAll(os.Stdin))
	x := util.OrDie(parser.ParseBytes("", in))
	return x
}

func (x Range) Covers(y Range) bool {
	return y.Start >= x.Start && y.End <= x.End
}

func main() {
	in := parse(os.Stdin)
	sum := util.Sum(util.Map(func(p *RangePair) int {
		if p.A.Covers(*p.B) || p.B.Covers(*p.A) {
			return 1
		}
		return 0
	}, in.R))
	fmt.Printf("%d\n", sum)
}
