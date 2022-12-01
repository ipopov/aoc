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

func bisect(vals []string, target rune) int64 {
	var idx int

	lo, hi := 0, len(vals)

	char_index := 0
	for hi-lo > 1 {
		fmt.Println("lo", lo, "hi", hi)
		more_common := rune(vals[(lo+hi)/2][char_index])
		fmt.Printf("more_common %c\n", more_common)
		idx = search(func(x int) int {
			return int(vals[x][char_index] - byte('0'))
		}, lo, hi)
		fmt.Println("idx", idx)
		if more_common == target {
			hi = idx
		} else {
			lo = idx
		}
		char_index += 1
	}
	fmt.Println("lo", lo, "hi", hi)

	ret, _ := strconv.ParseInt(vals[lo], 2, 64)
	return ret
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var vals []string
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}

	sort.Strings(vals)

	oxy := bisect(vals, '1')
	co2 := bisect(vals, '0')

	fmt.Println(oxy * co2)
}
