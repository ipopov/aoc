package main

import "fmt"
import "errors"
import "aoc/util"

import "os"
import "io"

import p "github.com/alecthomas/participle/v2"
import l "github.com/alecthomas/participle/v2/lexer"

type Play int
const (
  Rock Play = iota
  Paper
  Scissors
)

func (x Play) Beats(y Play) bool {
  ret := (x - y + 3) % 3 == 1
  return ret
}

func (p *Play) Capture(s []string) error {
  switch s[0] {
  case "A", "X":
    *p = Rock
    return nil
  case "B", "Y":
    *p = Paper
    return nil
  case "C", "Z":
    *p = Scissors
    return nil
  }
  return errors.New("blah")
}

type Round struct {
  OpponentPrediction *Play  `@PlayA Space`
  Recommendation *Play  `@PlayB Newline`
}

type Match struct {
	R []*Round `@@+`
}

func parse(r io.Reader) *Match {
	lexer := l.MustSimple(
		[]l.SimpleRule{
			{"PlayA", `[ABC]`},
			{"PlayB", `[XYZ]`},
			{"Space", ` `},
			{"Newline", `\n`},
		})

	parser := p.MustBuild[Match](p.Lexer(lexer))

	in := util.OrDie(io.ReadAll(os.Stdin))
	x := util.OrDie(parser.ParseBytes("", in))
	return x
}

func BaseScore(p Play) int {
  switch p {
  case Rock:
  return 1
  case Paper:
  return 2
  case Scissors:
  return 3
  }
  panic("")
}

func RoundScore(a, b Play) int {
  if a == b {
  return 3
  }
  if a.Beats(b) {
  return 6
  }
  return 0
}

func main() {
	in := parse(os.Stdin)
  sum := 0
  for _, r := range in.R {
    sum += BaseScore(*r.Recommendation) + RoundScore(*r.Recommendation, *r.OpponentPrediction)
  }
  fmt.Printf("%d\n", sum)
}
