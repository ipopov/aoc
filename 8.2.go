package main

import "fmt"
import "sort"
import "strings"
import "os"
import "bufio"

type perm struct {
	x string
}

func generate_perms(alphabet string) []perm {
	if len(alphabet) == 0 {
		return []perm{perm{x: ""}}
	}
	var ret []perm
	for _, alpha := range alphabet {
		cp := strings.ReplaceAll(alphabet, string(alpha), "")
		for _, x := range generate_perms(cp) {
			x.x = x.x + string(alpha)
			ret = append(ret, x)
		}
	}
	return ret
}

func (p *perm) remap(initial string) string {
	rs := []rune(initial)
	mapping := []rune(p.x)
	for i, x := range rs {
		rs[i] = mapping[x-'a']
	}
	sort.Slice(rs, func(i, j int) bool { return rs[i] < rs[j] })
	return string(rs)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	valid := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

  perms := generate_perms("abcdefg")
	var x int
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " | ")
		left := strings.Split(s[0], " ")
		right := strings.Split(s[1], " ")
	BTarget:
		for _, p := range perms {
			for _, l := range left {
				rl := p.remap(l)
				if _, ok := valid[rl]; !ok {
					continue BTarget
				}
			}
			var this_result int
			for _, r := range right {
				digit := valid[p.remap(r)]
				this_result = this_result*10 + digit
			}
			x += this_result
		}
	}

	fmt.Println(x)
}
