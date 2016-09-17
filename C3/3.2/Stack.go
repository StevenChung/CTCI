package main

import (
  "errors"
  // "fmt"
)

type Node struct {
  Val int
  Next *Node
}

type Stack struct {
  Last *Node
  MinimumStack *MinStack // recall that this is initialized as nil
  LocalMin *Node
}
type MinStack struct {
  Last *Node
}

func (ms *MinStack) Addy(v int) *MinStack {
  newNode := &Node{Val: v}
  if ms.Last == nil {
    ms.Last = newNode
  } else {
    newNode.Next = ms.Last
    ms.Last = newNode
  }
  return ms
}

func (ms *MinStack) Removey() (*Node, error) {
  // LIFO
  var err error
  var item *Node
  if ms.Last == nil {
    err = errors.New("Empty Stack")
  } else {
    item = ms.Last
    if ms.Last.Next != nil {
      ms.Last = ms.Last.Next
    } else {
      ms.Last = nil
    }
  }
  return item, err
}


func (s *Stack) Add(v int) *Stack {
  newNode := &Node{Val: v}
  if s.Last == nil {
    s.Last = newNode
    s.LocalMin = newNode
    s.MinimumStack = &MinStack{}
    // initialize place in memory
    s.MinimumStack.Addy(v)
  } else {
    newNode.Next = s.Last
    s.Last = newNode
    if s.LocalMin.Val > newNode.Val {
      s.MinimumStack.Addy(v)
      s.LocalMin = newNode
    }
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
