package main

import "fmt"
import "os"
import "bufio"

func print_board(in [][]int) {
	for _, x := range in {
		for _, y := range x {
			fmt.Print(y)
		}
		fmt.Println()
	}
}

func update[T any](b [][]T, f func(x, y int, v *T)) {
	for i, x := range b {
		for j, _ := range x {
			f(i, j, &b[i][j])
		}
	}
}

type point struct {
	x, y int
}

func neighbors(i, j int) []point {
	ret := []point{}
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 || i+x < 0 || i+x > 9 || j+y < 0 || j+y > 9 {
				continue
			}
			ret = append(ret, point{i + x, j + y})
		}
	}
	return ret
}

func propagate(b [][]int, x, y int) int {
	ret := 0
	b[x][y]++
	if b[x][y] == 10 { // also a termination condition
		ret++
		for _, point := range neighbors(x, y) {
			ret += propagate(b, point.x, point.y)
		}
	}
	return ret
}

func step(b [][]int) (int, bool) {
	flashes := 0
	for i, x := range b {
		for j, _ := range x {
			flashes += propagate(b, i, j)
		}
	}
	update(b, func(i, j int, b *int) {
		if *b > 9 {
			*b = 0
		}
	})
	all_zero := true
	update(b, func(i, j int, b *int) {
		if *b != 0 {
			all_zero = false
		}
	})
	return flashes, all_zero
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	b := [][]int{}
	for scanner.Scan() {
		q := func(in []byte) []int {
			ret := make([]int, len(in))
			for i, b := range in {
				ret[i] = int(b - '0')
			}
			return ret
		}
		b = append(b, q(scanner.Bytes()))
	}
	i := 0
	for ; ; i++ {
		_, all_zero := step(b)
		if all_zero {
			break
		}
	}
	fmt.Println(i + 1)
}
