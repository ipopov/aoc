package main

import "fmt"
import "strconv"
import "strings"
import "os"
import "bufio"

type card struct {
	x [5][5]int
}

func subset(x, y []int) bool {
	var y_map map[int]bool = map[int]bool{}
	for _, x := range y {
		y_map[x] = true
	}
	for _, x := range x {
		if !y_map[x] {
			return false
		}
	}
	return true
}

func f(c card, nums []int) bool {
	for _, x := range c.x {
		if subset(x[:], nums) {
			return true
		}
	}
	for i := 0; i < 5; i++ {
		if subset([]int{c.x[0][i], c.x[1][i], c.x[2][i], c.x[3][i], c.x[4][i]}, nums) {
			return true
		}
	}
	return false
}

func score(c card, nums []int) int {
	var score int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !subset([]int{c.x[i][j]}, nums) {
				score += c.x[i][j]
			}
		}
	}
	return score
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums_str := strings.Split(scanner.Text(), ",")
	var nums []int
	for _, i := range nums_str {
		x, _ := strconv.ParseInt(i, 10, 32)
		nums = append(nums, int(x))
	}
	fmt.Println(nums)

	var cs []card
	for scanner.Scan() {
		var x [5][5]int
		for i := 0; i < 5; i++ {
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &x[i][0], &x[i][1], &x[i][2], &x[i][3], &x[i][4])
		}
		cs = append(cs, card{x: x})
	}

	for i := 0; i < len(nums); i++ {
		for _, card := range cs {
			if f(card, nums[0:i]) {
				fmt.Println(nums[i-1] * score(card, nums[0:i]))
				return
			}
		}
	}

}
