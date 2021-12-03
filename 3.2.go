package main

import "fmt"
import "strconv"
import "os"
import "bufio"
import "sort"

func search(arr func(int) int, lo, hi int) int {
	if arr(lo) == 1 {
		return lo
	}
	if lo+1 == hi {
		return lo + 1
	}

	mid := (hi + lo) / 2

	if arr(mid) == 0 {
		return search(arr, mid, hi)
	} else {
		return search(arr, lo, mid)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var vals []string
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}

	sort.Strings(vals)

	var idx int

	lo, hi := 0, len(vals)
	// Start out chasing "most"
	char_index := 0
	for hi-lo > 1 {
		fmt.Println("lo", lo, "hi", hi)
		more_common := rune(vals[(lo+hi)/2][char_index])
		fmt.Printf("more_common %c\n", more_common)
		idx = search(func(x int) int {
			return int(vals[x][char_index] - byte('0'))
		}, lo, hi)
		fmt.Println("idx", idx)
		if more_common == '0' {
			hi = idx
		} else {
			lo = idx
		}
		char_index += 1
	}
	fmt.Println("lo", lo, "hi", hi)

	oxy, _ := strconv.ParseInt(vals[lo], 2, 64)

	lo, hi = 0, len(vals)
	// Start out chasing "least"
	char_index = 0
	for hi-lo > 1 {
		fmt.Println("lo", lo, "hi", hi)
		more_common := rune(vals[(lo+hi)/2][char_index])
		fmt.Printf("more_common %c\n", more_common)
		idx = search(func(x int) int {
			return int(vals[x][char_index] - byte('0'))
		}, lo, hi)
		fmt.Println("idx", idx)
		if more_common == '0' {
			lo = idx
		} else {
			hi = idx
		}
		char_index += 1
	}
	fmt.Println("lo", lo, "hi", hi)

	co2, _ := strconv.ParseInt(vals[lo], 2, 64)

	fmt.Println(oxy * co2)
}
