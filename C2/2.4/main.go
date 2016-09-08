/*
Write code to partition a linked list around
a value x, such that all nodes less than x
come before all nodes greater than or equal to x.
lf x is contained within the list, the values of x
only need to be after the elements less than x (see below).
The partition element x can appear anywhere in the "right
partition"; it does not need to appear between the left and
right partitions.
*/

package main

import (
  "fmt"
)
// O(n) time // O(n) space
func (l *DoublyLinkedList) Partition(n int) *DoublyLinkedList {
  partitionedList := &DoublyLinkedList{}
  // iterate through linkedlist
  // slice tracking nodes that are greater than or equal to (if it is in the list)
  // slice tracking nodes that are less than
  currentNode := l.Head
  for currentNode != nil {
    if currentNode.Value.Numpy < n {
      partitionedList.AddToHead(Value{"test", currentNode.Value.Numpy})
    } else {
      partitionedList.AddToTail(Value{"test", currentNode.Value.Numpy})
    }
    currentNode = currentNode.Next
  }
  l = partitionedList
  return l

}

func main() {
  newList := &DoublyLinkedList{}
  // memory allocation/value intialization
  // of *DoublyLinkedList type
  // pointer methods!
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 2})
  newList.AddToHead(Value{"Test", 32})
  newList.AddToHead(Value{"Test", -4})
  newList.AddToHead(Value{"Test", 45})
  newList.AddToTail(Value{"Testy", 2})
  newList.AddToHead(Value{"Test", 14})
  newList.InsertAtIndex(4, Value{"Test", 7})

  newList = newList.Partition(20)
  for i := 0; i < newList.Length; i++ {
    fmt.Println(newList.RetrieveAtIndex(i))
  }
}
