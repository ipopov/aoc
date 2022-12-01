package main

import "fmt"
import "sort"
import "os"
import "bufio"

func traverse(b []byte, i *int, fault, completion *int) {
	for *i != len(b) && *fault == 0 {
		c := b[*i]
		starters := map[byte]byte{
			']': '[',
			'>': '<',
			')': '(',
			'}': '{',
		}
		if _, ok := starters[c]; ok {
			return
		}
		*i++
		traverse(b, i, fault, completion)

		if *fault != 0 {
			return
		}
		if *i == len(b) {
			scores := map[byte]int{
				'[': 2,
				'<': 4,
				'(': 1,
				'{': 3,
			}
			*completion = *completion*5 + scores[c]
			return
		}
		if starters[b[*i]] != c {
			scores := map[byte]int{
				']': 57,
				'>': 25137,
				')': 3,
				'}': 1197,
			}
			*fault = scores[b[*i]]
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
	fault_sum := 0
	nonzero := []int{}
	for _, in := range input {
		i := 0
		fault := 0
		completion_score := 0
		traverse(in, &i, &fault, &completion_score)
		fault_sum += fault
		if completion_score != 0 {
			nonzero = append(nonzero, completion_score)
		}
	}
	sort.Ints(nonzero)
	fmt.Println(nonzero[len(nonzero)/2])
}
