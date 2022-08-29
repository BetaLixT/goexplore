package main

import (
	"fmt"
	"strconv"
)

type Node struct {
  PointedBy []*Node
  PointingTo []*Node
  Value string
}

func NewNode(value string) *Node {
  return &Node{
    PointedBy: []*Node{},
    PointingTo: []*Node{},
    Value: value,
  }
}

func (n *Node) AddParent(p *Node) {
  n.PointedBy = append(n.PointedBy, p)
  p.PointingTo = append(p.PointingTo, n)
}

func (n *Node) ListGeneration() (nodes []*Node) {
  nodes = []*Node{}
  check := map[*Node]struct{}{}
  for _, p := range n.PointedBy {
    for _, c := range p.PointingTo {
      if _, ok := check[c]; !ok {
        check[c] = struct{}{}
        nodes = append(nodes, c)
      }
    } 
  }
  return
}

func main() {
  userCount := 10
  classCount := 3
  
  // user index
  users := make([]*Node, userCount)
  for i := 0; i < userCount; i++ {
    users[i] = NewNode("user:" + strconv.Itoa(i))
  }

  // class index
  classes := make([]*Node, classCount)
  for idx := range classes {
    classes[idx] = NewNode("class:" + strconv.Itoa(idx))
  }

  for idx, u := range users { 
    if idx % 2 == 0 {
      u.AddParent(classes[0])
    }
    if idx % 3 == 0 {
      u.AddParent(classes[1])
    }
    if idx % 2 != 0 && idx % 3 != 0 {
      u.AddParent(classes[2])
    }
  }

  n1 := users[3].ListGeneration()
  for _, x := range n1 {
    fmt.Printf("%s ", x.Value)
  }
  n2 := users[6].ListGeneration()
  fmt.Printf("\n")
  for _, x := range n2 {
    fmt.Printf("%s ", x.Value)
  }
  n3 := users[1].ListGeneration()
  fmt.Printf("\n")
  for _, x := range n3 {
    fmt.Printf("%s ", x.Value)
  }
}
