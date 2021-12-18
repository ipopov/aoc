package main

import "fmt"
import "os"
import "bufio"

type state map[string]int

func to_state(s string)state {
  ret := state{}
  ret[string([]byte{'_', s[0]})] += 1
  for i := 1; i < len(s)-2; i++ {
    ret[s[i:i+2]] += 1
  }
  ret[string([]byte{s[len(s)-1], '_'})] += 1
  return ret
}

func g(m map[string]byte, s state) state {
  ret := state{}
  for k, v := range s {
    if c, ok := m[k]; ok {
      ret[string([]byte{k[0], c})] += v
      ret[string([]byte{c,k[1]})] += v
    } else {
      ret[k] += v
    }
  }
  return ret
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
  s := to_state(seed)
  for i := 0; i < 10; i++ {
    s = g(m, s)
  }
  fmt.Println(s)
  cs := map[byte]int{}
  for k, v := range s {
    cs[k[0]] += v
    cs[k[1]] += v
  }
  delete(cs, '_')
  max := max_element(cs, func(x, y int) bool { return x < y })
  min := max_element(cs, func(x, y int) bool { return x > y })

  fmt.Println(cs)
  fmt.Println((cs[max] - cs[min])/2)
}
