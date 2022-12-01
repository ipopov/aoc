package main

import (
	"bufio"
	//"container/heap"
	"fmt"
	"os"
)

////}	inf  bool
////}	dist int
////}}
//
//type PQ []point
//
//func (h PQ) Len() int {
//	return len(h)
//}
//func (h PQ) Less(i, j int) bool {
//  if h[i].inf != h[j].inf {
//    return h[i].inf < h[j].inf
//  }
//	return h[i].dist < h[j].dist
//}
//func (h PQ) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *PQ) Push(x interface{}) {
//	*h = append(*h, x.(traversal))
//}
//
//func (h *PQ) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}

type board [][]int

type point struct {
	x, y int
}

type Done struct {
	m map[point]int
}

func (d Done) Contains(p point) bool {
	_, ok := d.m[p]
	return ok
}

func (d *Done) Put(p point, c int) {
	d.m[p] = c
}

type Heap struct {
	m map[point]int // to cost
}

func (h Heap) Empty() bool {
	return len(h.m) == 0
}

func (h *Heap) Update(p point, c int) {
	it, ok := h.m[p]
	if !ok || c < it {
		h.m[p] = c
	}
}

func (h *Heap) PopMin() (point, int) {
	// Yes... this should be a heap. But it's a pain in Go.
	p := point{}
	c := -1
	for k, v := range h.m {
		if c == -1 || v < c {
			p = k
			c = v
		}
	}
	delete(h.m, p)
	return p, c
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
	h := Heap{
		m: map[point]int{},
	}
	h.Update(point{0, 0}, 0)
	for !h.Empty() {
		n, c := h.PopMin()
		d.Put(n, c)
		for _, neighbor := range neighbors(n.x, n.y, len(b), len(b[0])) {
			if d.Contains(neighbor) {
				continue
			}
			h.Update(neighbor, c+b[neighbor.x][neighbor.y])
		}
	}
	//fmt.Println(d)
	fmt.Println(d.m[point{len(b) - 1, len(b[0]) - 1}])
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
	//li, lj := len(b), len(b[0])

	//for i := 0; i < li; i++ {
	//	for j := 0; j < lj; j++ {
	//		fmt.Print(b[i][j])
	//	}
	//	fmt.Println()
	//}
	dijkstra(b)
}
