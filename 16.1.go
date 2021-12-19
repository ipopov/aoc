package main

import "fmt"
import "os"
import "bufio"

type BitString struct {
	x   []int8
	pos int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := [][]byte{}
	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
	}
	sum := 0
	for _, _ = range input {
		outcome := 0
		sum += outcome
	}
	fmt.Println(sum)
}
