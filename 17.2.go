package main

import "fmt"
import "os"
import "bufio"

type compute_min struct {
  seen_first bool
  min int
}
func (x *compute_min)add(i int ) {
  if !x.seen_first {
    x.seen_first = true
    x.min = i
  }
  if i < x.min {
    x.min = i
  }
}

type compute_max struct {
  seen_first bool
  max int
}
func (x *compute_max)add(i int) {
  if !x.seen_first {
    x.seen_first = true
    x.max = i
  }
  if i > x.max {
    x.max = i
  }
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan() 
  var x, u, y, v int
  fmt.Sscanf(scanner.Text(), "target area: x=%d..%d, y=%d..%d", &x, &u, &y, &v)
	fmt.Println(x, u, y, v)

  type pair struct {
    x, y int
  }
  m := map[pair]bool{}
  for vy := y; vy <= -y; vy++ {
    for vx := 0; vx <= u; vx++ {
      vx_i, vy_i := vx, vy
      x_i, y_i := 0, 0
      for t := 0;; t++ {
        x_i += vx_i
        y_i += vy_i
        if vx_i != 0 {
          vx_i--
        }
        vy_i--
        if x_i >= x && x_i <= u && y_i >= y && y_i <= v {
          m[pair{vx, vy}] = true
        }
        if y_i < y {
          break
        }
      }
    }
  }
  fmt.Println(m)
  fmt.Println(len(m))
}
