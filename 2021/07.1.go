package main

import "fmt"
import "strconv"
import "strings"
import "os"
import "bufio"
import "sort"

func abs(x float32) float32 {
	if x < 0 {
		return -x
	} else {
		return x
	}

}

func median(x []float32) float32 {
	l := len(x)
	if l%2 == 1 {
		return x[l/2]
	} else {
		return (x[l/2] + x[l/2-1]) / 2
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums_str := strings.Split(scanner.Text(), ",")
	var nums []float32
	for _, i := range nums_str {
		x, _ := strconv.ParseFloat(i, 32)
		nums = append(nums, float32(x))
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	m := median(nums)
	var sum float32
	for _, x := range nums {
		sum += abs(m - x)
	}

	fmt.Println(sum)
}
