package main

import "fmt"
import "os"
import "bufio"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	const (
		kX = 1000
		kY = 1000
	)
	var s [kX][kY]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var x, y, u, v int
		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x, &y, &u, &v)
		i, j := x, y
		for i != u || j != v {
			s[i][j] += 1
			if i < u {
				i++
			} else if i > u {
				i--
			}
			if j < v {
				j++
			} else if j > v {
				j--
			}
		}
		s[i][j] += 1
	}
	var count int
	for i := 0; i < kX; i++ {
		for j := 0; j < kY; j++ {
			if s[i][j] >= 2 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
