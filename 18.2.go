package main

import "fmt"
import "math"
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


func descend_right(i iterator) iterator {
	for ; !i.n.isLeaf; {
    i.n = i.n.right
    i.depth++
	}
  return i
}
func descend_left(i iterator) iterator {
	for ; !i.n.isLeaf; {
    i.n = i.n.left 
    i.depth++
	}
  return i
}

func left(i iterator) iterator {
	p := i.n.parent
	if p == nil {
		return iterator{
      n: nil,
    }
	}
	for i.n == p.left {
		p, i.n = p.parent, p
    i.depth--
		if p == nil {
      return iterator{
        n: nil,
      }
		}
	}
  i.n = p.left
  return descend_right(i)
}

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
  l:= make_leaf(int(math.Floor(float64(n.value) / 2))      )
  r:= make_leaf(int(math.Ceil(float64(n.value) / 2)))
  l.parent = n
  r.parent = n
  n.isLeaf= false
  n.left = l
  n.right = r
  
  return true
}

func maybeExplode(i iterator) bool {
  if i.depth < 5 {
    return false
  }
  
  parent_it := iterator {
    n: i.n.parent,
    depth: i.depth -1,
  }
  left_neighbor := left(iterator {
    n: parent_it.n.left,
    depth: i.depth,
    })
  right_neighbor := right(iterator {
    n: parent_it.n.right,
    depth: i.depth,
    })
  if x := left_neighbor.n; x != nil {
    x.value += parent_it.n.left.value
  }
  if x := right_neighbor.n; x != nil {
    x.value += parent_it.n.right.value
  }
  parent_it.n.isLeaf = true
  parent_it.n.value = 0
  parent_it.n.left = nil
  parent_it.n.right = nil
  
  return true
}

func magnitude(n *node) int {
  if n.isLeaf {
    return n.value
  }
  return 3 * magnitude(n.left) + 2*magnitude(n.right)
}

func parseString(s []byte) *node {
	tokens := tokenize(s)
	n := parse(&tokens)
  return n
}

func (n *node) toString() string {
  if n.isLeaf {
    return fmt.Sprintf("%d", n.value)
  }
  return fmt.Sprintf("[%s,%s]", n.left.toString(), n.right.toString())
}

func add(x, y *node) *node {
  ret := make_node(x, y)
Outer:
  for did_reduction:=true; did_reduction; {
    root_iterator := iterator {
      n: ret,
      depth: 0,
    }
    did_reduction = false
    for x := descend_left(root_iterator); x.n != nil; x = right(x) {
      if maybeExplode(x) {
        did_reduction = true
        continue Outer
      }
    }
    for x := descend_left(root_iterator); x.n != nil; x = right(x) {
      if maybeSplit(x.n) {
        did_reduction = true
        continue Outer
      }
    }
  }
  fmt.Printf("%s\n", ret.toString())
  return ret
}

func main() {
  var inputs [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
    inputs = append(inputs, []byte(scanner.Text()))
  }
  max := 0
  for i_idx, i := range inputs {
    for j_idx, j := range inputs {
      if i_idx == j_idx {
        continue
        }
      m := magnitude(add(parseString(i), parseString(j)))
      if m>max {
       max = m
       }
    }
  }
  fmt.Printf("%d\n", max)
}
