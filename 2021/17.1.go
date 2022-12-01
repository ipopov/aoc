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

  allowable_times_discrete := map[int]bool{}
  allowable_times_min := compute_min{}

  // x bounds
  for vx_start := 0; vx_start < u; vx_start++ {
    vx := vx_start
    pos := 0
    for t := 0;; {
      if pos > x && pos < u {
        allowable_times_discrete[t] = true
        if vx == 0 {
          allowable_times_min.add(t)
        }
      }
      if vx == 0 || pos > u {
        break
      }
      t++
      pos += vx
      vx--
    }
  }

  fmt.Println(allowable_times_discrete)
  fmt.Println(allowable_times_min)

  highest_y_velocity := compute_min{}
  // y bounds
  for vy_start := 0; vy_start < -y; vy_start++ {
    vy := vy_start
    pos := 0
    for t := 0;; {
      if pos > -v && pos < -y {
        if _, ok := allowable_times_discrete[t]; ok {
          highest_y_velocity.add(vy_start)
        } else if allowable_times_min.seen_first && t >= allowable_times_min.min {
          highest_y_velocity.add(vy_start)
        }
        symmetric_t := t + 2 * vy_start
        if _, ok := allowable_times_discrete[symmetric_t]; ok {
          highest_y_velocity.add(-vy_start)
        } else if allowable_times_min.seen_first && symmetric_t >= allowable_times_min.min {
          highest_y_velocity.add(-vy_start)
        }
      }
      if pos > -y {
        break
      }
      t++
      pos += vy
      vy++
    }
  }
  fmt.Println(highest_y_velocity)
  if !highest_y_velocity.seen_first || highest_y_velocity.min >= 0 {
    panic("")
  }
  ret := 0
  for i := 0; i <= - highest_y_velocity.min; i++ {
    ret += i
  }
  fmt.Println(ret)
}
