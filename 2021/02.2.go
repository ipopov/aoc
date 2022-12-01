package main

import "fmt"
import "os"
import "bufio"

type state struct {
	x, y int
	aim  int
}

var foo map[string]func(*state, int)

func (x *state) mutate(cmd string, val int) {
	f := foo[cmd]
	f(x, val)
}

func main() {
	foo = map[string]func(*state, int){
		"forward": func(x *state, v int) {
			x.x += v
			x.y += x.aim * v
		},
		"down": func(x *state, v int) { x.aim += v },
		"up":   func(x *state, v int) { x.aim -= v },
	}
	var s state
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var cmd string
		var val int
		fmt.Sscanf(scanner.Text(), "%s %d", &cmd, &val)
		s.mutate(cmd, val)
	}
	fmt.Println(s)
	fmt.Println(s.x * s.y)
}
