/*
Reverse a linked list
*/
package main

import (
  "fmt"
)
// iterative solution
// func (l *LinkedList) Reverse() {
//   var NewNext *Node
//   currentNode := l.Head
//   for currentNode != nil {
//     oldNext := currentNode.Next
//     currentNode.Next = NewNext
//     NewNext = currentNode
//     if oldNext != nil {
//       currentNode = oldNext
//     } else {
//       l.Head = currentNode
//       break
//     }
//   }
// }

// recursive solution
func (l *LinkedList) RecursiveReverse(n *Node) {
  if n.Next == nil {
    l.Head = n
    return
  }

  l.RecursiveReverse(n.Next)
  oldNext := n.Next
  oldNext.Next = n
  n.Next = nil
  // the old next is equal to n's next
  // the old next's next is reversed (set to n)
  // n's next is set to nil
  // at the head, this would mean it become the tail (i.e. not have a next)
}

func main() {
  newList := &LinkedList{}
  // new vs &LinkedList{} vs just starting a variable with type *LinkedList
  // new and &LinkedList{} saves memory with nil values
  // new variable with type *LinkedList is a nil reference
  newList.Append(2)
  newList.Append(3)
  newList.Append(4)

  newList.RecursiveReverse(newList.Head)

  currentNode := newList.Head
  for currentNode != nil {
    fmt.Println(currentNode)
    currentNode = currentNode.Next
  }
}
