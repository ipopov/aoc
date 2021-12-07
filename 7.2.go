package main

import "fmt"
import "strconv"
import "strings"
import "os"
import "bufio"
import "sort"

func abs(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums_str := strings.Split(scanner.Text(), ",")
	var nums []float64
	for _, i := range nums_str {
		x, _ := strconv.ParseFloat(i, 64)
		nums = append(nums, float64(x))
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	min, max := nums[0], nums[len(nums)-1]

	var min_cost float64 = -1
  // Go over all possible midpoints
	for i := min; i <= max; i++ {
		var cost float64
		for _, n := range nums {
			dist := abs(n - i)
			cost += abs(dist * (dist + 1) / 2)
		}
		if min_cost == -1 || cost < min_cost {
			min_cost = cost
		}
	}

	fmt.Printf("%20.0f\n", min_cost)
}
