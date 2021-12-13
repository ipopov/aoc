package main

import "fmt"
import "strconv"
import "strings"
import "os"
import "bufio"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums_str := strings.Split(scanner.Text(), ",")
	var nums []int
	for _, i := range nums_str {
		x, _ := strconv.ParseInt(i, 10, 32)
		nums = append(nums, int(x))
	}

	var fish []int = make([]int, 9)
	for _, i := range nums {
		fish[i] += 1
	}

	for i := 0; i < 256; i++ {
		today := fish[0]
		fish = append(fish[1:7], today+fish[7], fish[8], today)
	}

	var sum int
	for _, x := range fish {
		sum += x
	}
	fmt.Println(sum)
}
