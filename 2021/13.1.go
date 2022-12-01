package main

import "fmt"
import "constraints"
import "os"
import "bufio"

type point struct {
	x, y int
}

type fold struct {
	axis  byte
	value int
}

func abs[T constraints.Signed](x T) T {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func f(ps []point, fs []fold) {
	f := fs[0]
	m := map[point]bool{}
	for _, p := range ps {
		var z *int
		if f.axis == 'x' {
			z = &p.x
		} else {
			z = &p.y
		}
		*z = f.value - abs(f.value-*z)
		m[p] = true
	}
	fmt.Println(len(m))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ps := []point{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		p := point{}
		fmt.Sscanf(scanner.Text(), "%d,%d", &p.x, &p.y)

		ps = append(ps, p)
	}
	fs := []fold{}
	for scanner.Scan() {
		f := fold{}
		fmt.Sscanf(scanner.Text(), "fold along %c=%d", &f.axis, &f.value)
		fs = append(fs, f)
	}
	f(ps, fs)

	// fmt.Println(ps)
	// fmt.Println(fs)
}
