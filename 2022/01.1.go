package main

import "fmt"

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

	in, _ := io.ReadAll(os.Stdin)
	x, _ := parser.ParseBytes("", in)
	return x
}

func main() {
	in := parse(os.Stdin)
  max := 0
  for _, elf := range in.E {
    total := 0
    for _, item := range elf.I {
      total += item.Calories
    }
    if total > max {
      max = total
    }
  }
  fmt.Printf("%d\n", max)
}
