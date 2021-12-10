package main

import "fmt"
import "strings"
import "os"
import "bufio"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var x int
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " | ")
		//left := strings.Split(s[0], " ")
		right := strings.Split(s[1], " ")
		for _, q := range right {
			l := len(q)
			if l == 2 || l == 3 || l == 4|| l == 7  {
				x++
			}
		}
	}
	fmt.Println(x)
}
