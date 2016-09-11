/*
Write Code to Remove Duplicates from an UNSORTED LinkedList
Constant Space?
Single or Double?
*/

package main

import (
  "fmt"
)


func (l *DoublyLinkedList) RemoveDuplicates() {
  // O(1) / O(n) ------
  // hashtable := make(map[int]bool)
  // i := 0
  // for i < l.Length {
  //   node, _ := l.RetrieveAtIndex(i)
  //   if hashtable[node.Value.Numpy] {
  //     l.DeleteFromIndex(i)
  //   } else {
  //     hashtable[node.Value.Numpy] = true
  //     i++
  //   }
  // }


  // O(n^2) // O(1) -----

  for i := 0; i < l.Length; i++ {
    currentNode, _ := l.RetrieveAtIndex(i)
    j := i + 1
    for j < l.Length {
      nodeToCheck, _ := l.RetrieveAtIndex(j)
      if nodeToCheck.Value.Numpy == currentNode.Value.Numpy {
        l.DeleteFromIndex(j)
      } else {
        j++
      }
    }
  }

}

func main() {
  newList := new(DoublyLinkedList)
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 7})
  newList.AddToTail(Value{"Testy", 2})
  newList.AddToHead(Value{"Test", 14})
  newList.InsertAtIndex(4, Value{"Test", 7})
  for i := 0; i < newList.Length; i++ {
    fmt.Println(newList.RetrieveAtIndex(i))
  }
  fmt.Println("~~~~~~~~~~~~")
  newList.RemoveDuplicates()
  for i := 0; i < newList.Length; i++ {
    fmt.Println(newList.RetrieveAtIndex(i))
  }

}
// RUNNERS FOR LINKED LISTS (i.e. ITERATION)
