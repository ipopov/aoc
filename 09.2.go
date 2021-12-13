package main

import "fmt"
import "sort"
import "os"
import "bufio"

type index struct {
	x, y int
}

func lowest(b [][]byte, i, j int) bool {
	li, lj := len(b), len(b[0])
	for _, x := range []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		ii, jj := i+x.x, j+x.y
		if ii < 0 || ii >= li || jj < 0 || jj >= lj {
			continue
		}
		if b[i][j] >= b[ii][jj] {
			return false
		}
	}
	return true
}

func traverse(b [][]byte, i, j int, count *int) {
	if i < 0 || i >= len(b) || j < 0 || j >= len(b[0]) {
		return
	}
	if b[i][j] == '9' {
		return
	}
	b[i][j] = '9'
	*count += 1
	for _, x := range []index{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		traverse(b, i+x.x, j+x.y, count)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	board := [][]byte{}
	for scanner.Scan() {
		var xxx []byte
		xxx = append(xxx, scanner.Bytes()...)
		board = append(board, xxx)
	}
	li, lj := len(board), len(board[0])

	sinks := []index{}
	for i := 0; i < li; i++ {
		for j := 0; j < lj; j++ {
			if lowest(board, i, j) {
				sinks = append(sinks, index{i, j})
			}
		}
	}
	sizes := []int{}
	for _, x := range sinks {
		sizes = append(sizes, 0)
		traverse(board, x.x, x.y, &sizes[len(sizes)-1])
	}
	sort.Ints(sizes)
	result := 1
	for x := len(sizes) - 3; x <= len(sizes)-1; x++ {
		result *= sizes[x]
	}
	fmt.Println(sizes)
	fmt.Println(result)
}
