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
		if x != u && y != v {
			continue
		}
		for i := min(x, u); i <= max(x, u); i++ {
			for j := min(y, v); j <= max(y, v); j++ {
				s[i][j] += 1
			}
		}
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
