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
}

func main() {
  newList := &LinkedList{}
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
