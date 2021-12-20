package main

import "fmt"
import "os"
import "bufio"

type BitString struct {
	x   []int8
	pos int
}

func (bs *BitString) Bit() int8 {
	ret := bs.x[bs.pos]
	bs.pos++
	return ret
}

func (bs *BitString) Bits(w int) int {
	var ret int
	for x := 0; x < w; x++ {
    ret <<= 1
		ret |= int(bs.Bit())
	}
	return ret
}

func (bs *BitString) ParseLiteral() int {
  var ret int
  for ;; {
    nibble := bs.Bits(5)
    ret <<= 4
    ret |= nibble & 0xf
    if nibble & 0x10 == 0 {
      break
    }
  }
  return ret
}

func (bs *BitString) ParsePacket() int {
  ret := 0
  version := bs.Bits(3)
  packet_type := bs.Bits(3)
  fmt.Println(version, "version", packet_type, "type")
  if packet_type == 4 {
    l := bs.ParseLiteral()
    ret = l
  } else {
    children := []int{}  // versions
    if length_type := bs.Bit(); length_type == 0 {
      // TODO: Don't reach in for "pos" here.
      x:= bs.Bits(15);
      fmt.Println(x, "bits for children")
      for l_bits := bs.pos + x; l_bits > bs.pos; {
        children = append(children, bs.ParsePacket())
      }
    } else {
      x:= bs.Bits(11);
      fmt.Println(x, "count for children")
      for ; x > 0; x-- {
        children = append(children, bs.ParsePacket())
      }
    }
    switch packet_type {
      case 0:
      for _, x := range children {
        ret += x
      }
      case 1:
      ret = 1
      for _, x := range children {
        ret *= x
      }
      case 2:
      ret = children[0]
      for _, x := range children[1:] {
        if x < ret {
          ret = x
        }
      }
      case 3:
      ret = children[0]
      for _, x := range children[1:] {
        if x > ret {
          ret = x
        }
      }
      case 5:
      if children[0] > children[1] {
        ret = 1
      } else {
        ret = 0
      }
      case 6:
      if children[0] < children[1] {
        ret = 1
      } else {
        ret = 0
      }
      case 7:
      if children[0] == children[1] {
        ret = 1
      } else {
        ret = 0
      }
    }
  }
  return ret
}

func main() {
  bs := BitString{
    x : []int8{},
  }
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := []byte(scanner.Text())
  for _, x := range input {
    v := int8(0)
    if x >= 'A' && x <= 'F' {
      v = int8(10 + x - 'A')
    } else {
      v = int8(x - '0')
    }
    for i := 3; i >= 0; i-- {
      bs.x = append(bs.x, (v >>i)  & 1)
    }
  }
	fmt.Println(bs)
  // D2FE28
	fmt.Println(bs.ParsePacket())
}
