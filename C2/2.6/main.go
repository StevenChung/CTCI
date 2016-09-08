// Palindrome: Implement a function to check if a linked list is a palindrome.

package main

import (
  "fmt"
)

func (l *DoublyLinkedList) isPalindrome() bool {
  fastRunner, slowRunner := l.Head, l.Head
  firstHalf := []int{}
  fastFinished := false
  handledOdd := false
  // fast runner (iterates two)
  // slow runner (iterates one)
  // both start at Head
  for slowRunner != nil {
    if fastFinished == false {
      firstHalf = append(firstHalf, slowRunner.Value.Numpy)
    } else {
      if slowRunner.Value.Numpy != firstHalf[len(firstHalf) - 1] {
        return false
      } else {
        firstHalf = firstHalf[:len(firstHalf)-1]
      }
    }

    if fastRunner.Next == nil {
      fastFinished = true
      if handledOdd == false {
        handledOdd = true
        firstHalf = firstHalf[:len(firstHalf)-1]
      }
    } else if fastRunner.Next.Next == nil {
      fastFinished = true
    } else {
      fastRunner = fastRunner.Next.Next
    }
    slowRunner = slowRunner.Next
  }
  // when fastRunner.Next.Next is nil, stop
  // continue fastRunner iteration and check against slice of values
  return true
}

func main() {
  newList := &DoublyLinkedList{}
  // memory allocation/value intialization
  // of *DoublyLinkedList type
  // pointer methods!
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 3})
  newList.AddToHead(Value{"Test", 1})
  newList.AddToHead(Value{"Test", 1})
  fmt.Println(newList.isPalindrome(), true)

  newListTwo := &DoublyLinkedList{}
  // memory allocation/value intialization
  // of *DoublyLinkedList type
  // pointer methods!
  newListTwo.AddToHead(Value{"Test", 2})
  newListTwo.AddToHead(Value{"Test", 1})
  newListTwo.AddToHead(Value{"Test", 1})
  newListTwo.AddToHead(Value{"Test", 1})
  newListTwo.AddToHead(Value{"Test", 1})
  newListTwo.AddToHead(Value{"Test", 2})
  fmt.Println(newListTwo.isPalindrome(), true)

  newListThree := &DoublyLinkedList{}
  // memory allocation/value intialization
  // of *DoublyLinkedList type
  // pointer methods!
  newListThree.AddToHead(Value{"Test", 1})
  newListThree.AddToHead(Value{"Test", 1})
  newListThree.AddToHead(Value{"Test", 3})
  newListThree.AddToHead(Value{"Test", 2})
  newListThree.AddToHead(Value{"Test", 1})
  fmt.Println(newListThree.isPalindrome(), false)

  newListFour := &DoublyLinkedList{}
  // memory allocation/value intialization
  // of *DoublyLinkedList type
  // pointer methods!
  newListFour.AddToHead(Value{"Test", 2})
  newListFour.AddToHead(Value{"Test", 1})
  newListFour.AddToHead(Value{"Test", 8})
  newListFour.AddToHead(Value{"Test", 7})
  newListFour.AddToHead(Value{"Test", 1})
  newListFour.AddToHead(Value{"Test", 2})
  fmt.Println(newListFour.isPalindrome(), false)

}
