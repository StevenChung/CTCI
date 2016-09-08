/*
Return Kth to Last: Implement an
algorithm to find the kth to last
element of a singly linked list.

Note: You implemented a Length property on your list, so
we will attempt to handle this sans Length (trivial)
*/

package main

import (
  "fmt"
  "errors"
)
// while-loop solution
func (l *DoublyLinkedList) RetrieveKthToLast(k int) (*Node, error) {
  var NodeToReturn *Node
  currentNode := l.Head
  var err error

  for currentNode != nil {
    checkIfTail := currentNode
    for i := 0; i < k; i++ {
      if checkIfTail.Next != nil {
        checkIfTail = checkIfTail.Next
      } else {
        err = errors.New("Invalid k")
        break
      }
    }

    if checkIfTail.Value.Numpy == l.Tail.Value.Numpy {
      NodeToReturn = currentNode
      break
    }
    currentNode = currentNode.Next
  }

  if currentNode == nil {
    err = errors.New("Invalid k")
  }
  return NodeToReturn, err
}

func main() {
  newList := &DoublyLinkedList{}
  // memory allocation/value intialization
  // of *DoublyLinkedList type
  // pointer methods!
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 7})
  newList.AddToTail(Value{"Testy", 2})
  newList.AddToHead(Value{"Test", 14})
  newList.InsertAtIndex(4, Value{"Test", 7})


  for i := newList.Length - 1; i >= 0; i-- {
    fmt.Println(newList.RetrieveKthToLast(i))
  }

}
