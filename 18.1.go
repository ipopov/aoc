package main

import "fmt"
import "os"
import "bufio"

type node struct {
	isLeaf             bool
	value               int
	parent, left, right *node
}

func make_node(l, r *node) *node {
	ret := &node{
		isLeaf: false,
		left:    l,
		right:   r,
	}
	ret.left.parent = ret
	ret.right.parent = ret
	return ret
}

type iterator struct {
  n *node
  depth int
}

func make_leaf(v int) *node {
	return &node{
		isLeaf: true,
		value:   v,
	}
}

func descend_left(i iterator) iterator {
	for ; !i.n.isLeaf; {
    i.n = i.n.left 
    i.depth++
	}
  return i
}

//func left(n *node) *node {
//	p := n.parent
//	if p == nil {
//		return nil
//	}
//	for n == p.left {
//		p, n = p.parent, p
//		if p == nil {
//			return nil
//		}
//	}
//	q := p.left
//	for ; !q.isLeaf; q = q.right {
//	}
//	return q
//}
//
func right(i iterator) iterator {
	p := i.n.parent
	if p == nil {
		return iterator{
      n: nil,
    }
	}
	for i.n == p.right {
		p, i.n = p.parent, p
    i.depth--
		if p == nil {
      return iterator{
        n: nil,
      }
		}
	}
  i.n = p.right
  return descend_left(i)
}

type tokenType int

const (
	Open  tokenType = 0
	Close           = 1
	Int             = 2
)

type token struct {
	t tokenType
	i int
}

func parse(tokens *[]token) *node {
	head := (*tokens)[0]
	switch head.t {
	case Int:
		ret := make_leaf(head.i)
		*tokens = (*tokens)[1:]
		return ret
	case Open:
		*tokens = (*tokens)[1:]
		l := parse(tokens)
		r := parse(tokens)
		*tokens = (*tokens)[1:]
		return make_node(l, r)
	}
	panic("")
}

func tokenize(in []byte) []token {
	ret := []token{}
	for len(in) != 0 {
		if in[0] == '[' {
			ret = append(ret, token{t: Open})
      in = in[1:]
      continue
		}
		if in[0] == ']' {
			ret = append(ret, token{t: Close})
      in = in[1:]
      continue
		}
		if in[0] == ',' {
      in = in[1:]
      continue
		}
    i := 0
    for in[0] >= '0' && in[0] <= '9' {
      i = i*10 + int(in[0] - '0')
      in = in[1:]
    }
		ret = append(ret, token{t: Int, i:i})
	}
	return ret
}

func maybeSplit(n *node) bool {
  if (n.value < 10) {
    return false
  }
  l:= make_leaf(n.value / 2)
  r:= make_leaf(n.value / 2)
  l.parent = n
  r.parent = n
  n.isLeaf= false
  n.left = l
  n.right = r
  
  return true
}

func main() {
	tokens := tokenize([]byte("[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]"))
	fmt.Printf("%u\n", tokens)
	n := parse(&tokens)
  root_iterator := iterator {
  n: n,
  depth: 0,
  }

	for x := descend_left(root_iterator); x.n != nil; x = right(x) {
		if maybeSplit(x.n) {
      break
    }
	}
	for x := descend_left(root_iterator); x.n != nil; x = right(x) {
		fmt.Printf("%d: %d\n", x.depth, x.n.value)
	}

	fmt.Printf("%u\n", n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var x, u, y, v int
	fmt.Sscanf(scanner.Text(), "target area: x=%d..%d, y=%d..%d", &x, &u, &y, &v)
	fmt.Println(x, u, y, v)
}
