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

type Outcome int

const (
	Lose Outcome = iota
	Draw
	Win
)

func (p *Play) Capture(s []string) error {
	switch s[0] {
	case "A":
		*p = Rock
		return nil
	case "B":
		*p = Paper
		return nil
	case "C":
		*p = Scissors
		return nil
	}
	return errors.New("blah")
}

func (p *Outcome) Capture(s []string) error {
	switch s[0] {
	case "X":
		*p = Lose
		return nil
	case "Y":
		*p = Draw
		return nil
	case "Z":
		*p = Win
		return nil
	}
	return errors.New("blah")
}

type Round struct {
	OpponentPrediction *Play    `@PlayA Space`
	Recommendation     *Outcome `@PlayB Newline`
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

func Recommendation(opponent Play, p Outcome) Play {
	switch p {
	case Win:
		return (opponent + 1) % 3
	case Draw:
		return opponent
	case Lose:
		return (opponent + 2) % 3
	}
	panic("")
}

func MatchScore(p Outcome) int {
	switch p {
	case Win:
		return 6
	case Draw:
		return 3
	case Lose:
		return 0
	}
	panic("")
}

func main() {
	in := parse(os.Stdin)
	sum := 0
	for _, r := range in.R {
		sum += MatchScore(*r.Recommendation) + BaseScore(Recommendation(*r.OpponentPrediction, *r.Recommendation))
	}
	fmt.Printf("%d\n", sum)
}
