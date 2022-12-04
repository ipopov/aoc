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

func (x Range) Contains(point int) bool {
	return point >= x.Start && point <= x.End
}

func (x Range) Covers(y Range) bool {
	return x.Contains(y.Start) && x.Contains(y.End)
}

func Overlap(x, y Range) bool {
	return x.Contains(y.Start) || y.Contains(x.Start)
}

func main() {
	in := parse(os.Stdin)
	part1 := util.CountIf(func(p *RangePair) bool {
    return p.A.Covers(*p.B) || p.B.Covers(*p.A)
	}, in.R)
	part2 := util.CountIf(func(p *RangePair) bool {
    return Overlap(*p.A, *p.B)
	}, in.R)
	fmt.Printf("%d\n", part1)
	fmt.Printf("%d\n", part2)
}
