package main

import "fmt"
import "os"
import "bufio"

func traverse(b []byte, i *int, outcome *int) {
	for *i != len(b) && *outcome == 0 {
		c := b[*i]
		starters := map[byte]byte{
			']': '[',
			'>': '<',
			')': '(',
			'}': '{',
		}
		scores := map[byte]int{
			']': 57,
			'>': 25137,
			')': 3,
			'}': 1197,
		}
		if _, ok := starters[c]; ok {
			return
		}
		*i++
		traverse(b, i, outcome)

		if *outcome != 0 {
			return
		}
		if *i == len(b) {
			return
		}
		if starters[b[*i]] != c {
			*outcome = scores[b[*i]]
		}
		*i++
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := [][]byte{}
	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
	}
	sum := 0
	for _, in := range input {
		i := 0
		outcome := 0
		traverse(in, &i, &outcome)
		sum += outcome
	}
	fmt.Println(sum)
}
