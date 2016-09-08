/*

Sum Lists: You have two numbers represented by a linked list,
where each node contains a single digit. The digits are
stored in reverse order, such that the 1's digit
is at the head of the list. Write a function
that adds the two numbers and returns the sum as a linked list.
EXAMPLE
Input: (7-) 1 -) 6) + (5 -) 9 -) 2).Thatis,617 + 295.
Output: 2 -) 1 -) 9.That is, 912.

FOLLOW UP
Suppose the digits are stored in forward order.
Repeat the above problem. EXAMPLE
Input: (6 -) 1 -) 7) + (2 -) 9 -) 5).
Thatis,617 + 295. Output: 9 -) 1 -) 2.That is, 912.

Note: Always consider negative integers!
*/


package main

import (
  "fmt"
  "strconv"
  "errors"
)

func (l *DoublyLinkedList) SingleInteger() (int, error) {
  var err error
  sliceOfStrings := []string{}
  currentNode := l.Head
  firstCheck := true
  for currentNode != nil {
    if firstCheck {
      sliceOfStrings = append(sliceOfStrings, strconv.Itoa(currentNode.Value.Numpy))
      firstCheck = false
    } else {
      if currentNode.Value.Numpy < 0 {
        err = errors.New("Negative value found after head")
        break
      } else {
        sliceOfStrings = append(sliceOfStrings, strconv.Itoa(currentNode.Value.Numpy))
      }
    }
    currentNode = currentNode.Next
  }

  str := ""
  for _, val := range(sliceOfStrings) {
    str += val
  }
  numpy, _ := strconv.Atoi(str)
  return numpy, err

}

func SumTwoLists(listOne *DoublyLinkedList, listTwo *DoublyLinkedList) (int, error) {
  currentNodeOne, currentNodeTwo := listOne.Head, listTwo.Head
  integerSlice := []int{}
  for currentNodeOne != nil && currentNodeTwo != nil {

  }
}

func main() {
  firstList := &DoublyLinkedList{}
  firstList.AddToHead(Value{"Test", 1})
  firstList.AddToHead(Value{"Test", 2})
  firstList.AddToHead(Value{"Test", 3})
  integerOne, erryOne := firstList.SingleInteger()

  secondList := &DoublyLinkedList{}
  secondList.AddToHead(Value{"Test", 2})
  secondList.AddToHead(Value{"Test", 2})
  secondList.AddToHead(Value{"Test", 3})
  integerTwo, erryTwo := secondList.SingleInteger()

  if erryOne != nil && erryTwo != nil {
    fmt.Println(erryOne, erryTwo, "invalid lists")
  } else {
    fmt.Println(integerOne + integerTwo)
  }
  // naive solution is to iterate through both lists and generate integers
  // better solution is to pass in two linkedlists and iterate through both at the same time
  // also, the O(1) space solution is clever
  // consider lists 2 -> 1 -> 7 // 5 -> 1 -> 4 -> 2
  // 5142
  //  712
  // Iterates through as if you're adding by hand!
  // carry over if a given basic sum is greater than 10


}
