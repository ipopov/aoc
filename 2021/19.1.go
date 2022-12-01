package main

import "fmt"

import "os"
import "io"

import p "github.com/alecthomas/participle/v2"
import l "github.com/alecthomas/participle/v2/lexer"

type Beacon struct {
	X int `@Int Comma`
	Y int `@Int Comma`
	Z int `@Int`
}

type Probe struct {
	ScannerId int       `HeaderStart @Int HeaderEnd`
	B         []*Beacon `@@+`
}

type Input struct {
	P []*Probe `@@+`
}

func parse(r io.Reader) *Input {
	lexer := l.MustSimple(
		[]l.SimpleRule{
			{"HeaderStart", `--- scanner`},
			{"HeaderEnd", `---`},
			{"Comma", `,`},
			{"Int", `-?\d+`},
			{"Whitespace", `[\n ]+`},
		})

	parser := p.MustBuild[Input](p.Lexer(lexer), p.Elide("Whitespace"))

	in, _ := io.ReadAll(os.Stdin)
	x, _ := parser.ParseBytes("", in)
	return x
}

type transformation struct {
  x, y, z int
}

func main() {
	in := parse(os.Stdin)
	fmt.Printf("%+v\n", in)
  //x, y := in.P[0], in.P[1]
  t := transformation { 1, 1, 1}
	fmt.Printf("%+v\n", t)
}
