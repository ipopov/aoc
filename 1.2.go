package main

import "fmt"
import "os"
import "bufio"
import "strconv"

type stream interface {
	next() bool
	get() float32
}

type intStream struct {
	i       int
	scanner *bufio.Scanner
}

func (x *intStream) get() float32 {
	return float32(x.i)
}

func (x *intStream) next() bool {
	if !x.scanner.Scan() {
		return false
	}
	i, err := strconv.Atoi(x.scanner.Text())
	if err != nil {
		return false
	}
	x.i = i
	return true
}

type averagingStream struct {
	delegate stream
	buf      [3]float32
	next_idx int
	ready    bool
}

func (x *averagingStream) get() float32 {
	k := 3
	var sum float32 = 0
	for _, x := range x.buf {
		sum += x
	}
	return sum / float32(k)
}

func (x *averagingStream) next() bool {
	if !x.ready {
		for i := 0; i < 3; i++ {

			if !x.delegate.next() {
				return false
			}

			x.buf[i] = x.delegate.get()
		}
		x.ready = true
		return true
	}
	if !x.delegate.next() {
		return false
	}
	x.buf[x.next_idx] = x.delegate.get()
	x.next_idx = (x.next_idx + 1) % 3
	return true
}
func main() {
	it := averagingStream{
		delegate: &intStream{
			i:       0,
			scanner: bufio.NewScanner(os.Stdin),
		},
		next_idx: 0,
		ready:    false,
	}
	first := true
	var prev float32
	count := 0
	for it.next() {
		i := it.get()
		if first {
			first = false
			prev = i
			continue
		}
		if i > prev {
			count++
		}
		prev = i
	}
	fmt.Println("result: ", count)
}
