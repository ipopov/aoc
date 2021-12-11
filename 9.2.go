package main

import "fmt"
import "os"
import "bufio"

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	board := [][]byte{}
	for scanner.Scan() {
		var xxx []byte
		xxx = append(xxx, scanner.Bytes()...)
		board = append(board, xxx)
	}
	li, lj := len(board), len(board[0])

	for i := 0; i < li; i++ {
		for j := 0; j < lj; j++ {
			v := int(board[i][j] - '0')
			if v == 9 {
				fmt.Printf("9")
			} else {
				fmt.Printf("\033[48;5;232;38;5;%dm%c\033[0m", 234+v, '0'+v)
			}
		}
		fmt.Println()
	}
}
