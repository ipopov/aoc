package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

type pq_entry struct {
  p point
  c int
}

type PQ []pq_entry

func (h PQ) Len() int {
	return len(h)
}
func (h PQ) Less(i, j int) bool {
	return h[i].c < h[j].c
}
func (h PQ) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PQ) Push(x any) {
	*h = append(*h, x.(pq_entry))
}

func (h *PQ) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type board [][]int

type Done struct {
  m map[point]int
}

func (d Done)Contains(p point) bool {
  _, ok := d.m[p]
  return ok
}

func (d *Done)Put(p point, c int) {
  d.m[p] = c
}

type Heap struct {
  q PQ
}

func (h *Heap)Init() {
  h.q = PQ{}
  heap.Init(&h.q)
}

func (h Heap)Empty() bool {
  return len(h.q) == 0
}

func (h *Heap)Update(p point, c int) {
  // Just push (duplicates are ok)
  heap.Push(&h.q, pq_entry{
    p: p,
    c: c,
  })
}

func (h *Heap)PopMin() (point, int) {
  pqe := heap.Pop(&h.q).(pq_entry)
  return pqe.p, pqe.c
}

func neighbors(i, j, dim_i, dim_j int) []point {
	ret := []point{}
	for _, x := range []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		ii, jj := i+x.x, j+x.y
		if ii < 0 || ii >= dim_i || jj < 0 || jj >= dim_j {
			continue
		}
		ret = append(ret, point{ii, jj})
  }
  return ret
}

func dijkstra(b board) {
  d := Done{
    m: map[point]int{},
  }
  h := Heap{}
  h.Init()
  h.Update(point{0, 0}, 0)
  for !h.Empty() {
    n, c := h.PopMin()
    if (d.Contains(n)) {
      continue
    }
    d.Put(n, c)
    for _, neighbor := range neighbors(n.x, n.y, len(b), len(b[0])) {
      if d.Contains(neighbor) {
        continue
      }
      h.Update(neighbor, c + b[neighbor.x][neighbor.y])
    }
  }
  //fmt.Println(d)
  fmt.Println(d.m[point{len(b)-1, len(b[0])-1}])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	b := board{}
	for scanner.Scan() {
		var xxx []int
		for _, x := range scanner.Bytes() {
			xxx = append(xxx, int(x-'0'))
		}
		b = append(b, xxx)
	}
	full_b := board{}
	li, lj := len(b), len(b[0])

	for i := 0; i < 5 * li; i++ {
    full_b = append(full_b, make([]int, 5*lj))
		for j := 0; j < 5 * lj; j++ {
      x_off, y_off := i / li, j / lj
      x_idx, y_idx := i % li, j % lj
      full_b[i][j] = (x_off + y_off + b[x_idx][y_idx]-1) % 9 + 1
    }
  }

  dijkstra(full_b)
}
