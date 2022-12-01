package main

import "fmt"
import "strings"
import "os"
import "bufio"

type graph map[string][]string

type visited struct {
	xs []string
}

func (v *visited) add(s string) {
	if s[0] >= 'A' && s[0] <= 'Z' {
		return
	}
	v.xs = append(v.xs, s)
}

func (v *visited) contains(s string) bool {
	for _, x := range v.xs {
		if x == s {
			return true
		}
	}
	return false
}

func traverse(g *graph, start string, v visited) int {
	if start == "end" {
		return 1
	}

	ret := 0
	v.add(start)
	for _, n := range (*g)[start] {
		if v.contains(n) {
			continue
		}
		ret += traverse(g, n, v)
	}
	return ret
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	g := graph{}
	//var x int
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "-")
		g[s[0]] = append(g[s[0]], s[1])
		g[s[1]] = append(g[s[1]], s[0])
	}
	fmt.Println(g)
	fmt.Println(traverse(&g, "start", visited{}))

}
