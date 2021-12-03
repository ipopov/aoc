package main

import "fmt"
import "os"
import "bufio"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var a []int
	for scanner.Scan() {
		if a == nil {
			a = make([]int, len(scanner.Text()))
		}
		for i, digit := range scanner.Text() {
			if digit == '0' {
				a[i] -= 1
			} else {
				a[i] += 1
			}
		}
	}

  var gamma, epsilon int
	for _, digit := range a {
    gamma <<= 1
    epsilon <<= 1
    if digit > 0 {
      gamma |= 1
    } else {
      epsilon |= 1
    }
  }

	fmt.Println(a)
	fmt.Println(gamma, epsilon)
	fmt.Println(gamma * epsilon)
}
