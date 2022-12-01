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

func max[T constraints.Ordered](x, y T) T {
	if x >= y {
		return x
	} else {
		return y
	}
}

func abs[T constraints.Signed](x T) T {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func f(ps []point, fs []fold) map[point]bool {
	m := map[point]bool{}
	for _, p := range ps {
		m[p] = true
	}
	for _, f := range fs {
		m_new := map[point]bool{}
		for p, _ := range m {
			var z *int
			if f.axis == 'x' {
				z = &p.x
			} else {
				z = &p.y
			}
			*z = f.value - abs(f.value-*z)
			m_new[p] = true
		}
		m = m_new
	}
	return m
}

func p(m map[point]bool) {
	x_size, y_size := 0, 0
	for p, _ := range m {
		x_size = max(x_size, p.x)
		y_size = max(y_size, p.y)
	}
	for y := 0; y < y_size+1; y++ {
		for x := 0; x < x_size+1; x++ {
			var b byte
			if _, ok := m[point{x, y}]; ok {
				b = '#'
			} else {
				b = ' '
			}
			fmt.Printf("%c", b)
		}
		fmt.Println()
	}
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
	p(f(ps, fs))
}
