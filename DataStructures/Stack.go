package main

import (
  "fmt"
  "errors"
)

type Node struct {
  Val int
  Next *Node
}

type Stack struct {
  Last *Node
}

func (s *Stack) Add(v int) *Stack {
  newNode := &Node{Val: v}
  if s.Last == nil {
    s.Last = newNode
  } else {
    newNode.Next = s.Last
    s.Last = newNode
  }
  return s
}

func (s *Stack) Remove() (*Node, error) {
  // LIFO
  var err error
  var item *Node
  if s.Last == nil {
    err = errors.New("Empty Stack")
  } else {
    item = s.Last
    if s.Last.Next != nil {
      s.Last = s.Last.Next
    } else {
      s.Last = nil
    }
  }
  return item, err
}

func (s *Stack) Peek() (*Node, error) {
  // returns the top of the stack
  var returnval *Node
  var err error
  if s.Last != nil {
    returnval = s.Last
  } else {
    err = errors.New("Empty stack")
  }
  return returnval, err
}

func (s *Stack) IsEmpty() bool {
  // returns the top of the stack
  if s.Last == nil {
    return true
  } else {
    return false
  }
}

func main() {
  newStack := &Stack{}
  newStack.Add(5)
  newStack.Add(7)
  fmt.Println(newStack.Peek())
  fmt.Println(newStack.Last)
}
