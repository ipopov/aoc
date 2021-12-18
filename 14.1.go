package main

import "fmt"
import "os"
import "bufio"
import "strings"

func f(m map[string]byte, s string) string {
  var b strings.Builder
  b.WriteByte(s[0])
  for i := 1; i < len(s); i++ {
    if r, ok := m[s[i-1:i+1]]; ok {
      b.WriteByte(r)
    }
    b.WriteByte(s[i])
  }
  return b.String()
}

func max_element(m map[byte]int, compare func(int, int)bool ) byte {
  var max_t byte
  var max_v int 
  for k, v := range m {
    max_t = k
    max_v = v
    break
  }
  for k, v := range m {
    if compare(max_v, v) {
      max_t = k
      max_v = v
    }
  }
  return max_t
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
  m := map[string]byte {}
  scanner.Scan() 
  seed := scanner.Text()
  scanner.Scan() 
	for scanner.Scan() {
    var from string
    var to byte
		fmt.Sscanf(scanner.Text(), "%s -> %c", &from, &to)
    m[from] = to
	}
  for i := 0; i < 10; i++ {
    seed = f(m, seed)
  }
  cs := map[byte]int{}
  for _, c := range []byte(seed) {
    cs[c] += 1
  }
  max := max_element(cs, func(x, y int) bool { return x < y })
  min := max_element(cs, func(x, y int) bool { return x > y })

  fmt.Println(cs)
  fmt.Println(cs[max] - cs[min])
}
